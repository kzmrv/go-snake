package main

import (
	"log"
	"net/http"
)

func main_web() {
	log.Fatal(http.ListenAndServe(":8080",
		http.FileServer(http.Dir("/home/vasyl/Projects/go/snake/web")),
	))
}
