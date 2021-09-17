/*  1. This repo was created to understand how the go language functions
    2. Started by creating a simple rest-api
    3. This is the starting point of the Repo

    Tutorial Series: https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_
    Github Repo: https://github.com/nicholasjackson/building-microservices-youtube
    Credits: @nicholasjackson 
*/
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
    "github.com/4molybdenum2/microservices-arch/handlers"
)

func main() {
    l := log.New(os.Stdout, "products-api", log.LstdFlags)
    hh := handlers.SayHello(l)
    ph := handlers.NewProducts(l)

    sm := http.NewServeMux()

    sm.Handle("/", hh)
    sm.Handle("/products", ph)

    server := &http.Server{
        Addr: ":8001",
        Handler: sm,
        ErrorLog: l,
        IdleTimeout: 120 * time.Second,
        ReadTimeout: 1 * time.Second,
        WriteTimeout: 1 * time.Second,
    }
    
    go func() {
        l.Println("Starting server on PORT: 8001")
        err := server.ListenAndServe()
        if err != nil{
            l.Printf("Error starting server %s\n", err)
            os.Exit(1)
        }
    }()

    sigChannel := make(chan os.Signal, 1)
    signal.Notify(sigChannel, os.Interrupt)
    signal.Notify(sigChannel, os.Kill)

    sig := <- sigChannel
    l.Println("Graceful Shutdonwn", sig)
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    server.Shutdown(ctx)
}