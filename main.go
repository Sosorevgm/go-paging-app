package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"users-paging-app/handlers"
)

func main() {
	// load environment
	err := godotenv.Load("/myexe/paging-users.env")
	if err != nil {
		log.Fatal(err)
	}
	dsn := os.Getenv("DSN")
	port := os.Getenv("PORT")

	// open db connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// launch handlers
	r := gin.Default()
	s := handlers.Service{Db: db}
	r.GET("/api/users", s.GetUsers)

	// start listen
	err = r.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
