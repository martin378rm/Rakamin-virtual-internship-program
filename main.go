package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status     string                 `json:"status"`
	StatusCode int                    `json:"status_code"`
	Error      string                 `json:"error"`
	Data       map[string]interface{} `json:"data"`
}

type Todo struct {
	Items string `json:"items"`
}

type Todolist []Todo

var todolist Todolist = []Todo{
	{
		Items: "Go",
	},
	{
		Items: "Python",
	},
	{
		Items: "Javac",
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Test Index")
	writeResponse(w, "Success", http.StatusOK, "no error", map[string]interface{}{"key": "Someone doing same thing in another parallel universe"})
}

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello todo")
	writeResponse(w, "Success", http.StatusOK, "no error", map[string]interface{}{"result": todolist})
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello todo")
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		writeResponse(w, "Error", http.StatusBadRequest, err.Error(), nil)
		return
	}
	todolist = append(todolist, todo)
	writeResponse(w, "Success", http.StatusOK, "no error", map[string]interface{}{"result": todolist})
}
func writeResponse(res http.ResponseWriter, status string, statusCode int, err string, data map[string]interface{}) error {
	res.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(res).Encode(Response{
		Status:     status,
		StatusCode: statusCode,
		Error:      err,
		Data:       data,
	})
}

func readContact(r *http.Request) (Todo, error) {
	var t Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		return Todo{}, err
	}
	return t, nil
}

func main() {
	r := echo.New()
	r.GET("/", echo.WrapHandler(http.HandlerFunc(Index)))
	r.GET("/todo", echo.WrapHandler(http.HandlerFunc(GetAllTodo)))
	r.POST("/todo", echo.WrapHandler(http.HandlerFunc(CreateTodo)))

	r.Start(":9090")
}
