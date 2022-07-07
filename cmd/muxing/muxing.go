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
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", nameHandler).Methods(http.MethodGet)
	router.HandleFunc("/bad", badHandler).Methods(http.MethodGet)
	router.HandleFunc("/data", dataHandler).Methods(http.MethodPost)
	router.HandleFunc("/headers", headersHandler).Methods(http.MethodPost)
	router.HandleFunc("/", catchAllHandler)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if value, ok := vars["PARAM"]; ok {
		fmt.Fprintf(w, "Hello, %s!", value)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func badHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "I got message:\n%s", string(data))
	}
}

func headersHandler(w http.ResponseWriter, r *http.Request) {
	a, errA := strconv.Atoi(r.Header.Get("a"))
	b, errB := strconv.Atoi(r.Header.Get("b"))

	if errA != nil || errB != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Add("a+b", strconv.Itoa(a+b))
	}
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
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
