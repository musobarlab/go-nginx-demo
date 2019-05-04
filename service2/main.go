package main

import (
	"fmt"
	"net/http"
	"strings"
)

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHelloHandler)

	fmt.Printf("service running on port %s\n", ":9001")
	if err := http.ListenAndServe(":9001", nil); err != nil {
		panic(err)
	}

}
