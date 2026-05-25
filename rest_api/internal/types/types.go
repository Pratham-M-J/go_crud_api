package types

type Student struct{
	Id int `validate:"required",gt=0`
	Name string `validate:"required,min=2,max=100"`
	Email string `validate:"required,email"`
	Age int `validate:"required,gte=18,lte=100"`
}