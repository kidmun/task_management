package models

type Role string

const (
	Admin      Role = "Admin"
	NormalUser Role = "NormalUser"
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
