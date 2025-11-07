package networks

type Network struct {
	IP        string `json:"ip"`
	PrefixLen int    `json:"prefixlen"`
	View      string `json:"view"`
}

type NetworkView struct {
	Network string `json:"network"`
	View    string `json:"view"`
}
