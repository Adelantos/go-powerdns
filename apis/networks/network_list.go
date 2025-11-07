package networks

import (
	"context"
	"fmt"
	"net/url"
)

func (c *client) ListNetworks(ctx context.Context, serverID string) ([]NetworkView, error) {
	path := fmt.Sprintf("/servers/%s/networks", url.PathEscape(serverID))
	var resp struct {
		Networks []NetworkView `json:"networks"`
	}
	if err := c.httpClient.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if resp.Networks == nil {
		return []NetworkView{}, nil
	}
	return resp.Networks, nil
}
