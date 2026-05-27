package types

type Student struct{
	Name string `validate:"required,min=2,max=100"`
	Email string `validate:"required,email"`
	Age int `validate:"required,gte=18,lte=100"`
}