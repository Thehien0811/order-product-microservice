package usecase

type CreateOrderInput struct {
	ID       string
	Name     string
	Quantity int
}

type DetailOrder struct {
	ID       string
	Name     string
	Quantity int
}
