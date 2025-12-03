package model

type User struct {
	RealName string `json:"realname"`
	Email    string `json:"email"`
	Password string `json:"password"` // Use SHA-1 hashed password
}
