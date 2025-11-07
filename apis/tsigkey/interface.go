package tsigkey

import "context"

type Client interface {
	// Get all TSIGKeys on the server, except the actual key
	ListTSIGKey(ctx context.Context, serverID string) ([]TSIGKey, error)
	// Add a TSIG key
	CreateTSIGKey(ctx context.Context, serverID string, in TSIGKey) (*TSIGKey, error)
	// Get a specific TSIGKeys on the server, including the actual key
	GetTSIGKey(ctx context.Context, serverID, tsigKeyID string) (*TSIGKey, error)
	// Update the TSIGKey with tsigkey_id
	UpdateTSIGKey(ctx context.Context, serverID, tsigKeyID string, in TSIGKey) (*TSIGKey, error)
	// Delete the TSIGKey with tsigkey_id
	DeleteTSIGKey(ctx context.Context, serverID, tsigKeyID string) error
}
