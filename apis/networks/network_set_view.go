package networks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/mittwald/go-powerdns/pdnshttp"
)

func (c *client) SetNetworkView(ctx context.Context, serverID, ip string, prefixLen int, view string) error {
	path := fmt.Sprintf("/servers/%s/networks/%s/%s",
		url.PathEscape(serverID),
		url.PathEscape(ip),
		url.PathEscape(strconv.Itoa(prefixLen)),
	)
	// WithJSONRequestBody(view) will serialize to "viewname" on the wire.
	return c.httpClient.Put(ctx, path, nil, pdnshttp.WithJSONRequestBody(view))
}
