package main

import (
	"blog/core/db"
	"blog/src/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func runApp(db *gorm.DB) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")

	router := gin.Default()
	apiBlogRouter := router.Group("/api/blog")
	{
		handlers.UserHandler(apiBlogRouter, db)
		handlers.ProfileHandler(apiBlogRouter, db)
		handlers.ArticleHandler(apiBlogRouter, db)
	}

	runningMsg := fmt.Sprintf("Server is running at http://localhost:%s", port)
	portStr := fmt.Sprintf(":%s", port)

	fmt.Println(runningMsg)
	log.Fatal(http.ListenAndServe(portStr, router))
}
func main() {
	// gormDB, err := gormDB.GetConnection()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	db.RunSeed()
	// runApp(db)
}
