package order

import "testing"

func TestOrderProcess(t *testing.T){
	order := NewOrder(1,"House",PENDING)
	order.OrderProcess()
	got := order.OrderStatus()
	want := PROCESS
	if got != want {
        t.Errorf("got %d want %d", got, want)
    }
}

func TestOrderSend(t *testing.T){
	order := NewOrder(1,"House",PROCESS)
	order.OrderSend()
	got := order.OrderStatus()
	want := SEND
	if got != want {
        t.Errorf("got %d want %d", got, want)
    }
}

func TestOrderDelivered(t *testing.T){
	order := NewOrder(1,"House",SEND)
	order.OrderProcess()
	got := order.OrderStatus()
	want := DELIVERED
	if got != want {
        t.Errorf("got %d want %d", got, want)
    }
}