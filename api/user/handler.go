package user

import (
	. "bibo.api/api/rest"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var dao = UserDAO{}

func CreateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	user.Id = primitive.NewObjectID()

	user.PasswordSalt = []byte(os.Getenv("AUTH_SECRET"))
	passHash, err := encryptPassword(user.Password)

	if err != nil {
		fmt.Println(err) // TODO: Error handling
	}

	user.PasswordHash = passHash

	if err := dao.Insert(user); err != nil {
		return ResponseInternalError(c)
	}

	token, err := generateUserToken(user)
	formattedResponse := map[string]interface{}{"token": token}

	if err != nil {
		return ResponseInternalError(c)
	}

	return ResponseCreated(c, formattedResponse)
}

func encryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func generateUserToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"name": user.Nickname,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	})

	return token.SignedString([]byte(os.Getenv("AUTH_SECRET")))
}