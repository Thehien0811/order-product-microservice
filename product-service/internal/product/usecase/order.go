package usecase

type CreateProductInput struct {
	ID       string
	Name     string
	Quantity int32
}

type DetailProduct struct {
	ID       string
	Name     string
	Quantity int32
}
