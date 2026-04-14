package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bakibillahrahat/auth-system/internal/database"
	"github.com/bakibillahrahat/auth-system/internal/handlers"
)

func main() {
	// 1. Database Connect befor starting the server
	database.ConnectDB()

	// 2. Setup the (Router/Endpoints)
	// When someone send the request using /api/v1/register, It calls the handlers.Register function
	http.HandleFunc("api/v1/register", handlers.Register)

	// 3. Setup port (Which is come from docker or .env file if not get it will run on port 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 4. Server is Running
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}