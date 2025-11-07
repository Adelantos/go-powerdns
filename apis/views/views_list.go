package views

import (
	"context"
	"fmt"
	"net/url"
)

func (c *client) ListViews(ctx context.Context, serverID string) (*ViewsList, error) {
	path := fmt.Sprintf("/servers/%s/views", url.PathEscape(serverID))
	var out ViewsList
	if err := c.httpClient.Get(ctx, path, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
