package usecase

type CreateOrderInput struct {
	ID        string
	ProductID string
	Quantity  int32
}

type DetailOrder struct {
	ID        string
	ProductID string
	Quantity  int32
}
