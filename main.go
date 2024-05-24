package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"practice/controllers"
	"practice/database"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed to load environment file")
	} else {
		fmt.Println("Environment file loaded successfully")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Ping Failed")
		panic(err)
	} else {
		fmt.Println("DB Ping Success")
	}

	database.DbMigrate(DB)

	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons", controllers.UpdatePerson)
	router.DELETE("/persons", controllers.DeletePerson)

	router.Run(":8080")
}
