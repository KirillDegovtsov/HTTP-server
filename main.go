package main

import (
	"flag"
	"log"
	objectHttp "my_project/api/http"
	pkgObject "my_project/pkg/http"
	"my_project/repository/ram_storage"
	"my_project/usecases/service"

	"github.com/go-chi/chi/v5"
)

// @title My API
// @version 1.0
// description This is a sample server.

// @host localhost:8080
// @BasePath /
func main() {
	addr := flag.String("addr", ":8080", "address for http server")

	objectRepo := ram_storage.NewObject()
	objectService := service.NewObject(objectRepo)
	objectHandlers := objectHttp.NewHandler(objectService)

	r := chi.NewRouter()
	// r.Get("/swagger/*", httpSwagger.WrapHandler)
	objectHandlers.WithObjectHandlers(r)

	log.Printf("Starting server on #{*addr}")
	if err := pkgObject.CreateAndRunServer(r, *addr); err != nil {
		log.Fatal("Failed to start server: #{err}")
	}
}
