package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var db *pgx.Conn

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func params() string {

	connURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		getEnv("DB_USERNAME", ""),
		getEnv("DB_PASSWORD", ""),
		getEnv("DB_HOST", ""),
		getEnv("DB_PORT", ""),
		getEnv("DB_NAME", ""),
	)

	return connURL
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {
	var err error
	db, err = pgx.Connect(context.Background(), params())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close(context.Background())

	router := gin.Default()
	router.GET("/api/v1/records", getRecords)
	router.GET("/api/v1/records/:id", getRecord)
	router.POST("/api/v1/records", addRecord)
	router.PUT("/api/v1/records/:id", updateRecord)
	router.DELETE("/api/v1/records/:id", deleteRecord)

	router.Run("localhost:8080")
}
