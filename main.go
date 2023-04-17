package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(".env file not found")
	}

	mux := NewRouter()

	fmt.Println("[INFO] Server listening")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		fmt.Println(err.Error())
	}
}
