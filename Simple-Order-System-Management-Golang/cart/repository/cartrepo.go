package repository

import (
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
	var selector = bson.M{"idcust": 2}
	err = collection.Find(selector).All(&result)
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
	if err != nil{
		_,err = collection.UpsertId(cart.IDCart, cart)
		if err != nil {
			return err
		}
	}else{
	match := bson.M{"idcart" : cart.IDCart}
	change := bson.M{"$push":bson.M{"item":bson.M{"$each":cart.Items}}}
	err = collection.Update(match,change)
	//_,err = collection.UpsertId(cart.IDCart, cart)
	if err != nil {
		return err
	}
}
	return nil
}


func Remove(item string, qty int64, idcart int64) error {
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

func insert() error { //Testing aja untuk buat attribut baru
	var session, err = configs.ConnectMO()
	if err != nil {
		return err
	}
	defer session.Close()

	var collection = session.DB("simpleorder").C("cart")
	err = collection.Insert(&models.Cart{IDCart: 1, IDCust: 2, Items: []models.Item{models.Item{Id: int64(1), Name: "mobil"},
		models.Item{Id: int64(1), Name: "mobil"}}})
	if err != nil {
		return err
	}
	return nil
}