package main

import (
	"GoDatabase/src/handlers"
	"GoDatabase/src/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	// * Prepare the environment variables.
	// * For now we will just use hard coded environment key value pairs.
	utils.Environment()

	// * Get all the environment variables.
	port := os.Getenv("PORT")

	// * Logs the lived routes.
	log.Print(utils.TextGreen("Server live at port http://localhost:" + port))

	// * These are all the server routes.
	http.HandleFunc("GET /ready", handlers.ReadyHandler)
	http.HandleFunc("GET /redis/ready", handlers.RedisReadyHandler)
	http.HandleFunc("GET /sql/ready", handlers.SQLReadyHandler)
	http.HandleFunc("POST /sql/create/user/table", handlers.SQLCreateUserTableHandler)

	// * You can setup the server here.
	s := &http.Server{
		Addr:           "localhost:" + port,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
