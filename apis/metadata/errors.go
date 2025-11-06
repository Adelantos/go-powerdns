package metadata

import "errors"

var (
	ErrReadOnlyKind = errors.New("metadata kind is read-only via HTTP metadata endpoint")
	ErrNotViaHTTP   = errors.New("metadata kind is not available via HTTP metadata endpoint; use Zones API or server settings instead")
)
