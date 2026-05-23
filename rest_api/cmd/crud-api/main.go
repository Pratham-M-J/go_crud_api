package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Pratham-M-J/crud_api/internal/config"
)

func main() {
	fmt.Println("Welcome to crud api")
	//load config
	cfg := config.MustLoad()
	//database setup
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to CRUD API"))
	})
	//setup server
	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	fmt.Printf("Server Started %s", cfg.HTTPServer.Addr)

	err := server.ListenAndServe()

	if err != nil{
		log.Fatal("failed to start server")
	}

	
}