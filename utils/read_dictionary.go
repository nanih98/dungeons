package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/nanih98/dungeons/logger"
)

// Read the dictionary and return a list of string with each entry
func ReadDictionary(log *logger.CustomLogger, dictionary string) []string {
	log.Debug(fmt.Sprintf("Reading subdomains from %s", dictionary))
	var words []string

	dicitonaryExists := checkFileExists(dictionary)

	if !dicitonaryExists {
		log.Fatal(fmt.Errorf("Credentials file %s don't exist", dictionary))
	}

	readFile, err := os.Open(dictionary)

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
