package main

import (
	roomcontroller "app/src/controller/room"
	socketmanager "app/src/controller/socket"
	usercontroller "app/src/controller/user"
	"fmt"

	"flag"
	"github.com/rs/cors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func serveWs(pool *socketmanager.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := socketmanager.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	client := &socketmanager.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func main() {
	flag.Parse()
	// Create a simple file server
	r := mux.NewRouter()
	c := cors.AllowAll()
	pool := socketmanager.NewPool()
	go pool.Start()
	http.Handle("/", r)
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
	r.HandleFunc("/users", usercontroller.GetAllUsers)
	r.HandleFunc("/score", usercontroller.ScoreData)
	r.HandleFunc("/sign-in", usercontroller.SignIn)
	r.HandleFunc("/sign-up", usercontroller.SignUp)
	r.HandleFunc("/logout", usercontroller.Logout)
	r.HandleFunc("/rooms", roomcontroller.GetAllRoom)
	r.HandleFunc("/joined", roomcontroller.JoinedUser)
	r.HandleFunc("/create-game", roomcontroller.CreateRoom)


	log.Println("http server started on :8000")

	handler := c.Handler(r)
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
