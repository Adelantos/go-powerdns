package pdns

import (
	"context"

	"github.com/mittwald/go-powerdns/apis/cryptokeys"
	"github.com/mittwald/go-powerdns/apis/metadata"
	"github.com/mittwald/go-powerdns/apis/networks"
	"github.com/mittwald/go-powerdns/apis/tsigkey"
	"github.com/mittwald/go-powerdns/apis/views"

	"github.com/mittwald/go-powerdns/apis/cache"
	"github.com/mittwald/go-powerdns/apis/search"
	"github.com/mittwald/go-powerdns/apis/servers"
	"github.com/mittwald/go-powerdns/apis/zones"
)

// Client is the root-level interface for interacting with the PowerDNS API.
// You can instantiate an implementation of this interface using the "New" function.
type Client interface {

	// Status checks if the PowerDNS API is reachable. This does a simple HTTP connection check;
	// it will NOT check if your authentication is set up correctly (except you're using TLS client
	// authentication.
	Status() error

	// WaitUntilUp will block until the PowerDNS API accepts HTTP requests. You can use the "ctx"
	// parameter to make this method wait only for (or until) a certain time (see examples).
	WaitUntilUp(ctx context.Context) error

	// Servers returns a specialized API for interacting with PowerDNS servers
	Servers() servers.Client

	// Zones returns a specialized API for interacting with PowerDNS zones
	Zones() zones.Client

	// Search returns a specialized API for searching
	Search() search.Client

	// Cache returns a specialized API for caching
	Cache() cache.Client

	// Cryptokeys returns a specialized API for cryptokeys
	Cryptokeys() cryptokeys.Client

	// Metadata returns a specialized API for metadata
	Metadata() metadata.Client

	// Views returns a specialized API for views
	Views() views.Client

	// Networks returns a specialized API for networks
	Networks() networks.Client

	// TsigKeys returns a specialized API for TSIG keys
	TsigKeys() tsigkey.Client
}
