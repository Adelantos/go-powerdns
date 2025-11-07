package tsigkey

import (
	"context"
	"fmt"
	"net/url"
)

func (c *client) GetTSIGKey(ctx context.Context, serverID, tsigKeyID string) (*TSIGKey, error) {
	path := fmt.Sprintf("/servers/%s/tsigkeys/%s",
		url.PathEscape(serverID),
		url.PathEscape(tsigKeyID),
	)
	var out TSIGKey
	if err := c.httpClient.Get(ctx, path, &out); err != nil {
		return nil, err
	}
	return &out, nil // this response includes the actual key material
}
