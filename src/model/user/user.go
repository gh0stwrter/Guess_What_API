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

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Password string        `bson:"password"`
	Score    int           `bson:"score"`
}

type Claims struct {
	IDUser             bson.ObjectId
	Name               string
	jwt.StandardClaims `json:"omitempty"`
}

var userCollection = orm.Collection("user")

var user User
var jwtKey = []byte("my_secret_key")
var orm = model.DatabaseSession()

func CreateUser(name string, password string) string {

	hash, _ := hashPassword(password)
	fmt.Println(hash)
	userCollection.Insert(User{
		Name:     name,
		Password: hash,
	})

	return "User Create successfully"
}

var users []User

func Login(name string, password string) string {
	fmt.Println(name)
	res := userCollection.Find(db.Cond{"name": name})
	err := res.One(&user)

	fmt.Println(user)
	if err != nil {
		return "Wrong Name"
	}

	match := CheckPasswordHash(password, user.Password)
	fmt.Println(match)

	if match == false {
		return "Wrong Password try again"
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
func FindAllUsers() []User {
	res := userCollection.Find()
	res.All(&users)
	return users
}

func SetScore(score int, id bson.ObjectId) string {
	res := userCollection.Find(db.Cond{"_id": id})
	fmt.Println(id)
	err := res.One(&user)
	if err != nil {
		return "Error"
	}
	fmt.Println(user.Score)
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
