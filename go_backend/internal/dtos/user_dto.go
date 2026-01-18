package dtos


type RegisterUserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserDto struct {
	Identifier string `json:"identifier"` // Can be email or username
	Password   string `json:"password"`
}