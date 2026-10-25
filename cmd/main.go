package main

import (
	"fmt"
	"taskQ/internal/redis"
	 "log"
    "net/http"
)

func main (){
	if err :=  redis.Connect(); err != nil {
		fmt.Printf("Error connecting to Redis: %v\n", err)
		return
	}

	  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Backend is running"))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}