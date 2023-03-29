package main

import (
	"github.com/zeynepotu/go-course/pkg/config"
	"github.com/zeynepotu/go-course/pkg/handlers"
	"net/http"
)

func main() {
	var app config.AppConfig
	http.HandleFunc("/", handlers.Home)
	http.ListenAndServe("localhost:8080", nil)
}
