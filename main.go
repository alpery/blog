package main

import (
	"blog/internal/database"
	"blog/internal/middleware"
	"blog/internal/pkg/auth"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the JWT Secret Key from the environment variables
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set in .env file")
	}

	// Initialize the database connection
	db, err := database.InitMongoDB("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize middleware
	jwtMiddleware := middleware.NewJWTMiddleware([]byte(jwtSecret))

	// Create a new instance of the AuthenticationHandler
	authHandler := &handler.AuthHandler{
		DB:  db,
		JWT: auth.NewJWTHandler([]byte(jwtSecret)),
	}

	// Create a new instance of the BlogHandler
	blogHandler := &handler.BlogHandler{
		DB: db,
	}

	// Serve the authentication endpoints
	http.Handle("/api/auth", authHandler.HandleAuth())

	// Serve the blog endpoints
	http.Handle("/api/blog", blogHandler.HandleBlog())
	http.Handle("/api/blog/{id}", blogHandler.HandleBlog())

	// Serve the protected endpoints with JWT middleware
	http.Handle("/api/protected", jwtMiddleware.HandleProtected(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Protected content"))
	})))

	// Start the server
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
