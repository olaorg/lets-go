package main

import "net/http"
import "log"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("HelloWorld endpoint received a request.")
		w.Write([]byte("Hello, world!"))
	})

	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye endpoint received a request.")
		w.Write([]byte("Goodbye, fellows!"))
	})

	http.ListenAndServe(":9090", nil)
}