package main

import (
	"log"
	"net/http"

	"todo-list/app"
)

func main() {
	log.Print("Listen And Server at http://localhost:3000")
	http.ListenAndServe(":3000", app.NewHandler())
}
