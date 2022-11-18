package utils

import (
	"math/rand"
)

// RandomServer return a random server of a given list of string Ex: [ns1.example.com,ns2.example.com,ns3.example.com]
func RandomServer(servers []string) string {
	randomIndex := rand.Intn(len(servers))
	pick := servers[randomIndex]
	return pick
}
