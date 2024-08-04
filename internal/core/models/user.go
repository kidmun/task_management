package models

import "context"

type Role string

const (
	Admin      Role = "Admin"
	NormalUser Role = "NormalUser"
)
const (
	CollectionUser = "users"
)

type UserInput struct {
	Username string `json:"username" binding:"required" bson:"username"`
	Password string `json:"password" binding:"required" bson:"password"`
}
type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"-" bson:"password"`
	Role     Role   `json:"role" bson:"role"`
}
type UserRepository interface {
	RegisterUser(c context.Context, user User) (*User, error)
	LoginUser(c context.Context, userInput UserInput) (string, error)
	RegisterAdmin(c context.Context, user User) (*User, error)
	FindUserByUsername(username string) (*User, error)
	CheckPassword(user *User, password string) bool
}

type UserUsecase interface {
	RegisterUser(c context.Context, user User) (*User, error)
	LoginUser(c context.Context, userInput UserInput) (string, error)
	RegisterAdmin(c context.Context, user User) (*User, error)
}