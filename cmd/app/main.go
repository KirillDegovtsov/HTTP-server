package main

import (
	"log"
	objectHttp "my_project/api/http"
	"my_project/cmd/config"
	pkgObject "my_project/pkg/http"
	"my_project/repository/ram_storage"
	"my_project/usecases/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.MustLoad()

	objectRepo := ram_storage.NewObject()
	objectService := service.NewObject(objectRepo)
	objectHandlers := objectHttp.NewHandler(objectService)

	r := chi.NewRouter()
	// r.Get("/swagger/*", httpSwagger.WrapHandler)
	objectHandlers.WithObjectHandlers(r)

	log.Printf("Starting server")
	if err := pkgObject.CreateAndRunServer(r, cfg.Address); err != nil {
		log.Fatal("Failed to start server: #{err}")
	}
}
