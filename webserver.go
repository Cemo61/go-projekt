package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Willkommen auf meiner Website")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server läuft auf Port 8080")
	http.ListenAndServe(":8080", nil)
}
