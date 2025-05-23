package db

import (
	"blog/core/models"
	"blog/core/seeders"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDialect() *gorm.Dialector {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get variables
	dbURL := os.Getenv("DATABASE_URL")
	dialector := postgres.Open(dbURL)

	return &dialector
}

func GetConnection() (*gorm.DB, error) {
	dialector := getDialect()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error
			ParameterizedQueries:      true,        // Don't include arguments in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(*dialector, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err) // Use %w for wrapping errors in Go 1.13+
	}

	err = db.AutoMigrate(&models.User{}, &models.Profile{}, &models.UserFollowing{}, &models.Article{}, &models.Tag{}, &models.Category{}, &models.ArticleTag{})
	if err != nil {
		return nil, fmt.Errorf("error in database prepare: %w", err)
	}

	return db, nil
	// TODO: Setup connection pool
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get generic database object: %w", err)
	// }

	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetConnMaxLifetime(time.Hour)

	// return sqlDB, nil
}

// runSeedUser seeds the database with a specified number of user records.
// It creates a new UserSeeder instance using the provided database connection,
// sets the number of users to seed, and calls the Seed method.
// If an error occurs during seeding, it panics. Otherwise, it prints the number of seeded users.
// Parameters:
// - database: the database connection to use for seeding.
func runSeedUser(database *gorm.DB) {
	userSeeder := seeders.NewUserSeeder(database)

	count := 100
	_, err := userSeeder.Seed(&count)
	if err != nil {
		panic(err)
	}
	fmt.Println("Seeded", count, "users")
}

// runSeedArticle seeds the database with a specified number of article records.
// It creates a new ArticleSeeder instance using the provided database connection,
// sets the number of articles to seed, and calls the Seed method.
// If an error occurs during seeding, it panics. Otherwise, it prints the number of seeded articles.
// Parameters:
// - database: the database connection to use for seeding.
func runSeedArticle(database *gorm.DB) {
	articleSeeder := seeders.NewArticleSeeder(database)

	count := 100
	_, err := articleSeeder.Seed(&count)
	if err != nil {
		panic(err)
	}
	fmt.Println("Seeded", count, "articles")
}

// runSeedMetadata seeds the database with category and tag records.
// It creates new CategorySeeder and TagSeeder instances using the provided database connection
// and calls their Seed methods without a specific count, resulting in all available categories
// and tags being seeded. If an error occurs during seeding, it panics.
// Parameters:
// - database: the database connection to use for seeding.
func runSeedMetadata(database *gorm.DB) {
	categorySeeder := seeders.NewCategorySeeder(database)
	tagSeeder := seeders.NewTagSeeder(database)

	_, err := categorySeeder.Seed(nil)
	if err != nil {
		panic(err)
	}

	_, err = tagSeeder.Seed(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Seeded categories and tags for articles")
}

func __closeConnection(database *gorm.DB) {
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func __runSeed(database *gorm.DB) {

	// runSeedUser(database)
	runSeedArticle(database)
	// runSeedMetadata(database)
}

func RunSeed() {
	database, err := GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer __closeConnection(database)

	__runSeed(database)
}
