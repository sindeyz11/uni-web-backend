package internal

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"uni-web/internal/infrastructure/config"
	"uni-web/internal/infrastructure/persistence"
	"uni-web/internal/interfaces"
	"uni-web/internal/interfaces/middleware"
)

func Run() {
	mux := http.NewServeMux()
	dbConf := config.NewConfig().MySqlConfig()

	services, err := persistence.NewRepositories(config.NewMysqlConn(dbConf))
	if err != nil {
		log.Fatal(err)
	}
	defer services.Close()

	// Handler for static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/static/"))))

	mux.HandleFunc("/task1/", interfaces.Task1)
	mux.HandleFunc("/task2/", interfaces.Task2)
	mux.HandleFunc("/task3/", interfaces.Task3)

	handler := middleware.Logging(mux)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
