package metadata

import (
	"context"
	"fmt"
	"net/url"
)

func (c *client) Delete(ctx context.Context, serverID, zoneID, kind string) error {
	if IsNotViaHTTP(kind) {
		return ErrNotViaHTTP
	}
	if IsReadOnlyHTTP(kind) {
		return ErrReadOnlyKind
	}
	path := fmt.Sprintf("/servers/%s/zones/%s/metadata/%s",
		url.PathEscape(serverID), url.PathEscape(zoneID), url.PathEscape(kind))
	return c.httpClient.Delete(ctx, path, nil)
}
