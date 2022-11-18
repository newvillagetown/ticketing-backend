package mongodb

import "time"

type RefreshToken struct {
	Created time.Time `json:"created"`
	Email   string    `json:"email"`
	Token   string    `json:"token"`
}
