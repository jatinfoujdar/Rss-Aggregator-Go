package main

import (
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
)

func main() {
    fmt.Println("Hello, World!")

     godotenv.Load(".env")

    portString := os.Getenv("PORT") 
    if portString == ""{
        log.Fatal("PORT environment variable is not set")
    }
        fmt.Printf("Server will run on port: %s\n", portString)
    
}