package zones

type ResourceRecordSet struct {
	Name       string              `json:"name"`
	Type       string              `json:"type"`
	TTL        int                 `json:"ttl"`
	ChangeType RecordSetChangeType `json:"changetype,omitempty"`
	Records    []Record            `json:"records"`
	Comments   []Comment           `json:"comments,omitempty"`
}

type Record struct {
	Content    string `json:"content"`
	Disabled   bool   `json:"disabled,omitempty"`
	ModifiedAt int    `json:"modified_at,omitempty"`
}

type Comment struct {
	Content    string `json:"content,omitempty"`
	Account    string `json:"account,omitempty"`
	ModifiedAt int    `json:"modified_at,omitempty"`
}
