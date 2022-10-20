package request

type ProductUpdateRequest struct {
	Id    int    `validate:"required"`
	Name  string `validate:"required,max=200"`
	Price int    `validate:"required,min=0,numeric"`
}
