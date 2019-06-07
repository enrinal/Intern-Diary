package customer

import "testing"

func TestOrder(t *testing.T) {
	areaTest := []struct {
		numOrder   int
		custStatus int
		want       string
	}{
		{1, RegularBuyer, "Order Added"},
		{1, SubcriptionBuyer, "Order Added"},
		{5, RegularBuyer, "Order Added"},
		{6, RegularBuyer, "Order limit exceeded"},
		{6, SubcriptionBuyer, "Order Added"},
	}
	for _, tt := range areaTest {
		buyer := NewCustomer(1, "Customer", 0, 1, tt.custStatus)
		got := buyer.AddOrder(tt.numOrder)
		if got != tt.want {
			t.Errorf("got %s want %s", got, tt.want)
		}
	}
}

func TestAddItem(t *testing.T) {
	areaTest := []struct {
		numItem int
		want    string
	}{
		{5, "Item Added"},
		{10, "Item Added"},
		{15, "Your basket limit exceeded"},
	}
	for _, tt := range areaTest {
		buyer := NewCustomer(1, "Customer", 0, 0, 1)
		item := make([]string, tt.numItem)
		got := buyer.AddItem(item)
		if got != tt.want {
			t.Errorf("got %s want %s", got, tt.want)
		}
	}
}
