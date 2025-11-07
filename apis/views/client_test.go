package views

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mittwald/go-powerdns/pdnshttp"
	"github.com/stretchr/testify/require"
)

func TestClientListViews(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/servers/localhost/views", r.URL.Path)
		called = true

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"views":["internal","external"]}`))
		require.NoError(t, err)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	out, err := client.ListViews(context.Background(), "localhost")
	require.NoError(t, err)
	require.True(t, called)
	require.NotNil(t, out)
	require.Equal(t, []string{"internal", "external"}, out.Views)
}

func TestClientListViewZones(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/servers/localhost/views/internal", r.URL.Path)
		called = true

		w.Header().Set("Content-Type", "application/json")
		// Match the client's expected shape: {"zones":[...]}
		_, err := w.Write([]byte(`{"zones":["example.org","example.net"]}`))
		require.NoError(t, err)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	zones, err := client.ListViewZones(context.Background(), "localhost", "internal")
	require.NoError(t, err)
	require.True(t, called)
	require.Equal(t, []string{"example.org", "example.net"}, zones.Zones)
}

func TestClientAddZoneToView(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/api/v1/servers/localhost/views/internal", r.URL.Path)
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.Equal(t, "{\"name\":\"example.org\"}\n", string(body))
		called = true

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write([]byte(`"example.org"`))
		require.NoError(t, err)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	err := client.AddZoneToView(context.Background(), "localhost", "internal", "example.org")
	require.NoError(t, err)
	require.True(t, called)
	//require.NotNil(t, zoneID)
	//require.Equal(t, "example.org", *zoneID)
}

func TestClientRemoveZoneFromView(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/api/v1/servers/localhost/views/internal/example.org", r.URL.Path)
		called = true

		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	err := client.RemoveZoneFromView(context.Background(), "localhost", "internal", "example.org")
	require.NoError(t, err)
	require.True(t, called)
}
