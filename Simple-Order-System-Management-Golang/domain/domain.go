package domain

type Customer struct {
	Id   int
	Name string
}

type Item struct {
	Id        int
	Name      string
	Available bool
}

type Order struct {
	Id       int
	Customer Customer
	Items    []Item
	Status   int
}

// func (order *Order) Add(item Item) error {
// 	if !item.Available {
// 		return errors.New("Item Unavaliable")
// 	}
// 	if len(Order.Items)>10{
// 		return error.New(`Can't Order more then 10`)
// 	}
// 	order.Items = append(order.Items, item)
// 	return nil
// }
