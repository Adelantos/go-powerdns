package tsigkey

import (
	"context"
	"fmt"
	"net/url"
)

func (c *client) Delete(ctx context.Context, serverID, tsigKeyID string) error {
	path := fmt.Sprintf("/servers/%s/tsigkeys/%s",
		url.PathEscape(serverID),
		url.PathEscape(tsigKeyID),
	)
	return c.httpClient.Delete(ctx, path, nil)
}
