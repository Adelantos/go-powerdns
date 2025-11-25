package metadata

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mittwald/go-powerdns/pdnshttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestIsReadOnlyHTTP(t *testing.T) {
	assert.True(t, IsReadOnlyHTTP(string(MDLuaAXFRScript)))
	assert.False(t, IsReadOnlyHTTP(string(MDAllowAXFRFrom)))
	assert.False(t, IsReadOnlyHTTP("CUSTOM"))
}

func TestIsNotViaHTTP(t *testing.T) {
	assert.True(t, IsNotViaHTTP(string(MDApiRectify)))
	assert.False(t, IsNotViaHTTP(string(MDLuaAXFRScript)))
	assert.False(t, IsNotViaHTTP("ALLOW-TSIG"))
}

func TestIsCustomKind(t *testing.T) {
	assert.True(t, IsCustomKind("X-App-Feature"))
	assert.True(t, IsCustomKind("x-other"))
	assert.False(t, IsCustomKind("App-Feature"))
	assert.False(t, IsCustomKind("XY"))
}

func TestClientCreateMetadata(t *testing.T) {
	defer gock.Off()

	gock.New("http://dns.example").
		Post("/api/v1/servers/localhost/zones/example.com./metadata").
		MatchHeader("X-API-Key", "secret").
		Reply(http.StatusNoContent)

	hc := &http.Client{Transport: gock.DefaultTransport}
	c := pdnshttp.NewClient("http://dns.example", hc, &pdnshttp.APIKeyAuthenticator{APIKey: "secret"}, io.Discard)
	mc := New(c)

	err := mc.Create(context.Background(), "localhost", "example.com.", Metadata{Kind: string(MDAllowAXFRFrom), Metadata: []string{"192.0.2.1"}})

	assert.NoError(t, err)
	assert.True(t, gock.IsDone())
}

func TestClientCreateMetadataRejectsReadOnly(t *testing.T) {
	mc := &client{}

	err := mc.Create(context.Background(), "localhost", "example.com.", Metadata{Kind: string(MDLuaAXFRScript)})

	assert.Equal(t, ErrReadOnlyKind, err)
}

func TestClientReplaceMetadata(t *testing.T) {
	defer gock.Off()

	gock.New("http://dns.example").
		Put("/api/v1/servers/localhost/zones/example.com./metadata/ALLOW-AXFR-FROM").
		MatchHeader("X-API-Key", "secret").
		Reply(http.StatusOK).
		SetHeader("Content-Type", "application/json").
		JSON(map[string]interface{}{
			"kind":     string(MDAllowAXFRFrom),
			"metadata": []string{"198.51.100.1"},
		})

	hc := &http.Client{Transport: gock.DefaultTransport}
	c := pdnshttp.NewClient("http://dns.example", hc, &pdnshttp.APIKeyAuthenticator{APIKey: "secret"}, io.Discard)
	mc := New(c)

	out, err := mc.Replace(context.Background(), "localhost", "example.com.", string(MDAllowAXFRFrom), Metadata{Metadata: []string{"198.51.100.1"}})

	assert.NoError(t, err)
	if assert.NotNil(t, out) {
		assert.Equal(t, string(MDAllowAXFRFrom), out.Kind)
		assert.Equal(t, []string{"198.51.100.1"}, out.Metadata)
	}
	assert.True(t, gock.IsDone())
}

func TestClientReplaceMetadataRejectsNotViaHTTP(t *testing.T) {
	mc := &client{}

	out, err := mc.Replace(context.Background(), "localhost", "example.com.", string(MDApiRectify), Metadata{})

	assert.Nil(t, out)
	assert.Equal(t, ErrNotViaHTTP, err)
}

func TestClientDeleteMetadata(t *testing.T) {
	defer gock.Off()

	gock.New("http://dns.example").
		Delete("/api/v1/servers/localhost/zones/example.com./metadata/ALLOW-AXFR-FROM").
		MatchHeader("X-API-Key", "secret").
		Reply(http.StatusNoContent)

	hc := &http.Client{Transport: gock.DefaultTransport}
	c := pdnshttp.NewClient("http://dns.example", hc, &pdnshttp.APIKeyAuthenticator{APIKey: "secret"}, io.Discard)
	mc := New(c)

	err := mc.Delete(context.Background(), "localhost", "example.com.", string(MDAllowAXFRFrom))

	assert.NoError(t, err)
	assert.True(t, gock.IsDone())
}

func TestClientDeleteMetadataRejectsReadOnly(t *testing.T) {
	mc := &client{}

	err := mc.Delete(context.Background(), "localhost", "example.com.", string(MDLuaAXFRScript))

	assert.Equal(t, ErrReadOnlyKind, err)
}

func TestClientListMetadata(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/servers/example-server/zones/example.net./metadata", r.URL.Path)
		called = true

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`[
                        {"kind":"ALLOW-AXFR-FROM","metadata":["192.0.2.1"]}
                ]`))
		require.NoError(t, err)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	out, err := client.List(context.Background(), "example-server", "example.net.")
	require.NoError(t, err)
	require.True(t, called)
	require.Len(t, out, 1)
	require.Equal(t, "ALLOW-AXFR-FROM", out[0].Kind)
	require.Equal(t, []string{"192.0.2.1"}, out[0].Metadata)
}

func TestClientGetMetadata(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/servers/example-server/zones/example.net./metadata/ALLOW-AXFR-FROM", r.URL.Path)
		called = true

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"kind":"ALLOW-AXFR-FROM","metadata":["192.0.2.1"]}`))
		require.NoError(t, err)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	out, err := client.Get(context.Background(), "example-server", "example.net.", string(MDAllowAXFRFrom))
	require.NoError(t, err)
	require.True(t, called)
	require.NotNil(t, out)
	require.Equal(t, "ALLOW-AXFR-FROM", out.Kind)
	require.Equal(t, []string{"192.0.2.1"}, out.Metadata)
}

func TestClientCreateMetadataReadOnlyKind(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatalf("unexpected call: %s %s", r.Method, r.URL.Path)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	err := client.Create(context.Background(), "example-server", "example.net.", Metadata{
		Kind: string(MDLuaAXFRScript),
	})
	require.Equal(t, ErrReadOnlyKind, err)
}

func TestClientReplaceMetadataGuardrails(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatalf("unexpected call: %s %s", r.Method, r.URL.Path)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	_, err := client.Replace(context.Background(), "example-server", "example.net.", string(MDLuaAXFRScript), Metadata{})
	require.Equal(t, ErrReadOnlyKind, err)

	_, err = client.Replace(context.Background(), "example-server", "example.net.", string(MDApiRectify), Metadata{})
	require.Equal(t, ErrNotViaHTTP, err)
}

func TestClientDeleteMetadataGuardrails(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatalf("unexpected call: %s %s", r.Method, r.URL.Path)
	}))
	defer srv.Close()

	c := pdnshttp.NewClient(srv.URL, srv.Client(), nil, io.Discard)
	client := New(c)

	err := client.Delete(context.Background(), "example-server", "example.net.", string(MDLuaAXFRScript))
	require.Equal(t, ErrReadOnlyKind, err)

	err = client.Delete(context.Background(), "example-server", "example.net.", string(MDApiRectify))
	require.Equal(t, ErrNotViaHTTP, err)
}
