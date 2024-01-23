package main

import (
	"myapp/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var todos = []domain.Todo{}

func GetAllTodo(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c echo.Context) error {
	idS, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Data Not Found"})
	}
	id := int64(idS)
	for _, todo := range todos {
		if todo.ID == id {
			return c.JSON(http.StatusOK, todo)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Data Not Found"})
}

func CreateTodo(c echo.Context) error {
	var t domain.Todo
	if err := c.Bind(&t); err != nil {
		return err
	}
	todos = append(todos, t)
	return c.JSON(http.StatusOK, t)
}

func DeleteTodo(c echo.Context) error {
	idS, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Data Not Found"})
	}
	id := int64(idS)
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{"message": "Data Deleted"})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Data Not Found"})
}

func UpdateTodo(c echo.Context) error {
	idS, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Data Not Found"})
	}
	id := int64(idS)
	for i, todo := range todos {
		if todo.ID == id {
			var t domain.Todo
			if err := c.Bind(&t); err != nil {
				return err
			}
			todos[i] = t
			return c.JSON(http.StatusOK, t)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Data Not Found"})
}

func main() {
	e := echo.New()
	e.GET("todo", GetAllTodo)
	e.GET("todo/:id", GetTodoByID)
	e.POST("todo", CreateTodo)
	e.PUT("todo/:id", UpdateTodo)
	e.DELETE("todo/:id", DeleteTodo)
	e.Start(":8080")

}
