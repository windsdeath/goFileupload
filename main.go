package main

import (
	"log"
	"net/http"

	// Note this is my path according to my GOPATH, chage it according to yours.
	"windsdeath/goUploadFile/controllers"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/upload", controllers.UploadFile)
	log.Println("Running")
	http.ListenAndServe(":8080", nil)
}
