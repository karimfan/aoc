package main

import (
	"aoc/common"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func PrintFish(fish []int) {
	for i, _ := range fish {
		fmt.Printf("%d, ", fish[i])
	}
	fmt.Printf("\n")
}

func SwapFish(value int, index int, fish []int) int {
	temp := fish[index]
	fish[index] = value
	return temp
}

func LineProcessor(line string) {
	temp := strings.Split(line, ",")
	var fish = make([]int, 9)
	for j := range temp {
		k, _ := strconv.Atoi(temp[j])
		fish[k] += 1
	}
	for j := 1; j <= 256; j++ {
		new_fish := fish[0]
		i := fish[8]
		for k := len(fish) - 2; k >= 0; k-- {
			i = SwapFish(i, k, fish)
		}

		fish[6] += new_fish
		fish[8] = new_fish
	}
	count := 0
	for i, _ := range fish {
		count += fish[i]
	}
	fmt.Printf("\n")
	fmt.Printf("count of fish %d\n", count)
}

func main() {
	inputFilePath := flag.String("input", "../input/day_6.txt", "Path of file to be processed")
	common.ProcessFile(*inputFilePath, LineProcessor)
}
