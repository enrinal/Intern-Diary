package repository

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/config"
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
	var selector = bson.M{"idcust": 2}
	err = collection.Find(selector).All(&result)
	if err != nil {
        return nil, err
	}
	return result,nil
}

func Add(item string, qty int64, idcart int64 ) error {
	var session, err = configs.ConnectMO()
	if err != nil {
		return err
	}
	defer session.Close()
	var collection = session.DB("simpleorder").C("cart")
	cart := bson.M{"idcart": idcart}
	additem := bson.M{"$push": bson.M{"item": bson.M{"item_name": item, "qty": qty}}}
	err := collection.Update(cart, additem)
	if err != nil {
		return err
	}
	return nil
}

func Remove(item string, qty int64) error {
	var session, err = configs.ConnectMO()
	if err != nil {
		return err
	}
	defer session.Close()
	var collection = session.DB("simpleorder").C("cart")

	var selector = bson.M{"idcart": idcart}
    err = collection.Remove(selector)
    if err != nil {
        return err
    }
	return nil
}