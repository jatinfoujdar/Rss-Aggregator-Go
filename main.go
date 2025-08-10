package main

import (
    "fmt"
    "log"
    "os"
    "net/http"

    "github.com/go-chi/cors"
    "github.com/go-chi/chi/v5"
    "github.com/joho/godotenv"
)

func main() {
    
     godotenv.Load()

    portString := os.Getenv("PORT") 
    if portString == ""{
        log.Fatal("PORT environment variable is not set")
    }

       router := chi.NewRouter()
        
        router.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"http://*", "https://*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"*"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300, 
    }))

        srv := &http.Server{
            Handler: router,
            Addr:   ":" + portString,
        }

        log.Printf("Starting server on port %s\n", portString)
        err := srv.ListenAndServe() 
         if err != nil {
         log.Fatal("Failed to start server: ", err)
     }


        fmt.Printf("Server will run on port: %s\n", portString)
           
}