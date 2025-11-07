package views

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mittwald/go-powerdns/pdnshttp"
)

func (c *client) AddZoneToView(ctx context.Context, serverID string, view string, zoneVariant string) error {
	path := fmt.Sprintf("/servers/%s/views/%s",
		url.PathEscape(serverID),
		url.PathEscape(view),
	)
	var created string
	body := struct {
		Name string `json:"name"`
	}{Name: zoneVariant}
	return c.httpClient.Post(ctx, path, &created, pdnshttp.WithJSONRequestBody(body))
}
