package metadata

import (
	"context"
	"fmt"
	"net/url"
)

func (c *client) Get(ctx context.Context, serverID, zoneID, kind string) (*Metadata, error) {
	path := fmt.Sprintf("/servers/%s/zones/%s/metadata/%s",
		url.PathEscape(serverID), url.PathEscape(zoneID), url.PathEscape(kind))
	var out Metadata
	if err := c.httpClient.Get(ctx, path, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
