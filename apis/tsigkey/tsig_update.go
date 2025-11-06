package tsigkey

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mittwald/go-powerdns/pdnshttp"
)

// Update replaces fields of the given TSIG key. PDNS lets you change
// name/algorithm/key; provide only fields you intend to change.
func (c *client) Update(ctx context.Context, serverID, tsigKeyID string, in TSIGKey) (*TSIGKey, error) {
	path := fmt.Sprintf("/servers/%s/tsigkeys/%s",
		url.PathEscape(serverID),
		url.PathEscape(tsigKeyID),
	)
	var updated TSIGKey
	if err := c.httpClient.Put(ctx, path, &updated, pdnshttp.WithJSONRequestBody(in)); err != nil {
		return nil, err
	}
	return &updated, nil
}
