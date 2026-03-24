package dtos

type InputUserDTO struct {
	First_Name string `json:"first_name" validate:"required,min=3,max=100"`
	Last_Name  string `json:"last_name" validate:"required,min=3,max=100"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6,max=100"`
}
