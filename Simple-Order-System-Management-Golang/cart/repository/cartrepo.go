package repository

import (
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/configs"
	"gopkg.in/mgo.v2"
)

func FindByCustomerId(idcust int64) ([]*models.Cart, error) {
	var session, err = configs.ConnectMO()
	if err != nil {
        fmt.Println("Error!", err.Error())
        return
    }
	defer session.Close()
	var collection = session.DB("simpleorder").C("cart")
	var result = []models.Cart{}
	var selector = bson.M{"idcust": 2}
	err = collection.Find(selector).All(&result)
	if err != nil {
        return nil, error
	}
	return result,nil
}

func Add(item string, qty int64) error {

	return nil
}

func Remove(item string, qty int64) error {
	return nil
}
