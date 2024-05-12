package main

import (
	"fmt"
	"google-translator/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Something wrong while reading the file.")
		return
	}
	router.Routing()
}
