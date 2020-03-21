package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zignd/turbo-octo-robot/api"
)

func main() {
	http.HandleFunc("/account", api.PostAccountHandler)
	http.HandleFunc("/", api.RootHandler)

	const addr = "localhost:8080"
	fmt.Printf("Running on http://%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
