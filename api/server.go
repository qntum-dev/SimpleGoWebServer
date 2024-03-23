package api

import (
	"net/http"
	"os"
	todo "todo/todo/libs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type errPortNotFound string

func (e errPortNotFound) Error() string {
	return "Port string not found"
}

func CreateServer() (*http.Server, error) {
	godotenv.Load("../.env")
	portString := os.Getenv("PORT")
	if portString == "" {
		return nil, errPortNotFound(portString)
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/healthy", todo.HandlerReadiness)
	v1Router.Get("/err", todo.HandleErr)
	router.Mount("/v1", v1Router)
	return &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}, nil

}
