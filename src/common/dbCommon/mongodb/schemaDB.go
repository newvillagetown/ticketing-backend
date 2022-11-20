package mongodb

import "time"

type RefreshToken struct {
	Created   time.Time `bson:"created"`
	Email     string    `bson:"email"`
	Token     string    `bson:"token"`
	IsDeleted bool      `bson:"isDeleted"`
}
