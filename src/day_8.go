package main

import (
	"aoc/common"
	"flag"
	"fmt"
	"math"
	"sort"
	"strings"
)

var count = 0
var decodedSum = 0.0

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func FindByLength(encoded []string, length int) []string {
	retVal := make([]string, 0)

	for _, v := range encoded {
		if len(v) == length {
			retVal = append(retVal, v)
		}
	}
	return retVal
}

func FindIntersection(encoded []string, target string, coverage int) []string {
	retVal := make([]string, 0)
	for _, v := range encoded {

		if DoesIntersect(v, target, coverage) {
			retVal = append(retVal, v)
		}
	}
	return retVal
}

func DoesIntersect(encoded string, target string, coverage int) bool {
	retVal := false
	match := 0

	for _, j := range target {
		if strings.Contains(encoded, string(j)) {
			match++
		}
	}
	if match == coverage {
		retVal = true
	}
	return retVal
}

func DayEightPartTwo(line string) {
	var vals = strings.Split(line, "|")
	var input = strings.Split(vals[0], " ")
	var output = strings.Split(vals[1], " ")[1:]
	var exponent = 3.0
	var number = 0.0

	_, decoder := CreateDecoder(input)
	for _, v := range output {
		v = SortString(v)
		number += math.Pow(10, exponent) * float64(decoder[v])
		exponent--
	}
	decodedSum += number
}

func CreateDecoder(input []string) (map[int]string, map[string]int) {
	var tostring = make(map[int]string)
	var toint = make(map[string]int)

	tostring[1] = SortString(FindByLength(input, 2)[0])
	tostring[4] = SortString(FindByLength(input, 4)[0])
	tostring[7] = SortString(FindByLength(input, 3)[0])
	tostring[8] = SortString(FindByLength(input, 7)[0])
	tostring[3] = SortString(FindIntersection(FindByLength(input, 5), tostring[1], len(tostring[1]))[0])
	tostring[9] = SortString(FindIntersection(FindByLength(input, 6), tostring[4], len(tostring[4]))[0])
	tostring[5] = SortString(FindIntersection(FindByLength(input, 5), tostring[4], 3)[0])

	for _, v := range FindByLength(input, 5) {
		v = SortString(v)
		if v != tostring[3] {
			if DoesIntersect(v, tostring[4], 3) {
				tostring[5] = v
			} else {
				tostring[2] = v
			}
		}
	}

	for _, v := range FindByLength(input, 6) {
		if SortString(v) != tostring[9] {
			if DoesIntersect(v, tostring[1], 2) {
				tostring[0] = SortString(v)
			} else {
				tostring[6] = SortString(v)
			}
		}
	}

	for i, v := range tostring {
		toint[v] = i
	}

	return tostring, toint
}

// line format is :
// cdbga acbde eacdfbg adbgf gdebcf bcg decabf cg ebdgac egca | geac ceag faedcb cg
func DayEightPartOne(line string) {
	var vals = strings.Split(line, "|")
	var output = strings.Split(vals[1], " ")
	for _, v := range output {
		if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
			count++
		}
	}
}

func main() {
	inputFilePath := flag.String("input", "../input/day_8.txt", "Path of file to be processed")
	common.ProcessFile(*inputFilePath, DayEightPartOne)
	fmt.Printf("Count of 1, 4, 7 & 8 is %d \n", count)
	common.ProcessFile(*inputFilePath, DayEightPartTwo)
	fmt.Printf("Sum of decoded output is %f\n", decodedSum)
}
