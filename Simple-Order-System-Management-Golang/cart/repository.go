package cart

type Repository interface {
	AddItem(item string, qty int64) error
	RemoveItem(item string) error
}
