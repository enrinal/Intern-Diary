package cart

type Usecase interface {
	AddItem(item string, qty int64) error
	RemoveItem(item string) error
}