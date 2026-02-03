package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func main() {
	// --- Load backends from config.json ---
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cfg struct {
		Backends []string `json:"backends"`
	}
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		panic(err)
	}

	// Initialize Backend objects
	var backends []*Backend
	for _, u := range cfg.Backends {
		parsedURL, err := url.Parse(u)
		if err != nil {
			panic(err)
		}
		backends = append(backends, &Backend{
			URL:          parsedURL,
			Alive:        true,
			CurrentConns: 0,
		})
	}

	// Create ServerPool using backends from config.json
	s := &ServerPool{
		Backends: backends,
		Current:  0,
	}
	//Creating the proxy object
	ProxyObject := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			backend := s.GetNextValidPeer()
			req.URL.Scheme = backend.URL.Scheme
			req.URL.Host = backend.URL.Host
			//println("Forwarding request to:", backend.URL.String())
		},
		// this is really non sense here but it can help me change the response form
		ModifyResponse: func(resp *http.Response) error {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			//fmt.Println("Backend response :", string(bodyBytes))
			// reset the body so the client still receives it
			resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			return nil
		},
	}

	//Health checker go routine

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			for _, backend := range s.Backends {
				backend.mux.Lock()
				resp, err := http.Get(backend.URL.String())
				if err != nil {
					backend.Alive = false
					backend.mux.Unlock()
					fmt.Println("backend ", backend.URL.String(), " is not alive")
					continue
				}
				backend.Alive = true
				backend.mux.Unlock()
				fmt.Println("backend ", backend.URL.String(), " is alive")
				resp.Body.Close()

			}

		}

	}()

	go func() {
		mux := http.NewServeMux()

		// GET /status
		mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			s.BackendsMutex.RLock()
			defer s.BackendsMutex.RUnlock()
			json.NewEncoder(w).Encode(s.Backends)
		})

		// POST /backends and DELETE /backends
		mux.HandleFunc("/backends", func(w http.ResponseWriter, r *http.Request) {
			var body struct {
				URL string `json:"url"`
			}
			json.NewDecoder(r.Body).Decode(&body)

			switch r.Method {
			case http.MethodPost:
				parsed, err := url.Parse(body.URL)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				newBackend := &Backend{
					URL:          parsed,
					Alive:        true,
					CurrentConns: 0,
				}
				s.BackendsMutex.Lock()
				s.Backends = append(s.Backends, newBackend)
				s.BackendsMutex.Unlock()
				w.WriteHeader(http.StatusCreated)

			case http.MethodDelete:
				s.BackendsMutex.Lock()
				for i, b := range s.Backends {
					if b.URL.String() == body.URL {
						s.Backends = append(s.Backends[:i], s.Backends[i+1:]...)
						break
					}
				}
				s.BackendsMutex.Unlock()
				w.WriteHeader(http.StatusOK)

			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		})

		http.ListenAndServe(":8090", mux)
	}()

	// launching the main server
	http.ListenAndServe(":8087", ProxyObject)

}
