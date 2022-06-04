package main

import (
	"aoc/common"
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func CalculateCost(crabs []int, val int) int {
	var cost = 0
	for _, c := range crabs {
		if c < val {
			cost += val - c
		} else {
			cost += c - val
		}
	}
	return cost
}

func FindMinimum(crabs []int, start int, end int) int {
	var mid = start + ((end - start) / 2)

	if start == end {
		return CalculateCost(crabs, crabs[start])
	}

	cost_at_end := CalculateCost(crabs, crabs[end])
	cost_at_mid := CalculateCost(crabs, crabs[mid])

	if cost_at_mid < cost_at_end {
		return FindMinimum(crabs, start, mid)
	} else {
		return FindMinimum(crabs, mid+1, end)
	}
}

func DaySevenProcessor(line string) {
	var vals = strings.Split(line, ",")
	var crabs = []int{}
	for _, v := range vals {
		k, _ := strconv.Atoi(v)
		crabs = append(crabs, k)
	}

	sort.Ints(crabs)
	min := FindMinimum(crabs, 0, len(crabs)-1)
	fmt.Println(min)
}

func main() {
	inputFilePath := flag.String("input", "../input/day_7.txt", "Path of file to be processed")
	common.ProcessFile(*inputFilePath, DaySevenProcessor)
}
