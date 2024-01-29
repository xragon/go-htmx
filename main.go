package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

// func getRoot(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got / request\n")
// 	io.WriteString(w, "This is my website!\n")
// }

// func getHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got /hello request\n")
// 	io.WriteString(w, "Hello, HTTP!\n")
// }

func main() {
	page := Page()
	hello := hello("STEEEEEVE")

	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(page))
	// mux.HandleFunc("/hello", templ.Handler(component))
	mux.Handle("/hello", templ.Handler(hello))

	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
