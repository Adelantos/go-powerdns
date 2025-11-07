package views

import (
	"context"
	"fmt"
	"net/url"
)

func (c *client) ListViewZones(ctx context.Context, serverID string, view string) (*ZoneList, error) {
	path := fmt.Sprintf("/servers/%s/views/%s",
		url.PathEscape(serverID),
		url.PathEscape(view),
	)
	var resp ZoneList

	if err := c.httpClient.Get(ctx, path, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
