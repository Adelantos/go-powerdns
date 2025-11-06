package metadata

// MetadataKind is a string alias; callers can still pass any string.
type MetadataKind string

// Writable kinds commonly used via the metadata endpoint (non-exhaustive).
const (
	MDAllowAXFRFrom        MetadataKind = "ALLOW-AXFR-FROM"
	MDDNSUpdateFrom        MetadataKind = "ALLOW-DNSUPDATE-FROM"
	MDAlsoNotify           MetadataKind = "ALSO-NOTIFY"
	MDFwdDNSUpdate         MetadataKind = "FORWARD-DNSUPDATE"
	MDNotifyDNSUpdate      MetadataKind = "NOTIFY-DNSUPDATE"
	MDIXFR                 MetadataKind = "IXFR"
	MDPublishCDNSKEY       MetadataKind = "PUBLISH-CDNSKEY"
	MDPublishCDS           MetadataKind = "PUBLISH-CDS"
	MDSlaveRenotify        MetadataKind = "SLAVE-RENOTIFY"
	MDTSIGAllowAXFR        MetadataKind = "TSIG-ALLOW-AXFR"
	MDTSIGAllowDNSUpdate   MetadataKind = "TSIG-ALLOW-DNSUPDATE"
	MDGSSAcceptorPrincipal MetadataKind = "GSS-ACCEPTOR-PRINCIPAL"
	MDGSSAllowAXFRPrin     MetadataKind = "GSS-ALLOW-AXFR-PRINCIPAL"
)

// GET-only at the metadata endpoint (writes forbidden).
const (
	MDAXFRMasterTSIG MetadataKind = "AXFR-MASTER-TSIG"
	MDLuaAXFRScript  MetadataKind = "LUA-AXFR-SCRIPT"
	MDNSEC3Narrow    MetadataKind = "NSEC3NARROW"
	MDNSEC3Param     MetadataKind = "NSEC3PARAM"
	MDPreSigned      MetadataKind = "PRESIGNED"
	MDSOAEdit        MetadataKind = "SOA-EDIT"
)

// Not exposed via metadata endpoint (use Zones API/server settings instead).
const (
	MDApiRectify       MetadataKind = "API-RECTIFY"
	MDEnableLuaRecords MetadataKind = "ENABLE-LUA-RECORDS"
	MDSOAEditAPI       MetadataKind = "SOA-EDIT-API"
)

// Advisory sets for client-side guardrails.
var readOnlyHTTP = map[MetadataKind]struct{}{
	MDAXFRMasterTSIG: {}, MDLuaAXFRScript: {}, MDNSEC3Narrow: {},
	MDNSEC3Param: {}, MDPreSigned: {}, MDSOAEdit: {},
}

var notViaHTTP = map[MetadataKind]struct{}{
	MDApiRectify: {}, MDEnableLuaRecords: {}, MDSOAEditAPI: {},
}

func IsReadOnlyHTTP(kind string) bool { _, ok := readOnlyHTTP[MetadataKind(kind)]; return ok }
func IsNotViaHTTP(kind string) bool   { _, ok := notViaHTTP[MetadataKind(kind)]; return ok }

// IsCustomKind returns true if the kind follows the X- prefix rule for custom metadata.
func IsCustomKind(kind string) bool {
	return len(kind) > 2 && (kind[0] == 'X' || kind[0] == 'x') && kind[1] == '-'
}
