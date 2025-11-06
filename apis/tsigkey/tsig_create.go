package tsigkey

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mittwald/go-powerdns/pdnshttp"
)

// Create adds a TSIG key. If in.Key == "", PDNS may generate it.
func (c *client) Create(ctx context.Context, serverID string, in TSIGKey) (*TSIGKey, error) {
	path := fmt.Sprintf("/servers/%s/tsigkeys", url.PathEscape(serverID))
	var created TSIGKey
	if err := c.httpClient.Post(ctx, path, &created, pdnshttp.WithJSONRequestBody(in)); err != nil {
		return nil, err
	}
	return &created, nil
}
