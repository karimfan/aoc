package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type power struct {
	onesCounter [12]int
	diags       int
}

func (p power) commonBits() (string, string) {
	mostCommonBits := ""
	leastCommonBits := ""

	for i := 0; i < 12; i++ {
		if p.onesCounter[i] >= (p.diags - p.onesCounter[i]) {
			mostCommonBits = mostCommonBits + "1"
			leastCommonBits = leastCommonBits + "0"
		} else {
			mostCommonBits = mostCommonBits + "0"
			leastCommonBits = leastCommonBits + "1"
		}
	}

	return mostCommonBits, leastCommonBits
}

func calculateGamma(readings []string) (string, string) {
	power := power{}
	power.diags = 0

	for i := 0; i < len(readings); i++ {

		line := readings[i]
		power.diags++

		for i := 0; i < len(line); i++ {
			char := string(line[i])

			if char == "1" {
				power.onesCounter[i]++
			}
		}
	}

	mostCommon, leastCommon := power.commonBits()

	return mostCommon, leastCommon
}

func filterReadings(readings []string, val string, position int) []string {
	filteredReadings := []string{}

	for i := 0; i < len(readings); i++ {
		reading := readings[i]
		k := string(reading[position])

		if k == val {
			//fmt.Printf("Filtered a reading %s\n", reading)
			filteredReadings = append(filteredReadings, reading)
		}
	}

	return filteredReadings
}

func readReadings(inputFile string) []string {
	readings := []string{}
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		readings = append(readings, line)
	}
	return readings

}

func main() {
	inputFilePath := flag.String("input", "../input/day_3.txt", "Path of file to be processed")

	readings := readReadings(*inputFilePath)

	mostCommon, leastCommon := calculateGamma(readings)
	i, _ := strconv.ParseInt(mostCommon, 2, 64)
	j, _ := strconv.ParseInt(leastCommon, 2, 64)
	fmt.Printf("Most common bits   	%d \n", i)
	fmt.Printf("Least common bits  	%d \n", j)
	fmt.Printf("Power = %d \n", i*j)

	oxygen := readings
	for i := 0; i < len(mostCommon); i++ {
		var k string

		mostCommon, _ := calculateGamma(oxygen)
		k = string(mostCommon[i])

		oxygen = filterReadings(oxygen, k, i)

		if len(oxygen) == 1 {
			break
		}
	}

	carbon := readings
	_, leastCommon = calculateGamma(carbon)
	for i := 0; i < len(leastCommon); i++ {
		var k string
		_, leastCommon = calculateGamma(carbon)
		k = string(leastCommon[i])
		carbon = filterReadings(carbon, k, i)

		if len(carbon) == 1 {
			break
		}
	}

	j, _ = strconv.ParseInt(carbon[0], 2, 64)
	i, _ = strconv.ParseInt(oxygen[0], 2, 64)
	fmt.Printf("oxygen=%d, carbon=%d, life support=%d \n", i, j, i*j)

}
