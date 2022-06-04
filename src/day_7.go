package main

import (
	"aoc/common"
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func CostWithoutDecay(crabs []int, val int) int {
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

func CostWithDecay(crabs []int, val int) int {
	var cost = 0
	var delta = 0
	for _, c := range crabs {
		if c < val {
			delta = val - c
			cost += delta
			cost += GaussSum(delta)
		} else {
			delta = c - val
			cost += delta
			cost += GaussSum(delta)
		}
	}

	return cost
}

func GaussSum(end int) int {
	return (end * (end - 1)) / 2
}

func FindMinimum(crabs []int, positions []int, start int, end int, cost func(crabs []int, val int) int) int {
	var mid = start + ((end - start) / 2)

	cost_at_start := cost(crabs, positions[start])
	cost_at_end := cost(crabs, positions[end])

	if start == end-1 {
		if cost_at_start <= cost_at_end {
			return cost_at_start
		} else {
			return cost_at_end
		}
	}

	if start == end {
		return cost_at_start
	}

	if cost_at_start < cost_at_end {
		return FindMinimum(crabs, positions, start, mid, cost)
	} else {
		return FindMinimum(crabs, positions, mid, end, cost)
	}
}

func DaySevenProcessor(line string) {
	var vals = strings.Split(line, ",")
	var crabs = []int{}
	var positions = []int{}
	for _, v := range vals {
		k, _ := strconv.Atoi(v)
		crabs = append(crabs, k)
	}

	sort.Ints(crabs)
	for i := 0; i <= crabs[len(crabs)-1]; i++ {
		positions = append(positions, i)
	}
	min := FindMinimum(crabs, crabs, 0, len(crabs)-1, CostWithoutDecay)
	fmt.Println(min)

	min = FindMinimum(crabs, positions, 0, len(positions)-1, CostWithDecay)
	fmt.Println(min)
}

func main() {
	inputFilePath := flag.String("input", "../input/day_7.txt", "Path of file to be processed")
	common.ProcessFile(*inputFilePath, DaySevenProcessor)
}
