package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/mat-cf/image-host/internal/handler"
	"github.com/mat-cf/image-host/internal/repository"
	"github.com/mat-cf/image-host/internal/service"
	"github.com/mat-cf/image-host/internal/storage"
)

func main() {

	dbURL := os.Getenv("DATABASE_URL")
	basePath := os.Getenv("STORAGE_PATH")
	baseURL := os.Getenv("BASE_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error connecting to database:L %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("error pinging database: %v", err)
	}
	
	imageStorage := storage.NewLocalStorage(basePath, baseURL)
	imageRepo := repository.NewPostgresImageRepository(db)
	imageService := service.NewImageService(imageRepo, imageStorage)
	imageHandler := handler.NewImageHandler(imageService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /upload", imageHandler.Upload)
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	log.Print("starting server on :8080")

	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}