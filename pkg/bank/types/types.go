package types

type Money int64

type Category string

type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
type PAN string

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}
