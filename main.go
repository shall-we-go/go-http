package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		b, err := json.Marshal(struct {
			Message string `json:"message"`
		}{
			Message: "pong",
		})
		if err != nil {
			log.Fatal("error:", err)
		}
		w.Write(b)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
