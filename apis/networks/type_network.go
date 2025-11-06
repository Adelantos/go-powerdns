package networks

type Network struct {
	IP        string `json:"ip"`
	PrefixLen int    `json:"prefixlen"`
	View      string `json:"view"`
}
