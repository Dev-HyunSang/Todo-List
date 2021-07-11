package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	if len(todoMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Users")
		return
	}
	users := []*ToDo{}
	for _, u := range todoMap {
		users = append(users, u)
	}
	data, _ := json.Marshal(users)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
	fmt.Print(data, string(data))
}

func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	todo, ok := todoMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID:", id)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(todo)
	fmt.Fprint(w, string(data))
	fmt.Print(w, string(data))
}

func AddTodoListHandler(w http.ResponseWriter, r *http.Request) {
	AddTodo := new(ToDo)
	err := json.NewDecoder(r.Body).Decode(&AddTodo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	lastID++
	AddTodo.ID = lastID
	AddTodo.CreatedAt = time.Now()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(AddTodo)
	fmt.Fprint(w, string(data))
	fmt.Println(string(data))
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	_, ok := todoMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID:", id)
		return
	}
	delete(todoMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User ID:", id)
}

func NewHandler() http.Handler {
	mux := mux.NewRouter()
	fs := http.FileServer(http.Dir("./public/"))

	mux.HandleFunc("/", IndexHandler).Methods("GET")
	mux.HandleFunc("/todo", TodoHandler).Methods("GET")
	mux.HandleFunc("/todo", AddTodoListHandler).Methods("POST")
	mux.HandleFunc("/todo/{id:[0-9]+}", GetTodoListHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", DeleteTodoHandler).Methods("DELETE")

	// Public => JS, CSS PathPrefix
	mux.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	return mux
}
