package dto
<<<<<<< HEAD

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type LoginResponse struct {
	User  UserDTO `json:"user"`
	Token string  `json:"token"`
}
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
