package common

import (
	"bufio"
	"log"
	"os"
)

func ProcessFile(inputFile string, processor func(line string)) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processor(scanner.Text())
	}
}
