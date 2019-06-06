package order

import "testing"

func TestOrderProcess(t *testing.T) {
	areaTest := []struct {
		orderStatus int
		want        string
	}{
		{PENDING, "Order Process"},
		{SEND, "Error Change Order Status to Process"},
	}
	for _, tt := range areaTest {
		order := NewOrder(1, "Item", tt.orderStatus)
		got := order.OrderProcess()
		if got != tt.want {
			t.Errorf("got %s want %s", got, tt.want)
		}
	}
}

func TestOrderSend(t *testing.T) {
	areaTest := []struct {
		orderStatus int
		want        string
	}{
		{PROCESS, "Order Send"},
		{DELIVERED, "Error Change Order Status to Send"},
	}
	for _, tt := range areaTest {
		order := NewOrder(1, "Item", tt.orderStatus)
		got := order.OrderSend()
		if got != tt.want {
			t.Errorf("got %s want %s", got, tt.want)
		}
	}
}

func TestOrderDelivered(t *testing.T) {
	areaTest := []struct {
		orderStatus int
		want        string
	}{
		{SEND, "Order Delivered"},
		{DELIVERED, "Error Change Order Status to Delivered"},
	}
	for _, tt := range areaTest {
		order := NewOrder(1, "Item", tt.orderStatus)
		got := order.OrderDelivered()
		if got != tt.want {
			t.Errorf("got %s want %s", got, tt.want)
		}
	}
}
