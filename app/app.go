package app

import (
	"net/http"
	"time"
)

type ToDo struct {
	ID        int       `json:"id"`
	Content   string    `json:"content'`
	Completed bool      `json:completed`
	CreatedAt time.Time `json:"created_at"`
}

var (
	todoMap map[int]*ToDo
)

func indexHandler(w http.Response, r *http.Request) {

}

func NewHandler() {
	http.Handle("/", http.FileServer(http.Dir("../public")))
}
