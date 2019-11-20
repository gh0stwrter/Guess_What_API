package modeluser

import (
	model "app/src/model"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"upper.io/db.v3"
)

type id bson.ObjectId

type User struct {
	ID       id     `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	Password string `bson:"password"`
	Score    int    `bson:"score"`
}

type Claims struct {
	IDUser id
	Name   string
	jwt.StandardClaims
}

var user User
var jwtKey = []byte("my_secret_key")
var orm = model.DatabaseSession()

func CreateUser(name string, password string) string {
	userCollection := orm.Collection("user")

	hash, _ := hashPassword(password)
	userCollection.Insert(User{
		Name:     name,
		Password: hash,
	})

	defer orm.Close()
	return "User Create successfully"
}

func Login(name string, password string) string {
	userCollection := orm.Collection("user")
	res := userCollection.Find(db.Cond{"name": name})
	res.One(&user)
	match := CheckPasswordHash(password, user.Password)
	if match == false {
		CreateUser(name, password)
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		IDUser: user.ID,
		Name:   user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "Error"
	}

	return tokenString
}

func SetScore(score int, id string) string {
	userCollection := orm.Collection("user")
	res := userCollection.Find(db.Cond{"_id": id})
	res.One(&user)
	res.Update(User{
		Score: score,
	})
	fmt.Println(user)

	return "Score insert"
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
