package main

import (
	"log"
	"net/http"

	"github.com/bashbeebo/gommerce/handlers"
)

const port string = ":8080"
func main() {
	http.HandleFunc("/",handlers.Home)
	err := http.ListenAndServe(port,nil)
	if err != nil {
		log.Println(err)
	}
}