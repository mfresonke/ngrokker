# ngrokker
`ngrokker` wraps the `ngrok` shell command, allowing you to programmatically create an introspective tunnel. For more information about ngrok, see https://ngrok.com.
## Prerequisites
You must have `ngrok` installed and available on your `$PATH`. See https://ngrok.com/download.
## Installing
This package is `go get`able.
```bash
go get github.com/mfresonke/ngrokker
```
## Example Usage
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mfresonke/ngrokker"
	"github.com/mfresonke/ngrokker/tunneler"
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
	var secureEndpoint tunneler.Endpoint
	for i, endpoint := range endpoints {
		fmt.Println("Endpoint", i+1, ": ", endpoint.URL)
		if endpoint.Secure {
			secureEndpoint = endpoint
		}
	}

	// Make a zero-configuration https request to your own machine!
	// Notice the lack of ":8080"!
	resp, _ := http.Get(secureEndpoint.URL + "/hello-world")
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
}

```

## Notice
All users of ngrok MUST accept the [ngrok terms of service](https://ngrok.com/tos) before opening a tunnel.
