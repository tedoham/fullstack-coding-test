package dto

type ListProductRequest struct {
	Limit  int
	Offset int
}

type SingleProductRequest struct {
	ProductId int
}
