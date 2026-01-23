package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	//test for initializing the backends
	parsedURL1, err := url.Parse("http://localhost:8081")
	if err != nil {
		panic(err)
	}
	parsedURL2, err := url.Parse("http://localhost:8082")
	if err != nil {
		panic(err)
	}
	parsedURL3, err := url.Parse("http://localhost:8083")
	if err != nil {
		panic(err)
	}

	backend1 := &Backend{URL: parsedURL1, Alive: true, CurrentConns: 0}
	backend2 := &Backend{URL: parsedURL2, Alive: true, CurrentConns: 0}
	backend3 := &Backend{URL: parsedURL3, Alive: true, CurrentConns: 0}

	// initializing the serverPool
	s := &ServerPool{
		Backends: []*Backend{backend1, backend2, backend3},
		Current:  0,
	}

	//Creating the proxy object
	ProxyObject := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			backend := s.GetNextValidPeer()
			req.URL.Scheme = backend.URL.Scheme
			req.URL.Host = backend.URL.Host
			println("Forwarding request to:", backend.URL.String())
		},
		// this is really non sense here but it can help me change the response form
		ModifyResponse: func(resp *http.Response) error {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			fmt.Println("Backend response :", string(bodyBytes))
			// reset the body so the client still receives it
			resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			return nil
		},
	}

	// launching the main server
	http.ListenAndServe(":8087", ProxyObject)
}
