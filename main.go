package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"rwa/internal/app"
)

const PORT = "HTTP_PORT"

func init() {
	err := godotenv.Load(".env")
	if err != nil {

		fmt.Println(err.Error())
	}
}

func main() {
	h := app.GetApp()
	addr := ":" + os.Getenv(PORT)

	fmt.Println("start server at", addr)
	err := http.ListenAndServe(addr, h)
	if err != nil {
		fmt.Println(err)
		return
	}
}
