package networks

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mittwald/go-powerdns/pdnshttp"
	"github.com/stretchr/testify/require"
)

func TestClientListNetworks(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/servers/server-id/networks", r.URL.Path)
		called = true

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{
			"networks": [
				{"network":"192.0.2.0/24","view":"internal"}
			]
		}`))
		require.NoError(t, err)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	out, err := client.ListNetworks(context.Background(), "server-id")
	require.NoError(t, err)
	require.True(t, called)
	//require.Len(t, out, 1)
	require.Equal(t, "192.0.2.0/24", out[0].Network)
}

func TestClientGetNetworkView(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/servers/server-id/networks/198.51.100.0/24", r.URL.Path)
		called = true

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"network":"198.51.100.0/24","view":"internal"}`))
		require.NoError(t, err)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	view, err := client.GetNetworkView(context.Background(), "server-id", "198.51.100.0", 24)
	require.NoError(t, err)
	require.True(t, called)
	require.NotNil(t, view)
	require.Equal(t, "internal", view.View)
	require.Equal(t, "198.51.100.0/24", view.Network)
}

func TestClientSetNetworkView(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		require.Equal(t, "/api/v1/servers/server-id/networks/2001:db8::1/64", r.URL.Path)
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.Equal(t, "{\"view\":\"customer-view\"}\n", string(body))
		called = true

		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	err := client.SetNetworkView(context.Background(), "server-id", "2001:db8::1", 64, "customer-view")
	require.NoError(t, err)
	require.True(t, called)
}
