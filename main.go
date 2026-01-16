package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

//"net/http/httputil"

func main() {
	// requesting doing a like client behavior
	requestUrl, err := url.Parse("http://localhost:8081")
	if err != nil {
		log.Fatal("the service you are requesting isn t available :( ")
	}
	var req *http.Request

	var reverseProxy httputil.ReverseProxy
	// set req Host, URL and Request URI to forward a request to the origin server
	req.Host = requestUrl.Host
	req.URL.Host = requestUrl.Host
	req.URL.Scheme = requestUrl.Scheme
	req.RequestURI = ""
	reverseProxy.Director(req)

	http.ListenAndServe(":8080", nil)

}
