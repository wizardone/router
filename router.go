package main

import (
	"fmt"
	"net/http"
)

func routerHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Print("Welcome\n")
}


func main() {
	/*s := &http.Server{
		Addr:           ":8080",
		Handler:        routerHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}*/
	fmt.Print("Starting router\n")
	http.HandleFunc("/", routerHandler)
	http.ListenAndServe(":8080", nil)
}
