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
	todoMap map[int]ToDo = map[int]ToDo{}
	lastID  int
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func TodosHandler(w http.ResponseWriter, r *http.Request) {
	if len(todoMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "NO TODOS")
		return
	}
	todos := []ToDo{}
	for _, i := range todoMap {
		todos = append(todos, i)
	}
	data, _ := json.Marshal(todos)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	fmt.Println(id, vars, todoMap[id])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	todo, ok := todoMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "NO TODO ID:", id)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(todo)
	fmt.Fprint(w, string(data))
}

func AddTodoListHandler(w http.ResponseWriter, r *http.Request) {

	todo := new(ToDo)
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	lastID++
	todo.ID = lastID
	todo.CreatedAt = time.Now()
	todoMap[lastID] = *todo

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(todo)
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
		fmt.Fprint(w, "NO TODO ID:", id)
		return
	}
	delete(todoMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "DELETED TODO ID:", id)
}

func UpdateTodoListHandler(w http.ResponseWriter, r *http.Request) {
	updateTodo := new(ToDo)
	err := json.NewDecoder(r.Body).Decode(updateTodo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	todo, ok := todoMap[updateTodo.ID]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "NO TODO ID:", updateTodo.ID)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(updateTodo)
	fmt.Fprint(w, string(data))
}

func NewHandler() http.Handler {
	mux := mux.NewRouter()
	fs := http.FileServer(http.Dir("./public/"))

	mux.HandleFunc("/", IndexHandler).Methods("GET")
	mux.HandleFunc("/todo", TodosHandler).Methods("GET")
	mux.HandleFunc("/todo", AddTodoListHandler).Methods("POST")
	// mux.HandleFunc("/todo").Methods("PUT")
	mux.HandleFunc("/todo/{id:[0-9]+}", GetTodoListHandler).Methods("GET")
	mux.HandleFunc("/todo/{id:[0-9]+}", DeleteTodoHandler).Methods("DELETE")

	// Public => JS, CSS PathPrefix
	mux.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	return mux
}
