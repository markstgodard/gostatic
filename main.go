package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var data string

func main() {
	data = strings.Repeat("Z", 1024*1024)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, data)
}
