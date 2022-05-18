package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	p1            Point
	p2            Point
	coveredPoints []Point
}

func (l *Line) PrintLine() {
	fmt.Printf("(%d,%d), (%d,%d)\n", l.p1.x, l.p1.y, l.p2.x, l.p2.y)
}

func (l *Line) IsVertical() bool {
	return (l.p1.x == l.p2.x)
}

func (l *Line) IsHorizontal() bool {
	return (l.p1.y == l.p2.y)
}

func (l *Line) IsDiagnonal() bool {
	return (math.Abs(float64(l.p2.x-l.p1.x)) == math.Abs(float64(l.p2.y-l.p1.y)))
}

func (line *Line) AddHorizontalPoints() {
	starting_y := math.Min(float64(line.p1.y), float64(line.p2.y))
	ending_y := math.Max(float64(line.p1.y), float64(line.p2.y))
	for i := starting_y; i <= ending_y; i++ {
		line.coveredPoints = append(line.coveredPoints, Point{line.p1.x, int(i)})
	}
}

func (line *Line) AddVerticalPoints() {
	starting_x := math.Min(float64(line.p1.x), float64(line.p2.x))
	ending_x := math.Max(float64(line.p1.x), float64(line.p2.x))
	for i := starting_x; i <= ending_x; i++ {
		line.coveredPoints = append(line.coveredPoints, Point{int(i), line.p1.y})
	}
}

func (line *Line) AddDiagonalPoints() {

	x_diff := -1
	y_diff := -1

	if line.p1.x < line.p2.x {
		x_diff = 1
	}

	if line.p1.y < line.p2.y {
		y_diff = 1
	}

	x1 := line.p1.x
	y1 := line.p1.y
	x2 := line.p2.x
	y2 := line.p2.y
	for x1 != x2 && y1 != y2 {
		line.coveredPoints = append(line.coveredPoints, Point{x1, y1})
		x1 += x_diff
		y1 += y_diff
	}
	// Add x2,y2
	line.coveredPoints = append(line.coveredPoints, Point{x2, y2})
}

func (line *Line) CoveredPoints() []Point {

	if !(line.IsVertical() || line.IsHorizontal() || line.IsDiagnonal()) {
		panic(fmt.Sprintf("Invalid line"))
	}

	if line.p1.x == line.p2.x {
		line.AddHorizontalPoints()
	} else if line.p1.y == line.p2.y {
		line.AddVerticalPoints()
	} else {
		line.AddDiagonalPoints()
	}

	return line.coveredPoints
}

func NewLine(x1, y1, x2, y2 int) *Line {
	p1 := Point{x1, y1}
	p2 := Point{x2, y2}
	return &Line{p1, p2, []Point{}}
}

// Parses a line according to this schema: 973,543 -> 601,915
func ParseLine(line string) *Line {

	delim := "->"
	slash := strings.Index(line, delim)

	left := strings.Trim(line[:slash-1], " ")
	right := strings.Trim(line[slash+len(delim):], " ")

	x1, _ := strconv.Atoi(left[:strings.Index(left, ",")])
	y1, _ := strconv.Atoi(left[strings.Index(left, ",")+1:])

	x2, _ := strconv.Atoi(right[:strings.Index(right, ",")])
	y2, _ := strconv.Atoi(right[strings.Index(right, ",")+1:])

	return NewLine(x1, y1, x2, y2)
}

func createLines(inputFile string) []Line {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []Line{}
	for scanner.Scan() {
		lines = append(lines, *ParseLine(scanner.Text()))
	}
	return lines
}

func IsLineHorizontalOrVertical(line Line) bool {
	return line.IsHorizontal() || line.IsVertical()
}

func IsLineHorizontalOrVerticalOrDiagonal(line Line) bool {
	return line.IsHorizontal() || line.IsVertical() || line.IsDiagnonal()
}

func ProcessLines(filter func(line Line) bool, lines []Line) map[Point]int {
	points := make(map[Point]int)
	for _, line := range lines {
		if filter(line) {
			for _, pt := range line.CoveredPoints() {
				points[pt] += 1
			}
		}
	}
	return points
}

func FilterPointsByCoverage(points map[Point]int, threshold int) int {
	count := 0
	for _, v := range points {

		if v >= threshold {
			count++
		}
	}

	return count
}

func main() {
	inputFilePath := flag.String("input", "../input/day_5.txt", "Path of file to be processed")
	lines := createLines(*inputFilePath)
	count := FilterPointsByCoverage(ProcessLines(IsLineHorizontalOrVertical, lines), 2)
	fmt.Printf("Points with >= 2 coverage = %d\n", count)
	count = FilterPointsByCoverage(ProcessLines(IsLineHorizontalOrVerticalOrDiagonal, lines), 2)
	fmt.Printf("Points with >= 2 coverage = %d\n", count)

}
