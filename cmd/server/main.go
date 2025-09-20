package main

import (
	"RIP/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Настройка маршрутов
	http.HandleFunc("/", handlers.ArtifactCatalogHandler)
	http.HandleFunc("/artifact/", handlers.ArtifactDetailHandler)
	http.HandleFunc("/tpq_request/", handlers.BuildingTPQCalcHandler)

	// Обслуживание статических файлов
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
