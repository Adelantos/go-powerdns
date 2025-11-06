package tsigkey

import (
	"context"
	"fmt"
	"net/url"
)

func (c *client) List(ctx context.Context, serverID string) ([]TSIGKey, error) {
	path := fmt.Sprintf("/servers/%s/tsigkeys", url.PathEscape(serverID))
	var out []TSIGKey
	if err := c.httpClient.Get(ctx, path, &out); err != nil {
		return nil, err
	}
	return out, nil // note: key material is not returned here by PDNS
}
