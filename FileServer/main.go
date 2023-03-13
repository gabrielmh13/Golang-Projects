package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/public/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Blog"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
