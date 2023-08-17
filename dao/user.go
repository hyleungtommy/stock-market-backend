package dao

import "time"

type User struct {
	User_id           int       `json:"user_id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Registration_date time.Time `json:"registration_date"`
	Funds             float32   `json:"funds"`
}
