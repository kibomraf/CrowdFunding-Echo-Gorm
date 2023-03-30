package users

import "time"

type Users struct {
	Id               int
	Name             string
	Occupation       string
	Email            string
	Password_hash    string
	Avatar_file_name string
	Role             string
	Token_vuejs      string
	Created_at       time.Time
	Updated_at       time.Time
}
