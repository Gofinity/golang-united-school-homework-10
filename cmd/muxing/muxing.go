package main

import (
	"fmt"
	"io"
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
	r := mux.NewRouter()

	r.HandleFunc("/name/{PARAM}", nameHandler).Methods("GET")
	r.HandleFunc("/bad", badRouteHandler).Methods("GET")
	r.HandleFunc("/data", dataHandler).Methods("POST")
	r.PathPrefix("/").HandlerFunc(defaultHandler)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), r); err != nil {
		log.Fatal(err)
	}
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

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading body:\n%v", err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I got message:\n%s", string(body))
}

func badRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %v!", vars["PARAM"])
}
