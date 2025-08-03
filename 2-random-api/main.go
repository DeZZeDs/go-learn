package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

const (
	randomNumberRange = 6
)

func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getRandomNumber()))
		return
	})

	server.ListenAndServe()
}

func getRandomNumber() []byte {
	return []byte(strconv.Itoa(rand.Intn(randomNumberRange) + 1))
}
