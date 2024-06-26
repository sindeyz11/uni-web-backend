package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"uni-web/internal/interfaces/handlers"

	"uni-web/internal/infrastructure/config"
	"uni-web/internal/infrastructure/persistence"
	"uni-web/internal/interfaces/middleware"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found. Using OS environment variables")
	}
	mux := http.NewServeMux()

	dbConf := config.NewConfig().PostgresConfig()
	services, err := persistence.NewRepositories(config.NewPostgresConn(dbConf))
	if err != nil {
		log.Fatal(err)
	}
	defer services.Close()

	formService := handlers.NewForm(services.Form, services.Language, services.User)

	// Handler for static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/static/"))))

	mux.HandleFunc("/task1/", handlers.Task1)
	mux.HandleFunc("/task2/", handlers.Task2)
	mux.HandleFunc("/task3/", formService.Task3)
	mux.HandleFunc("/task4/", formService.Task4)
	mux.HandleFunc("/task5/", formService.Task5)

	mux.HandleFunc("/login/", formService.LoginHandler)

	handler := middleware.Logging(mux)

	fmt.Printf("Starting server at port 8080\nhttp://127.0.0.1:8080/\n")

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
