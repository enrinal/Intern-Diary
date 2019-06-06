package customer

import "testing"

func TestOrder(t *testing.T) {
	buyer := NewCustomer(1, "Enrinal", 1, 1, 1)
	buyer.AddOrder()
	got := buyer.SumOrder()
	want := 2
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestAddItem(t *testing.T) {
	buyer := NewCustomer(1, "Enrinal", 1, 1, 1)
	s := make([]string, 5)
	got := buyer.AddItem(s)
	want := "Item Added"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
