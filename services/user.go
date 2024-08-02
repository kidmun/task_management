package services

import (
	"context"
	"errors"
	"task_management/config"
	"task_management/models"
	"time"

	"github.com/dgrijalva/jwt-go"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string      `json:"username"`
	Role     models.Role `json:"role"`
	jwt.StandardClaims
}

var userCollection *mongo.Collection

func InitUserCollection(client *mongo.Client) {
	userCollection = client.Database("task_management_db").Collection("users")
}
func FindUserByUsername(username string) (*models.User, error) {
	ctx := context.Background()
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func CheckPassword(user *models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
func RegisterUser(user *models.User) (*models.User, error) {
	ctx := context.Background()
	existingUser, err := FindUserByUsername(user.Username)
	if err == nil && existingUser != nil {
		return nil, errors.New("username already taken")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hash)
	user.Role = "NormalUser"
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func LoginUser(userInput *models.UserInput) (string, error) {
	user, err := FindUserByUsername(userInput.Username)
	if err != nil || !CheckPassword(user, userInput.Password) {
		return "", errors.New("wrong Credentials")
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	JwtSecret := config.GetEnv("Jwt_Secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func RegisterAdmin(user *models.User) (*models.User, error) {
	ctx := context.Background()
	existingUser, err := FindUserByUsername(user.Username)
	if err == nil && existingUser != nil {
		return nil, errors.New("username already taken")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hash)
	user.Role = "Admin"
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
