package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// URL to fetch
var webUrl string = "https://check.torproject.org"

// Specify Tor proxy ip and port
var torProxy string = "socks5://127.0.0.1:9050" // 9150 w/ Tor Browser

func main() {
	// Parse Tor proxy URL string to a URL type
	torProxyURL, err := url.Parse(torProxy)
	if err != nil {
		log.Fatal("Error parsing Tor proxy URL:", torProxy, ".", err)
	}

	// Set up a custom HTTP transport to use the proxy and create the client
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyURL)}
	client := &http.Client{Transport: torTransport, Timeout: time.Second * 5}

	// Make request
	resp, err := client.Get(webUrl)
	//forwarded := resp.Header.Get("X-FORWARDED-FOR")

	fmt.Println(resp.Request.RemoteAddr)

	if err != nil {
		log.Fatal("Error making GET request.", err)
	}

	defer resp.Body.Close()

	// Read response
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal("Error reading body of response.", err)
	// }
	//log.Println(string(body))
	log.Println("Return status code:", resp.StatusCode)
}
