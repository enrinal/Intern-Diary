package order


//PENDING const for declare statusOrder = 1
const PENDING = 1

//PROCESS const for declare statusOrder = 2
const PROCESS = 2

//SEND const for declare statusOrder = 3
const SEND = 3

//DELIVERED const for declare statusOrder = 4
const DELIVERED = 4

func NewOrder(ID int, Item string, Status int) *Order{
	order := new(Order)
	order.ID = ID
	order.Item = Item
	order.Status = Status
	return order
}

func (order *Order) OrderProcess()  {
	order.Status = PROCESS
}

func (order *Order) OrderSend()  {
	order.Status = SEND
}

func (order *Order) OrderDelivered()  {
	order.Status = DELIVERED
}

func (order Order) OrderStatus() int{
	return order.Status
}