package views

import "context"

type Client interface {
	// List all views in a server
	ListViews(ctx context.Context, serverID string) ([]string, error)

	// List the contents of a given view
	ListViewZones(ctx context.Context, serverID, view string) ([]string, error)

	// Adds a zone to a given view, creating it if needed
	AddZoneToView(ctx context.Context, serverID, view, zoneVariant string) (*string, error)

	// Removes the given zone from the given view
	RemoveZoneFromView(ctx context.Context, serverID, view, id string) error
}
