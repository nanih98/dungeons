package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Data struct {
	Domain      string        `json:"domain"`
	Nameservers []Nameservers `json:"nameservers"`
}

type Nameservers struct {
	CNAME string `json:"cname"`
	IPV4  string `json:"ipv4"`
	IPV6  string `json:"ipv6"`
}

func GetIPV4(server string) string {
	ip, err := net.LookupIP(server)
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%v", ip[0])
}

func GetIPV6(server string) string {
	ip, err := net.LookupIP(server)
	if err != nil {
		log.Println(err)
	}

	return fmt.Sprintf("%v", ip[1])
}

func (d *Data) GetNameservers() []string {
	var nameservers []string
	nameserver, _ := net.LookupNS(d.Domain)
	for _, ns := range nameserver {
		nameservers = append(nameservers, ns.Host)
	}
	return nameservers
}

func (d *Data) Print() []byte {
	marshal, err := json.MarshalIndent(*d, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling json...")
	}
	return marshal
}

func main() {
	start := new(Data)

	start.Domain = "google.com"
	nameservers := start.GetNameservers()

	for _, server := range nameservers {
		serverInfo := Nameservers{
			CNAME: server,
			IPV4:  GetIPV4(server),
			IPV6:  GetIPV6(server),
		}
		start.Nameservers = append(start.Nameservers, serverInfo)
	}
	fmt.Printf("%s", start.Print())
}
