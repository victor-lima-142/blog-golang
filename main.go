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
)

func RunSeedUser() {
	// db, err := db.GetConnection()
	// if err != nil {
	// 	panic(err)
	// }

	// userSeeder := seeders.NewUserSeeder(db)

	// count := 100
	// _, err = userSeeder.Seed(&count)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Seeded", count, "users")
}

func RunSeedArticle() {
	// db, err := db.GetConnection()
	// if err != nil {
	// 	panic(err)
	// }

	// articleSeeder := seeders.NewArticleSeeder(db)

	// count := 100
	// _, err = articleSeeder.Seed(&count)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Seeded", count, "articles")
}

func main() {
	db, err := db.GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load()

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
