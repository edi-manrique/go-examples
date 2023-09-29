package main

import (
	"log"
	"net/http"
)

func main(){
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("Servido estático corriendo en: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}