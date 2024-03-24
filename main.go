package main

import (
	"log"
	"net/http"
	"todos/config"
	"todos/controllers/categorycontroller"
	"todos/controllers/homecontroller"
)

func main() {
	config.ConnectDB()

	// 1. HomePage

	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
