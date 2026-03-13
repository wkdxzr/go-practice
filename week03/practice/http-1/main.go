package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type todo struct {
	Title     string `json:"title"`
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
}

func main() {
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos", nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	var todos []todo
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		fmt.Println("解析失败:", err)
		return
	}
	for _, todo := range todos {
		fmt.Printf("ID: %d, UserID: %d, Title: %s, Completed: %t\n",
			todo.Id, todo.UserId, todo.Title, todo.Completed)
	}
}
