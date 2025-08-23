package main

import (
    "fmt"
    "log"
    "os"
    "net/http"

    "github.com/go-chi/cors"
    "github.com/go-chi/chi/v5"
    "github.com/joho/godotenv"
    "database/sql"
    "github.com/jatinfoujdar/Rss-Aggregator-Go/internal/database"

    _ "github.com/lib/pq"
)

type apiConfig struct {
    DB *database.Queries
}

func main() {
    
     godotenv.Load()

    portString := os.Getenv("PORT") 
    if portString == ""{
        log.Fatal("PORT environment variable is not set")
    }

    dbURL := os.Getenv("DB_URL") 
    if  dbURL == ""{
        log.Fatal(" dbURL environment variable is not set")
    }

    conn, err := sql.Open("postgres", dbURL)
     if err != nil {
        log.Fatal("Failed to connect to database: ", err)
     }
     
   
     apiCfg := apiConfig{
        DB: database.New(conn),
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


    v1Router := chi.NewRouter()

    v1Router.Get("/healthz", handlerReadiness)
    v1Router.Get("/err",handlerErr)
    v1Router.Post("/users", apiCfg.handlerCreateUser)
    v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.HandlerGetUser))

    v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
    // v1Router.Get("/feeds",apiCfg.handlerGetFeeds) 

    v1Router.Post("/feedfollows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))

    router.Mount("/v1", v1Router)

        srv := &http.Server{
            Handler: router,
            Addr:   ":" + portString,
        }

        log.Printf("Starting server on port %s\n", portString)
        err = srv.ListenAndServe()
         if err != nil {
         log.Fatal("Failed to start server: ", err)
     }


        fmt.Printf("Server will run on port: %s\n", portString)
           
}