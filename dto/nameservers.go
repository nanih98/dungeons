package dto

// Data structure of the domain and the nameservers information
type Data struct {
	Domain      string        `json:"domain"`
	Nameservers []Nameservers `json:"nameservers"`
}

// Nameservers struct with the info necessary to operate
type Nameservers struct {
	CNAME string `json:"cname"`
	IPV4  string `json:"ipv4"`
	IPV6  string `json:"ipv6"`
}
