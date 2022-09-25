package utils

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

// Read the dictionary and return a list of string with each entry
func ReadDictionary(path string) []string {
	log.Println("Reading subdomains from", path)
	var words []string

	dicitonaryExists := checkFileExists(path)

	if !dicitonaryExists {
		log.Fatal(fmt.Errorf("Credentials file %s don't exist", path))
	}

	readFile, err := os.Open(path)

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

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}
