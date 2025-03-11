package main

import (
	"github.com/joho/godotenv"
	"rwa/cmd/command"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
}

func main() {
	err := command.Execute()
	if err != nil {
		return
	}
}
