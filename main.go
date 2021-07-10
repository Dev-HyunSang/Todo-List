package main

import (
	"net/http"

	"todo-list/app"
)

func main() {
	http.ListenAndServe(":3000", app.NewHandler())
}
