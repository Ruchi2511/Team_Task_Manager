package model

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}
type UserExist struct {
	ID       string `db:"id"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Name     string `db:"name"`
}
type User struct {
	ID    string `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Role  string `db:"role" json:"role"`
}
