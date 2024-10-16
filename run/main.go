package main

import (
	"net/http"

	"github.com/AkuLaper/akulaperbe/route"
)

func main() {
	http.HandleFunc("/", route.URL)
	http.ListenAndServe(":8080", nil)
}
