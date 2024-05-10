package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"uni-web/internal/infrastructure/config"
	"uni-web/internal/infrastructure/persistence"
	"uni-web/internal/interfaces"
	"uni-web/internal/interfaces/middleware"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found. Using OS environment variables")
	}
	mux := http.NewServeMux()

	dbConf := config.NewConfig().MySqlConfig()
	services, err := persistence.NewRepositories(config.NewMysqlConn(dbConf))
	if err != nil {
		log.Fatal(err)
	}
	defer services.Close()

	formService := interfaces.NewForm(services.Form, services.Language)

	// Handler for static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/static/"))))

	mux.HandleFunc("/task1/", interfaces.Task1)
	mux.HandleFunc("/task2/", interfaces.Task2)
	mux.HandleFunc("/task3/", formService.Task3)

	handler := middleware.Logging(mux)

	fmt.Printf("Starting server at port 8080\nhttp://127.0.0.1:8080/\n")

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
