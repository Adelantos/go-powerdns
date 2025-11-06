package networks

import "context"

type Client interface {
	// List all registered networks and views in a server
	ListNetworks(ctx context.Context, serverID string) ([]Network, error)

	// Return the view associated to the given network
	GetNetworkView(ctx context.Context, serverID, ip string, prefixLen int) (*string, error)

	// Sets the view associated to the given network
	SetNetworkView(ctx context.Context, serverID, ip string, prefixLen int, view string) error
}
