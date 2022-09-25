// package main

// import "fmt"

// type contactInfo struct {
// 	email   string
// 	zipcode string
// }

// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo
// }

// func (p person) print() {
// 	fmt.Printf("%+v", p)
// }

// func (p *person) updateFirstName(name string) {
// 	*&p.firstName = name
// }

// func main() {
// 	dani := person{
// 		firstName: "Daniel",
// 		lastName:  "Cascales",
// 		contactInfo: contactInfo{
// 			email:   "dani@test.com",
// 			zipcode: "08905",
// 		},
// 	}
// 	dani.updateFirstName("hola")
// 	dani.print()
// }

package main

import (
	"fmt"
	"log"
	"net"
)

type Data struct {
	Domain      string
	Nameservers []Nameservers
}

type Nameservers struct {
	CNAME string
	IPV4  string
	IPV6  string
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
		fmt.Println(serverInfo)
	}
}
