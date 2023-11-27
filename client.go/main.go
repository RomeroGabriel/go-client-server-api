package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getFile() *os.File {
	fileName := "cotacao.txt"
	finalFileName := "./" + fileName

	path, err := os.Getwd()
	if err == nil {
		finalFileName = path + "/" + fileName
	} else {
		log.Print(err.Error())
	}

	file, err := os.OpenFile(finalFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

func saveOutputIntoFile(data string) {
	linePattern := "Dollar: "
	file := getFile()
	defer file.Close()
	finalString := linePattern + data
	_, err := file.WriteString(finalString)
	if err != nil {
		panic(err)
	}

}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3000*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	saveOutputIntoFile(string(res))
}
