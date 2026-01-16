package main

import (
	"log"
	"net/http"
	"net/url"
)

//"net/http/httputil"

func main() {
	// requesting doing a like client behavior
	requestUrl, err := url.Parse("http://localhost:8081")
	if err != nil {
		log.Fatal("the service you are requesting isn t available :( ")
	}

	// the reverse proxy work
	var reverseProxy ReverseProxy
	//reverseProxy1 := httputil.NewSingleHostReverseProxy(requestUrl)

	http.ListenAndServe(":8080")

}
