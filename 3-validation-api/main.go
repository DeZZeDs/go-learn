package main

import (
	"main/config"
	"main/verify"
	"net/http"
)

func main() {
	config := config.BuildConfig()
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	verify.VerifyHandler(router, config)
	server.ListenAndServe()
}
