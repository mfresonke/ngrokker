package tunneler

// Interface is the main interface that represents a type that can open an introspective
//  tunnel. The Open method creates and starts the tunnel, and returns the introspective
//  url, while the Close method closes the tunnel and cleans up all associated resources
type Interface interface {
	Open(port int) ([]Endpoint, error)
	Close() error
}

// Endpoint represets a publicly accessible URL that tunnels to your machine
type Endpoint struct {
	URL    string
	Secure bool
}
