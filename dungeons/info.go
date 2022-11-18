package dungeons

// // Data structure of the domain and the nameservers information
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

//func GetIPV4(server string, log *logger.CustomLogger) string {
//	ip, err := net.LookupIP(server)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return fmt.Sprintf("%v", ip[0])
//}
//
//func GetIPV6(server string, log *logger.CustomLogger) string {
//	ip, err := net.LookupIP(server)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return fmt.Sprintf("%v", ip[1])
//}
//
//// GetNameservers returns the list of all the nameservers of the given domain
//func GetNameservers(domain string) []string {
//	var nameservers []string
//	nameserver, _ := net.LookupNS(domain)
//	for _, ns := range nameserver {
//		nameservers = append(nameservers, ns.Host)
//	}
//	return nameservers
//}

//func (d *Data) PrintTabWriter() {
//	w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', tabwriter.Debug)
//	fmt.Fprintln(w, "Nameservers\t Ipv4\t Ipv6")
//
//	for _, nameserver := range d.Nameservers {
//		fmt.Fprintln(w, nameserver.CNAME, "\t", nameserver.IPV4, "\t", nameserver.IPV6)
//	}
//	w.Flush()
//}

//func (d *Data) AppendNameserverData(server string, log *logger.CustomLogger) {
//	serverInfo := Nameservers{
//		CNAME: server,
//		IPV4:  GetIPV4(server, log),
//		IPV6:  GetIPV6(server, log),
//	}
//	d.Nameservers = append(d.Nameservers, serverInfo)
//}
//
//func (d *Data) PrintJson(log *logger.CustomLogger) {
//	marshal, err := json.MarshalIndent(*d, "", "  ")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%s", marshal)
//}
//
//func Info(log *logger.CustomLogger, domain, outputmode string) {
//	target := new(Data)
//	nameservers := GetNameservers(domain)
//	var wg sync.WaitGroup
//
//	for _, server := range nameservers {
//		wg.Add(1)
//		go func(server string) {
//			defer wg.Done()
//			target.AppendNameserverData(server, log)
//		}(server)
//	}
//	wg.Wait()
//
//	switch outputmode {
//	case "tabwriter":
//		target.PrintTabWriter()
//	case "json":
//		target.PrintJson(log)
//	default:
//		log.Fatal(fmt.Errorf("Outputmode incorrect. Use -output json|tabwriter"))
//	}
//}
