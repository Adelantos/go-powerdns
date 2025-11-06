package views

import (
	"context"
	"fmt"
	"net/url"
)

func (c *client) RemoveZoneFromView(ctx context.Context, serverID, view, id string) error {
	path := fmt.Sprintf("/servers/%s/views/%s/%s",
		url.PathEscape(serverID),
		url.PathEscape(view),
		url.PathEscape(id),
	)
	return c.httpClient.Delete(ctx, path, nil)
}
