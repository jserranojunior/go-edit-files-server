package main

import (
    "fmt"
	"net/http"
	"../editfiles"
)

func editFile(w http.ResponseWriter, req *http.Request) {	
	text := "It's Work!"
	var arrayContent []string	
	data := append(arrayContent, text)
	file := editfiles.SetValueFile("../index.html", data)
	fmt.Fprintf(w, file)
}

func main() {
	http.HandleFunc("/file", editFile)
	fmt.Println("Server 8090")
	http.ListenAndServe(":8090", nil)	
}