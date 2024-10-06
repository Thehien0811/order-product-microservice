package usecase

type CreateOrderInput struct {
	ID       string
	Name     string
	Quantity int32
}

type DetailOrder struct {
	ID       string
	Name     string
	Quantity int32
}
