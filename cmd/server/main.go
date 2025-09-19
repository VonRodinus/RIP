package main

import (
	"RIP/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Настройка маршрутов
	http.HandleFunc("/", handlers.CatalogHandler)
	http.HandleFunc("/artifact/", handlers.DetailHandler)
	http.HandleFunc("/order/", handlers.GetOrder)

	// Обслуживание статических файлов
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
