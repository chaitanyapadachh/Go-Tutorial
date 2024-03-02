package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	//Use the http.NewServeMux() function to initialize a new servemux, then
	//register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	// Use the http.ListenAndServe() functioin to start a new webserver
	//We pass in two parameters the TCP network address to listen on (:4000)
	//and the servemux we just created.
	//If http.ListenAndServe() returns an error
	//we use log.Fatal() function to log the errorMessage and Exit
	//Note that any error returned by http.ListenAndServe is always non-nill
	log.Println("Starting Server on Port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
