package models

import (
	"time"
)

type User struct {
	ID       int       `json:"id"`
	UserName string    `json:"username"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Number   int       `json:"number"`
	CreateAt time.Time `json:"createdAt"`
}
