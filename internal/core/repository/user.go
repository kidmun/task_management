package repository

import (
	"context"
	"errors"

	"task_management/internal/config"
	"task_management/internal/core/models"

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

type userRepository struct {
	database   *mongo.Database
	collection string
}
func NewUserRepository(db *mongo.Database, collection string) models.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}
// var userCollection *mongo.Collection

// func InitUserCollection(client *mongo.Client) {
// 	userCollection = client.Database("task_management_db").Collection("users")
// }
func (ur *userRepository) FindUserByUsername(username string) (*models.User, error) {
	ctx := context.Background()
	var user models.User
	err := ur.database.Collection(ur.collection).FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (ur *userRepository) CheckPassword(user *models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
func (ur *userRepository) RegisterUser(c context.Context, user models.User) (*models.User, error) {
	
	existingUser, err := ur.FindUserByUsername(user.Username)
	if err == nil && existingUser != nil {
		return nil, errors.New("username already taken")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hash)
	user.Role = "NormalUser"
	_, err = ur.database.Collection(ur.collection).InsertOne(c, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) LoginUser(c context.Context, userInput models.UserInput) (string, error) {
	user, err := ur.FindUserByUsername(userInput.Username)
	if err != nil || !ur.CheckPassword(user, userInput.Password) {
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

func (ur *userRepository) RegisterAdmin(c context.Context, user models.User) (*models.User, error) {
	
	existingUser, err := ur.FindUserByUsername(user.Username)
	if err == nil && existingUser != nil {
		return nil, errors.New("username already taken")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hash)
	user.Role = "Admin"
	_, err = ur.database.Collection(ur.collection).InsertOne(c, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
