package tunneler

// Interface represents a type that can open an introspective tunnel.
type Interface interface {
	// Open creates and starts the tunnel, and returns the introspective urls
	Open(port int) ([]Endpoint, error)
	// Close method closes the tunnel and cleans up all associated resources
	Close() error
}

// Endpoint represets a publicly accessible URL that tunnels to your machine
type Endpoint struct {
	URL    string
	Secure bool
}
