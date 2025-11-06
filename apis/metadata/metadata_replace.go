package metadata

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mittwald/go-powerdns/pdnshttp"
)

// Replace replaces the entire value-set for 'kind' and returns the resulting object (200 OK).
func (c *client) Replace(ctx context.Context, serverID, zoneID, kind string, in Metadata) (*Metadata, error) {
	if IsNotViaHTTP(kind) {
		return nil, ErrNotViaHTTP
	}
	if IsReadOnlyHTTP(kind) {
		return nil, ErrReadOnlyKind
	}
	path := fmt.Sprintf("/servers/%s/zones/%s/metadata/%s",
		url.PathEscape(serverID), url.PathEscape(zoneID), url.PathEscape(kind))
	var out Metadata
	if err := c.httpClient.Put(ctx, path, &out, pdnshttp.WithJSONRequestBody(in)); err != nil {
		return nil, err
	}
	return &out, nil
}
