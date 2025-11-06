package metadata

// Metadata mirrors the PDNS "Metadata" object.
type Metadata struct {
	Kind     string   `json:"kind"`     // e.g. "ALLOW-AXFR-FROM", "SOA-EDIT-API", "X-YourApp-Foo"
	Metadata []string `json:"metadata"` // values for this kind
}
