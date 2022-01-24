package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	var t Task
	for i := 0; i < 100; i++ {
		res, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", i))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			log.Fatal(err)
		}
		fmt.Println(t.Title)
	}
}
