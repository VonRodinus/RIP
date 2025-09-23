package main

import (
	"RIP/internal/db"
	"RIP/internal/handlers"
	"log"
	"net/http"
)

func main() {
	db.Init()

	http.HandleFunc("/", handlers.ArtifactCatalogHandler)
	http.HandleFunc("/artifact/", handlers.ArtifactDetailHandler)
	http.HandleFunc("/tpq_request/", handlers.BuildingTPQCalcHandler)
	http.HandleFunc("/add_artifact/", handlers.AddArtifactToRequestHandler)
	http.HandleFunc("/delete_request/", handlers.DeleteRequestHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
