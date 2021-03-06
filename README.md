# ngrokker
[![GoDoc](https://godoc.org/github.com/mfresonke/ngrokker?status.svg)](https://godoc.org/github.com/mfresonke/ngrokker)

`ngrokker` wraps the `ngrok` shell command, allowing you to programmatically create an introspective tunnel. For more information about ngrok, see https://ngrok.com.

Created for the [send2phone](https://github.com/mfresonke/send2phone) utility.

This package is in alpha state. Contributions are welcome and encouraged!
## Prerequisites
You must have `ngrok` installed and available on your `$PATH`. See https://ngrok.com/download.
## Installing
This package is `go get`able.
```bash
go get github.com/mfresonke/ngrokker
```
## Example Usage
### Code
```go
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

```
### Output
```
Endpoint 1 - https://9545bfed.ngrok.io
Endpoint 2 - http://9545bfed.ngrok.io
Making request from outside world to https://9545bfed.ngrok.io/hello-world
Thanks, @inconshreveable!
```

## Issues
See https://github.com/mfresonke/ngrokker/issues

## Contributing
Please fork the repo, make your changes, and create a PR. Make sure you `gofmt`, `golint`, and `go vet` your code!

## Notice
All users of ngrok MUST accept the [ngrok terms of service](https://ngrok.com/tos) before opening a tunnel.
