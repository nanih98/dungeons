package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"text/tabwriter"
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

func (d *Data) PrintTabWriter() {
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', tabwriter.Debug)
	var entry []string
	fmt.Println("Scanned domain:", d.Domain)
	fmt.Fprintln(w, "Nameserver\t Ipv4\t Ipv6")

	for _, nameserver := range d.Nameservers {
		entry = append(entry, nameserver.CNAME, nameserver.IPV4, nameserver.IPV6)
		fmt.Fprintln(w, nameserver.CNAME, "\t", nameserver.IPV4, "\t", nameserver.IPV6)
		entry = nil
	}
	w.Flush()
}

func (d *Data) GetNameservers() []string {
	var nameservers []string
	nameserver, _ := net.LookupNS(d.Domain)
	for _, ns := range nameserver {
		nameservers = append(nameservers, ns.Host)
	}
	return nameservers
}

func (d *Data) AppendNameserverData(server string) {
	serverInfo := Nameservers{
		CNAME: server,
		IPV4:  GetIPV4(server),
		IPV6:  GetIPV6(server),
	}
	d.Nameservers = append(d.Nameservers, serverInfo)
}

func (d *Data) PrintData() {
	marshal, err := json.MarshalIndent(*d, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling json...")
	}
	fmt.Printf("%s", marshal)
}

func main() {
	target := new(Data)

	target.Domain = "google.com"
	nameservers := target.GetNameservers()
	var wg sync.WaitGroup

	for _, server := range nameservers {
		wg.Add(1)
		go func(server string) {
			defer wg.Done()
			target.AppendNameserverData(server)
		}(server)
	}
	wg.Wait()

	// Print the final result (JSON MODE)
	target.PrintData()
	fmt.Println()
	// Print the final result (TABWRITE MODE)
	target.PrintTabWriter()
}
