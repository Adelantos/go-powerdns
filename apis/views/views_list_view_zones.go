package views

import (
	"context"
	"fmt"
	"net/url"
)

func (c *client) ListViewZones(ctx context.Context, serverID string, view string) ([]string, error) {
	path := fmt.Sprintf("/servers/%s/views/%s",
		url.PathEscape(serverID),
		url.PathEscape(view),
	)
	var out []string
	if err := c.httpClient.Get(ctx, path, &out); err != nil {
		return nil, err
	}
	return out, nil
}
