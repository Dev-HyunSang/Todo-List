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

func RemoveTodoListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "DLETET TODO LIST ID:", id)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

/* 본 함수는 ToDo struct에서의 Completed만을 다루는 함수입니다.
상태 업데이트만을 다루는 함수입니다. */
func UpdateStateTodoListHandler(w http.ResponseWriter, r *http.Request) {
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

	todo := new(ToDo)
	JsonErr := json.NewDecoder(r.Body).Decode(&todo)
	if JsonErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	todo.ID = id
	todo.CreatedAt = time.Now()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(todo)
	fmt.Fprint(w, string(data))
	fmt.Println(string(data))
}

func NewHandler() http.Handler {
	mux := mux.NewRouter()
	fs := http.FileServer(http.Dir("./public/"))

	mux.HandleFunc("/", IndexHandler).Methods("GET")
	mux.HandleFunc("/todo", TodosHandler).Methods("GET")
	mux.HandleFunc("/todo", AddTodoListHandler).Methods("POST")
	mux.HandleFunc("/todo/completed/{id:[0-9]+}", UpdateStateTodoListHandler).Methods("PUT")
	mux.HandleFunc("/todo/{id:[0-9]+}", GetTodoListHandler).Methods("GET")
	mux.HandleFunc("/todo/{id:[0-9]+}", RemoveTodoListHandler).Methods("DELETE")

	// Public => JS, CSS PathPrefix
	mux.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	return mux
}
