package request

type ProductCreateRequest struct {
	Name  string `validate:"required,min=1,max=200"`
	Price int64  `validate:"required,min=0"`
}
