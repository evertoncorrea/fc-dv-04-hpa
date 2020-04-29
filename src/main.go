package main

import (
	"fmt"
	"log"
	"math"
	"message"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	x := 0.0001
	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}

	fmt.Fprintf(w, message.Message())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
