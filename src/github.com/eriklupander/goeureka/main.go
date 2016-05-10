package main

import (
        "net/http"
        "sync"
        "log"
        "github.com/eriklupander/goeureka/service"
        "github.com/eriklupander/goeureka/eureka"
)

func main() {


        go startWebServer()

        eureka.Register()

        go eureka.StartHeartbeat()

        defer eureka.Deregister()

        // Block...
        wg := sync.WaitGroup{}
        wg.Add(1)
        wg.Wait()
}

func startWebServer() {
        router := service.NewRouter()
        log.Println("Starting HTTP service at 8080")
        err := http.ListenAndServe(":8080", router)
        if err != nil {
                log.Println("An error occured starting HTTP listener at port 8080")
                log.Println("Error: " + err.Error())
        }
}

