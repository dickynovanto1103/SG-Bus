package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	byte, err := w.Write([]byte("halo"))
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	log.Println("byte", byte, "written")

	//call the API here
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
