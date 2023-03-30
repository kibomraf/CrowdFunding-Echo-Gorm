package users

type InputRegister struct {
	Name       string `json:"name" validate:"required"`
	Occupation string `json:"occupation" validate:"required"`
	Email      string `json:"email"  validate:"required,email"`
	Password   string `json:"password" validate:"required"`
}
type InputLogin struct {
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
