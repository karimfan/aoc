package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func part_one() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	var count = 0
	var previous_value = math.MaxInt32
	for _, line := range txtlines {
		current_value, err := strconv.Atoi(line)

		if err != nil {
			os.Exit(1)
		}

		if current_value > previous_value {
			count++
		}
		previous_value = current_value
	}
	fmt.Fprintf(os.Stdout, "The number of times the current value was more than previous is %d \n", count)
}

func part_two() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	var slots [3]int
	var slider = 0
	var full_windows = 0
	var windows [2]int
	var window_size = 3
	count := 0
	for _, line := range txtlines {
		int_val, err := strconv.Atoi(line)

		if err != nil {
			os.Exit(1)
		}

		slots[slider] = int_val
		slider++

		// First full window
		if slider == window_size {

			windows[full_windows] = slots[0] + slots[1] + slots[2]

			slots[0] = slots[1]
			slots[1] = slots[2]
			slots[2] = 0
			full_windows++
			slider = window_size - 1
		}

		if full_windows == window_size-1 {
			full_windows--
			if windows[1] > windows[0] {
				count++
			}
			windows[0] = windows[1]
			windows[1] = 0
		}

	}
	fmt.Fprintf(os.Stdout, "The number of times the current window was more than previous is %d \n", count)

}

func main() {
	//part_one()
	part_two()
}
