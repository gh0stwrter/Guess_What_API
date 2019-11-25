package main

import (
	socketmanager "app/src/controller/socket"
	usercontroller "app/src/controller/user"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

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
func setupRoutes(w http.ResponseWriter, r *http.Request) {
	pool := socketmanager.NewPool()
	go pool.Start()

	serveWs(pool, w, r)
}
func main() {
	// Create a simple file server
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/ws", setupRoutes)
	r.HandleFunc("/score", usercontroller.ScoreData)
	r.HandleFunc("/create-game", usercontroller.CreateRoom)
	r.HandleFunc("/sign-in", usercontroller.SignIn)
	r.HandleFunc("/sign-up", usercontroller.SignUp)
	r.HandleFunc("/logout", usercontroller.Logout)

	log.Println("http server started on :8000")
	c := cors.AllowAll()

	handler := c.Handler(r)
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
