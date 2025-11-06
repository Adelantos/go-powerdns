package networks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

func (c *client) GetNetworkView(ctx context.Context, serverID, ip string, prefixLen int) (*string, error) {
	path := fmt.Sprintf("/servers/%s/networks/%s/%s",
		url.PathEscape(serverID),
		url.PathEscape(ip),
		url.PathEscape(strconv.Itoa(prefixLen)),
	)
	var out string
	if err := c.httpClient.Get(ctx, path, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
