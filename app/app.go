package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "../public/todo.html", http.StatusTemporaryRedirect)
}

func NewHandler() http.Handler {
	mux := mux.NewRouter()

	mux.Handle("/static", http.FileServer(http.Dir("../public")))
	mux.HandleFunc("/", IndexHandler).Methods("GET")
	return mux
}
