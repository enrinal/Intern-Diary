package repository

import (
	"fmt"

	configs "gitlab.warungpintar.co/enrinal/intern-diary/simple-order/config"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	"gopkg.in/mgo.v2/bson"
)

func FindByCustomerId(idcust int64) ([]models.Cart, error) {
	var session, err = configs.ConnectMO()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	var collection = session.DB("simpleorder").C("cart")
	var result = []models.Cart{}
	var selector = bson.M{"idcust": idcust}
	err = collection.Find(nil).Select(selector).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Add(cart models.Cart) error {
	var session, err = configs.ConnectMO()
	if err != nil {
		return err
	}
	defer session.Close()
	cartuser := models.Cart{}
	var collection = session.DB("simpleorder").C("cart")
	err = collection.Find(bson.M{"idcart": cart.IDCart}).One(&cartuser)
	fmt.Println(cartuser)
	if err != nil {
		_, err = collection.UpsertId(cart.IDCart, cart)
		if err != nil {
			return err
		}
	} else {
		match := bson.M{"idcart": cart.IDCart}
		change := bson.M{"$push": bson.M{"item": bson.M{"$each": cart.Items}}}
		err = collection.Update(match, change)
		if err != nil {
			return err
		}
	}
	return nil
}

func Remove(cart models.Cart) error {
	//Connection Begin
	var session, err = configs.ConnectMO()
	if err != nil {
		return err
	}
	defer session.Close()
	var collection = session.DB("simpleorder").C("cart")

	//Remove cart when item is 0
	if len(cart.Items) == 0 {
		err = collection.Remove(bson.M{"_id": cart.IDCart})
		if err != nil {
			return err
		}
		return nil
	}

	//Update cart when item change
	err = collection.Update(bson.M{"_id": cart.IDCart}, cart)
	if err != nil {
		return err
	}
	return nil
}
