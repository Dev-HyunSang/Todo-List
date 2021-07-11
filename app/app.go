package app

import (
	"encoding/json"
	"fmt"
	"html/template"
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
	lastID  int
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}
func AddTodoListHandler(w http.ResponseWriter, r *http.Request) {
	AddTodo := new(ToDo)
	err := json.NewDecoder(r.Body).Decode(AddTodo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	AddTodo.ID = lastID
	AddTodo.CreatedAt = time.Now()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(AddTodo)
	fmt.Fprint(w, string(data))
	fmt.Println(string(data))
}

func NewHandler() http.Handler {
	mux := mux.NewRouter()
	fs := http.FileServer(http.Dir("./public/"))

	mux.HandleFunc("/", IndexHandler).Methods("GET")
	mux.HandleFunc("/todo", AddTodoListHandler).Methods("POST")

	// Public => JS, CSS PathPrefix
	mux.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	return mux
}
