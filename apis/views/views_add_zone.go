package views

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mittwald/go-powerdns/pdnshttp"
)

func (c *client) AddZoneToView(ctx context.Context, serverID string, view string, zoneVariant string) (*string, error) {
	path := fmt.Sprintf("/servers/%s/views/%s",
		url.PathEscape(serverID),
		url.PathEscape(view),
	)
	var created string
	if err := c.httpClient.Post(ctx, path, &created, pdnshttp.WithJSONRequestBody(zoneVariant)); err != nil {
		return nil, err
	}
	return &created, nil
}
