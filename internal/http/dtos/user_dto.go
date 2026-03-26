package dtos

type InputRegisterDTO struct {
	First_Name string `json:"first_name" validate:"required,min=3,max=100"`
	Last_Name  string `json:"last_name" validate:"required,min=3,max=100"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6,max=100"`
}

type InputLoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type OutputUserDTO struct {
	UUID       string `json:"uuid"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
}

type RegisterOutput struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	User    OutputUserDTO `json:"user"`
}

type LoginOutput struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	User    OutputUserDTO `json:"user"`
	Token   string        `json:"token"`
}
