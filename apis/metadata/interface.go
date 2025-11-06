package metadata

import "context"

type Client interface {
	List(ctx context.Context, serverID, zoneID string) ([]Metadata, error)
	Get(ctx context.Context, serverID, zoneID, kind string) (*Metadata, error)
	Create(ctx context.Context, serverID, zoneID string, in Metadata) error
	Replace(ctx context.Context, serverID, zoneID, kind string, in Metadata) (*Metadata, error)
	Delete(ctx context.Context, serverID, zoneID, kind string) error
}
