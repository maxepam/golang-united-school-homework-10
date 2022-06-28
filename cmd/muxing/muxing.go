package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{param}", getNameHandler)
	router.HandleFunc("/bad", getBadHandler)
	router.HandleFunc("/data", postDataHandler).Methods("POST")
	router.HandleFunc("/headers", postHeadersHandler).Methods("POST")


	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

// handler to for get /name/{param} endpoint
func getNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello,  %v!", vars["param"])
}

// handler to for get /bad endpoint
func getBadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

// handler to for post /data endpoint
func postDataHandler(w http.ResponseWriter, r *http.Request) {
	param := r.PostFormValue("PARAM")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I got message:\n%v", param)
}

// handler to for post /headers endpoint
func postHeadersHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.Header.Get("a")
	bStr := r.Header.Get("b")
	a, err := strconv.Atoi(aStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
        	return
	}
	b, err := strconv.Atoi(bStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
        	return
	}
	sum := a + b
	w.Header().Set("a+b",  strconv.Itoa(sum))
	w.WriteHeader(http.StatusOK)
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
