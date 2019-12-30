package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func CreateServerReplicas(server *http.ServeMux, initPortNumber int, NumberOfReplicas int) {
	for index := 0; index < NumberOfReplicas; index++ {
		serverPortNumber := ":" + strconv.Itoa(initPortNumber+index)

		go func(server *http.ServeMux, serverPortNumber string) {
			log.Println("Creating Server with port", serverPortNumber)
			log.Fatal(http.ListenAndServe(serverPortNumber, server))
		}(server, serverPortNumber)
	}
	select {}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	route := http.NewServeMux()
	route.HandleFunc("/hello", IndexHandler)

	CreateServerReplicas(route, 8080, 20)
}
