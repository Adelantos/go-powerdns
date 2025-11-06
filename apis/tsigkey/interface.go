package tsigkey

import "context"

type Client interface {
	// Get all TSIGKeys on the server, except the actual key
	List(ctx context.Context, serverID string) ([]TSIGKey, error)
	// Add a TSIG key
	Create(ctx context.Context, serverID string, in TSIGKey) (*TSIGKey, error)
	// Get a specific TSIGKeys on the server, including the actual key
	Get(ctx context.Context, serverID, tsigKeyID string) (*TSIGKey, error)
	// Update the TSIGKey with tsigkey_id
	Update(ctx context.Context, serverID, tsigKeyID string, in TSIGKey) (*TSIGKey, error)
	// Delete the TSIGKey with tsigkey_id
	Delete(ctx context.Context, serverID, tsigKeyID string) error
}
