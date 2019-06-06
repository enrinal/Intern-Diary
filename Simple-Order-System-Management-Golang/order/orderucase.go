package order

const (
	PENDING   int = 1
	PROCESS   int = 2
	SEND      int = 3
	DELIVERED int = 4
)

func NewOrder(ID int, Item string, Status int) *Order {
	order := new(Order)
	order.ID = ID
	order.Item = Item
	order.Status = Status
	return order
}

func (order *Order) OrderProcess() string {
	if order.Status == PENDING {
		order.Status = PROCESS
		return "Order Process"
	}
	return "Error Change Order Status to Process"
}

func (order *Order) OrderSend() string {
	if order.Status == PROCESS {
		order.Status = SEND
		return "Order Send"
	}
	return "Error Change Order Status to Send"
}

func (order *Order) OrderDelivered() string {
	if order.Status == SEND {
		order.Status = DELIVERED
		return "Order Delivered"
	}
	return "Error Change Order Status to Delivered"
}

func (order Order) OrderStatus() int {
	return order.Status
}
