package main

import (
	"fmt"
	"net/http"
)

func main() {

	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello from backend 1 "))
		fmt.Println("Hello from backend 1")
	})

	//http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", handler)

}
