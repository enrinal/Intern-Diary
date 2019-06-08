package order

const (
	Pending = iota + 1
	Process
	Send
	Delivered
)

func NewOrder(ID int, Item string, Status int) *Order {
	order := new(Order)
	order.ID = ID
	order.Item = Item
	order.Status = Status
	return order
}

func (order *Order) OrderProcess() string {
	if order.Status == Pending {
		order.Status = Process
		return "Order Process"
	}
	return "Error Change Order Status to Process"
}

func (order *Order) OrderSend() string {
	if order.Status == Process {
		order.Status = Send
		return "Order Send"
	}
	return "Error Change Order Status to Send"
}

func (order *Order) OrderDelivered() string {
	if order.Status == Send {
		order.Status = Delivered
		return "Order Delivered"
	}
	return "Error Change Order Status to Delivered"
}

func (order Order) OrderStatus() int {
	return order.Status
}
