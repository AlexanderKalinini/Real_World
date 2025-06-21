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

	newApp, err := app.NewApp()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func(newApp *app.App) {
		err := newApp.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(newApp)

	addr := ":" + os.Getenv(PORT)

	fmt.Println("start server at", addr)
	err = http.ListenAndServe(addr, newApp.Router.Router)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
