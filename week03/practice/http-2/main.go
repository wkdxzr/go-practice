package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed string `json:"completed"`
}

var todos []todo

func init() {
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos", nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic("初始化数据失败" + err.Error())
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&todos)
}

func getTodoDetail(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 5 || path[4] == "" {
		http.Error(w, "无效的 URL 路径，格式应为 /api/todo/detail/数字", http.StatusBadRequest)
		return
	}

	todoid, err := strconv.Atoi(path[4])
	if err != nil {
		http.Error(w, "todoid必须是数字", http.StatusBadRequest)
		return
	}
	for _, todo := range todos {
		if todo.Id == todoid {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
}

func getTodoList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("json编码失败:", err)
		return
	}
	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println("响应写入失败:", err)
		return
	}
}

func main() {
	http.HandleFunc("/api/todo/detail/", getTodoDetail)
	http.HandleFunc("/api/todo/list", getTodoList)

	fmt.Println("服务器开始监听")
	http.ListenAndServe("localhost:8080", nil)
}
