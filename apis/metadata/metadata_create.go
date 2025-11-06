package metadata

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mittwald/go-powerdns/pdnshttp"
)

// Create adds values for a kind; existing values of the same kind are not overwritten.
// Returns 204 No Content on success.
func (c *client) Create(ctx context.Context, serverID, zoneID string, in Metadata) error {
	if IsNotViaHTTP(in.Kind) {
		return ErrNotViaHTTP
	}
	if IsReadOnlyHTTP(in.Kind) {
		return ErrReadOnlyKind
	}
	path := fmt.Sprintf("/servers/%s/zones/%s/metadata",
		url.PathEscape(serverID), url.PathEscape(zoneID))
	return c.httpClient.Post(ctx, path, nil, pdnshttp.WithJSONRequestBody(in))
}
