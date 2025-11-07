package tsigkey

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mittwald/go-powerdns/pdnshttp"
	"github.com/stretchr/testify/require"
)

func TestClientListTSIGKeys(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/servers/localhost/tsigkeys", r.URL.Path)
		called = true

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`[{"id":"1","name":"key-one","algorithm":"hmac-sha256"}]`))
		require.NoError(t, err)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	keys, err := client.List(context.Background(), "localhost")
	require.NoError(t, err)
	require.True(t, called)
	require.Len(t, keys, 1)
	require.Equal(t, "key-one", keys[0].Name)
}

func TestClientGetTSIGKey(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/servers/localhost/tsigkeys/key-one", r.URL.Path)
		called = true

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"id":"key-one","name":"key-one","algorithm":"hmac-sha256","key":"secret"}`))
		require.NoError(t, err)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	key, err := client.Get(context.Background(), "localhost", "key-one")
	require.NoError(t, err)
	require.True(t, called)
	require.NotNil(t, key)
	require.Equal(t, "secret", key.Key)
}

func TestClientCreateTSIGKey(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/api/v1/servers/localhost/tsigkeys", r.URL.Path)
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.JSONEq(t, `{"name":"key-one","algorithm":"hmac-sha256"}`, string(body))
		called = true

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write([]byte(`{"id":"key-one","name":"key-one","algorithm":"hmac-sha256","key":"secret"}`))
		require.NoError(t, err)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	created, err := client.Create(context.Background(), "localhost", TSIGKey{Name: "key-one", Algorithm: "hmac-sha256"})
	require.NoError(t, err)
	require.True(t, called)
	require.NotNil(t, created)
	require.Equal(t, "secret", created.Key)
}

func TestClientUpdateTSIGKey(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		require.Equal(t, "/api/v1/servers/localhost/tsigkeys/key-one", r.URL.Path)
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.JSONEq(t, `{"name":"","algorithm":"","key":"new-secret"}`, string(body))
		called = true

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write([]byte(`{"id":"key-one","name":"key-one","algorithm":"hmac-sha256","key":"new-secret"}`))
		require.NoError(t, err)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	updated, err := client.Update(context.Background(), "localhost", "key-one", TSIGKey{Key: "new-secret"})
	require.NoError(t, err)
	require.True(t, called)
	require.NotNil(t, updated)
	require.Equal(t, "new-secret", updated.Key)
}

func TestClientDeleteTSIGKey(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/api/v1/servers/localhost/tsigkeys/key-one", r.URL.Path)
		called = true

		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	err := client.Delete(context.Background(), "localhost", "key-one")
	require.NoError(t, err)
	require.True(t, called)
}
