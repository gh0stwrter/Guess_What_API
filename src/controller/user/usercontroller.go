package usercontroller

import (
	modeluser "app/src/model/user"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2/bson"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

type User struct {
	IDUser   bson.ObjectId `json:",omitempty"`
	Name     string        `json:",omitempty"`
	Password string        `json:",omitempty"`
	Token    string        `json:",omitempty"`
	Message  string        `json:",omitempty"`
	Score    int           `json:",omitempty"`
}

var user User

func Secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-Auth")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewDecoder(r.Body).Decode(&user)
	session, _ := store.Get(r, "session-Auth")

	token := modeluser.Login(user.Name, user.Password)
	tokenSession := session.Values["token"]
	tokenSession = token

	if tokenSession == token {
		session.Values["authenticated"] = true

	} else {
		session.Values["authenticated"] = false
	}
	data := User{Token: token,
		Message: "Token successful",
	}
	session.Save(r, w)
	json.NewEncoder(w).Encode(data)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewDecoder(r.Body).Decode(&user)
	modeluser.CreateUser(user.Name, user.Password)
	data := User{
		Message: "User create",
	}
	json.NewEncoder(w).Encode(data)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	session, _ := store.Get(r, "session-Auth")
	session.Values["authenticated"] = false
	tokenSession := session.Values["token"]

	if tokenSession == true {
		tokenSession = nil
	}

	session.Save(r, w)
}
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	usersFind := modeluser.FindAllUsers()
	json.NewEncoder(w).Encode(&usersFind)

}

func ScoreData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	res := modeluser.SetScore(user.Score, user.IDUser)
	fmt.Println(res)
	data := User{
		Score: user.Score,
	}
	fmt.Println(data)
	json.NewEncoder(w).Encode(data)
}
