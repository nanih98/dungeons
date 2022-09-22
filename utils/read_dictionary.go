package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile() []string {
	log.Println("Reading subdomains from /usr/local/share/SecLists/Discovery/DNS/subdomains-top1million-5000.txt")
	var words []string
	readFile, err := os.Open("/usr/local/share/SecLists/Discovery/DNS/subdomains-top1million-5000.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
	}
	readFile.Close()

	return words
}
