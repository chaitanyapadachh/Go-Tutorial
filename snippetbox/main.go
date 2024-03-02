package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

//Adding a View Function

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// Checking whether such a parameter like id exists and if a valid integer has been passed
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Disaply a specifice snippet with ID %d", id)
	// w.Write([]byte("Display a specific snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	//Using r.Method to check whether request is a POST method or not
	if r.Method != "POST" {

		//Setting Allow methods to POST so user can know which methods to use
		w.Header().Set("Allow", "POST")
		//If not Post then use w.WriteHeader to send 405 status and send message method not allowed
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	//Use the http.NewServeMux() function to initialize a new servemux, then
	//register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
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
