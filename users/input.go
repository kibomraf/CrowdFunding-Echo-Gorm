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
type ChekEmail struct {
	Email string `json:"email"  validate:"required,email"`
}
type UploadImage struct {
	Filename string `json:"file_name" validate:"required"`
	Data     string `json:"data" validate:"required"`
}
