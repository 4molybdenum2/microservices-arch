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
    "github.com/4molybdenum2/rest-api/handlers"
)

func main() {
    l := log.New(os.Stdout, "rest-api", log.LstdFlags)
    hh := handlers.SayHello(l)

    sm := http.NewServeMux()

    sm.Handle("/", hh)

    server := &http.Server{
        Addr: ":8001",
        Handler: sm,
        IdleTimeout: 120 * time.Second,
        ReadTimeout: 1 * time.Second,
        WriteTimeout: 1 * time.Second,
    }
    
    go func() {
        err := server.ListenAndServe()
        if err != nil{
            l.Fatal(err)
        }
    }()

    sigChannel := make(chan os.Signal)
    signal.Notify(sigChannel, os.Interrupt)
    signal.Notify(sigChannel, os.Kill)

    sig := <- sigChannel
    l.Println("Graceful Shutdonwn", sig)
    tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
    server.Shutdown(tc)
}