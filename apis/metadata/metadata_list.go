package metadata

import (
	"context"
	"fmt"
	"net/url"
)

func (c *client) List(ctx context.Context, serverID, zoneID string) ([]Metadata, error) {
	path := fmt.Sprintf("/servers/%s/zones/%s/metadata",
		url.PathEscape(serverID), url.PathEscape(zoneID))
	var out []Metadata
	if err := c.httpClient.Get(ctx, path, &out); err != nil {
		return nil, err
	}
	return out, nil
}
