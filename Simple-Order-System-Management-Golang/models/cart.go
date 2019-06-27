package models

type Cart struct {
	IDCart int64  `bson:"idcart"`
	IDCust int64  `bson:"idcust"`
	Items  []Item `bson:"item"`
	Qty    int64  `bson:"qty"`
	Prize  int64  `bson:"prize"`
}
