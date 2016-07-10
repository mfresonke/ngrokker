package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mfresonke/ngrokker"
)

func main() {
	// setup an introspective tunnel to port 8080
	tunnel := ngrokker.NewHTTPTunnel(true, false)
	endpoints, _ := tunnel.Open(8080)
	// don't forget to close the tunnel!
	defer tunnel.Close()

	// set up a http server to respond to requests on port 8080
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Thanks, @inconshreveable!")
	})
	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	// find the secure endpoint out of the two ngrok creates by default
	var secureEndpoint ngrokker.Endpoint
	for i, endpoint := range endpoints {
		fmt.Println("Endpoint", i+1, "-", endpoint.URL)
		if endpoint.Secure {
			secureEndpoint = endpoint
		}
	}

	// Make a zero-configuration https request to your own machine!
	// Notice the lack of ":8080"!
	reqURL := secureEndpoint.URL + "/hello-world"
	fmt.Println("Making request from outside world to", reqURL)
	resp, _ := http.Get(reqURL)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
}
