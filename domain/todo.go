package domain

import "log"

type Todo struct {
	ID   int64  `json:"id"`
	Task string `json:"task"`
}

func (todo *Todo) String() string {
	return todo.Task
}

func main() {
	t := Todo{
		ID:   1,
		Task: "Go",
	}
	log.Println(t.String())

}
