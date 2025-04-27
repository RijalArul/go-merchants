package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"go-merchants/src/middleware"
	"go-merchants/src/routes"
	"go-merchants/src/tools"

	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // CLI flag
    mode := flag.String("mode", "server", "mode: server | generate_customers | generate_merchants | generate_all")
    flag.Parse()

    switch *mode {
    case "generate_customers":
        tools.GenerateCustomers()
        return
    case "generate_merchants":
        tools.GenerateMerchants()
        return
    case "generate_all":
        tools.GenerateCustomers()
        tools.GenerateMerchants()
        log.Println("✅ Successfully generated customers and merchants!")
        return
    case "server":
        startServer()
    default:
        log.Fatal("Unknown mode, use: server | generate_customers | generate_merchants | generate_all")
    }
}

func startServer() {
    // Check necessary environment variables
    if os.Getenv("JWT_SECRET_KEY") == "" {
        log.Fatal("JWT_SECRET_KEY is not set in environment")
    }
    if os.Getenv("ENCRYPTION_SECRET_KEY") == "" {
        log.Fatal("ENCRYPTION_SECRET_KEY is not set in environment")
    }
    ensureDataExists()

	// Set up routes

	routes.IndexRoutes()

	handler := http.DefaultServeMux

    // Wrap semua middleware
    finalHandler := middleware.IPWhitelistMiddleware(
        middleware.RateLimitMiddleware(
            middleware.CORSMiddleware(handler),
        ),
    )



    log.Println("Server starting at :8080 🚀")
    err := http.ListenAndServe(":8080", finalHandler)
    if err != nil {
        log.Fatal(err)
    }

}

// ensureDataExists - Check if JSON files exist; generate if missing
func ensureDataExists() {
    // Check customers.json
    if _, err := os.Stat("data/customers.json"); os.IsNotExist(err) {
        log.Println("⚡ customers.json not found, generating fake customers...")
        tools.GenerateCustomers()
    } else {
        log.Println("✅ customers.json found")
    }

    // Check merchants.json
    if _, err := os.Stat("data/merchants.json"); os.IsNotExist(err) {
        log.Println("⚡ merchants.json not found, generating fake merchants...")
        tools.GenerateMerchants()
    } else {
        log.Println("✅ merchants.json found")
    }
}
