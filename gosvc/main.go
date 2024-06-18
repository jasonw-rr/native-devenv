package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func getBasicHandlerfunc() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Getting request...")
	})
}

func main() {
	fmt.Println("Starting server...")
	fmt.Println(os.Getenv("AAA"))
	fmt.Println(os.Getenv("BBB"))
	server := &http.Server{
		Addr:              "0.0.0.0:8080",
		WriteTimeout:      time.Second * 600,
		ReadHeaderTimeout: time.Second * 60,
		ReadTimeout:       time.Second * 60,
		IdleTimeout:       time.Second * 120,
		Handler:           getBasicHandlerfunc(),
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Shutting down server...")
	}
}
