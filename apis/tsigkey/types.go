package tsigkey

// TSIGKey mirrors the PDNS TSIGKey schema.
// PDNS commonly uses "id" as the resource identifier; "Name" is the TSIG key name.
// On list calls, PDNS omits the actual key material; on single GET it is present.
type TSIGKey struct {
	ID        string `json:"id,omitempty"`   // server-side identifier (often same as Name)
	Name      string `json:"name"`           // TSIG key name
	Algorithm string `json:"algorithm"`      // e.g. hmac-sha256
	Key       string `json:"key,omitempty"`  // base64 secret; omitted in list responses
	Type      string `json:"type,omitempty"` // usually "TSIGKey"
}
