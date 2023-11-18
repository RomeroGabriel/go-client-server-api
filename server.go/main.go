package main

import (
	"fmt"
	"log"
	"net/http"
)

func QuotationHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hi"))
	res.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/cotacao", QuotationHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server is listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
