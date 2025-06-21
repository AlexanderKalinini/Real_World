package main

import (
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
}

func main() {
	err := Execute()
	if err != nil {
		return
	}
}
