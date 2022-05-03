package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type submarine struct {
	depth               int
	horizontal_position int
	aim                 int
}

func (sub *submarine) readCourse(line string) {
	// The format is up|down|forward %d
	parts := strings.Split(line, " ")
	movement := strings.ToLower(parts[0])
	magnitude, err := strconv.Atoi(parts[1])

	if err != nil {
		log.Fatal(err)
	}

	switch movement {

	case "up":
		sub.depth -= magnitude

	case "down":
		sub.depth += magnitude

	case "forward":
		sub.horizontal_position += magnitude
	}
}

func (sub *submarine) readCourseWithAim(line string) {
	// The format is up|down|forward %d
	parts := strings.Split(line, " ")
	movement := strings.ToLower(parts[0])
	magnitude, err := strconv.Atoi(parts[1])

	if err != nil {
		log.Fatal(err)
	}

	switch movement {

	case "up":
		sub.aim -= magnitude

	case "down":
		sub.aim += magnitude

	case "forward":
		sub.horizontal_position += magnitude
		sub.depth = sub.depth + (sub.aim * magnitude)
	}
}

func processSubmarineCourse(inputFile string, processWithAim bool) submarine {

	sub := submarine{}
	sub.depth = 0
	sub.horizontal_position = 0
	sub.aim = 0

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if processWithAim == false {
			sub.readCourse(scanner.Text())
		} else {
			sub.readCourseWithAim(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sub
}

func main() {
	inputFilePath := flag.String("input", "../input/day_2.txt", "Path of file to be processed")

	sub := processSubmarineCourse(*inputFilePath, false)
	fmt.Printf("The sub is at a depth of %d and horizontal position of %d \n", sub.depth, sub.horizontal_position)
	fmt.Printf("The sub depth * horizontal position is %d\n", sub.depth*sub.horizontal_position)

	sub = processSubmarineCourse(*inputFilePath, true)
	fmt.Printf("The sub is at a depth of %d, horizontal position of %d and aim %d\n", sub.depth, sub.horizontal_position, sub.aim)
	fmt.Printf("The sub depth * horizontal position is %d\n", sub.depth*sub.horizontal_position)
}
