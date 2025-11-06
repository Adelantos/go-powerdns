package networks

import (
	"context"
	"fmt"
	"net/url"
)

func (c *client) ListNetworks(ctx context.Context, serverID string) ([]Network, error) {
	path := fmt.Sprintf("/servers/%s/networks", url.PathEscape(serverID))
	var out []Network
	if err := c.httpClient.Get(ctx, path, &out); err != nil {
		return nil, err
	}
	return out, nil
}
