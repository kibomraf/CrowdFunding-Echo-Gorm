package users

type FormatRegister struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Token      string `json:"token"`
}

func FormatterUsers(user Users, token string) FormatRegister {
	formatter := FormatRegister{
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Password:   user.Password_hash,
		Token:      token,
	}
	return formatter
}
