package main

import (
	"fmt"
	"log"
	"os"
	"todo/todo/api"

	"github.com/joho/godotenv"
)

func main() {

	s, err := api.CreateServer()
	if err != nil {
		log.Fatal(err)
		return
	}
	godotenv.Load()
	portString := os.Getenv("PORT")

	fmt.Printf("Listening on %v", portString)
	s.ListenAndServe()

}
