package main

import (
	"log"
	"net/http"

	handler "github.com/12ain13owz/project-borntodev/api"
	db "github.com/12ain13owz/project-borntodev/database"
	"github.com/12ain13owz/project-borntodev/middleware"
	_ "github.com/go-sql-driver/mysql"
)

// MySQL
// CREATE TABLE users (
// 	id INT AUTO_INCREMENT PRIMARY KEY,
// 	name VARCHAR(255),
// 	email VARCHAR(255),
// );

func SetUpRoutes() {
	users := http.HandlerFunc(handler.HandlerUsers)
	http.Handle("/api/user", middleware.EnableCORS(users))
	user := http.HandlerFunc(handler.HandlerUserByID)
	http.Handle("/api/user/", middleware.EnableCORS(user))
}

func main() {
	err := db.SetupDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Database.Close()

	SetUpRoutes()
	log.Fatal(http.ListenAndServe(":7000", nil))
}
