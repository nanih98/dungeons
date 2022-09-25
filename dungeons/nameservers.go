package dungeons

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"text/tabwriter"

	"github.com/nanih98/dungeons/dto"
)

// // Data structure of the domain and the nameservers information
// type Data struct {
// 	Domain      string        `json:"domain"`
// 	Nameservers []Nameservers `json:"nameservers"`
// }

// // Nameservers struct with the info necessary to operate
// type Nameservers struct {
// 	CNAME string `json:"cname"`
// 	IPV4  string `json:"ipv4"`
// 	IPV6  string `json:"ipv6"`
// }

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

func (d *dto.Data) PrintTabWriter() {
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

func (d *dto.Data) GetNameservers() []string {
	var nameservers []string
	nameserver, _ := net.LookupNS(d.Domain)
	for _, ns := range nameserver {
		nameservers = append(nameservers, ns.Host)
	}
	return nameservers
}

func (d *dto.Data) AppendNameserverData(server string) {
	serverInfo := dto.Nameservers{
		CNAME: server,
		IPV4:  GetIPV4(server),
		IPV6:  GetIPV6(server),
	}
	d.Nameservers = append(d.Nameservers, serverInfo)
}

func (d *dto.Data) PrintJson() {
	marshal, err := json.MarshalIndent(*d, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling json...")
	}
	fmt.Printf("%s", marshal)
}

func Info(domain string, outputmode string) {
	target := new(dto.Data)

	target.Domain = domain
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

	switch outputmode {
	case "tabwriter":
		target.PrintTabWriter()
	case "json":
		target.PrintJson()
	default:
		log.Fatal("Outputmode incorrect. Use -output json|tabwriter")
	}
}
