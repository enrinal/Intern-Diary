package repository

import (
	"testing"

	configs "gitlab.warungpintar.co/enrinal/intern-diary/simple-order/config"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	"gopkg.in/mgo.v2/bson"
)

func TestAdd(t *testing.T) {
	var session, _ = configs.ConnectMO()
	defer session.Close()
	var collection = session.DB("simpleorder").C("cart")
	cart := models.Cart{IDCart: 1000, IDCust: 2, Items: []models.Item{models.Item{Id: int64(1), Name: "mobil"},
		models.Item{Id: int64(1), Name: "mobil"}}}

	t.Log("test service ")
	if got := Add(cart); got != nil {
		t.Errorf("Error %s", got)
	}
	_ = collection.Remove(bson.M{"_id": 1000})

}

func TestFindByCustomerId(t *testing.T) {
	var session, _ = configs.ConnectMO()
	defer session.Close()
	var collection = session.DB("simpleorder").C("cart")
	cart := models.Cart{IDCart: 1000, IDCust: 2, Items: []models.Item{models.Item{Id: int64(1), Name: "mobil"},
		models.Item{Id: int64(1), Name: "mobil"}}}
	_, _ = collection.UpsertId(cart.IDCart, cart)
	if _, got := FindByCustomerId(2); got != nil {
		t.Errorf("Error %s", got)
	}
	_ = collection.Remove(bson.M{"_id": 1000})
}
