/*
* read the input and for every game create a 6x6 grid
* the top row and right most column store the scores for the rows & columns. default to 0 and incremented once
* a cell on the intersecting row + column is selected.
* bingo when one of the cells in the topmost row or rightmost column are 5
 */
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

const grid_width = 5
const grid_length = 5

type mode int

const (
	FirstWins mode = iota
	LastWins
)

type grid_item struct {
	marked bool
	value  int64
}

type board struct {
	grid             [grid_width][grid_length]grid_item
	horizontal_score [grid_length]int
	vertical_score   [grid_length]int
}

func (b *board) printBoard() {
	fmt.Printf("------------BOARD BEGIN----------------\n")
	for i := 0; i < grid_width; i++ {
		for j := 0; j < grid_length; j++ {
			fmt.Printf("%d=%t,", b.grid[i][j].value, b.grid[i][j].marked)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("----------------BOARD END----------------\n")

}

// row can look like so "36 33 18 54 10" Can have multiple spaces between values
func (b *board) addRow(row string, row_number int) {
	var values = strings.Split(row, " ")
	// There could be empty strings in the row. Omit them
	index := 0
	for j := 0; j < len(values); j++ {
		val := values[j]
		if val != "" {
			b.grid[row_number][index].value, _ = strconv.ParseInt(val, 10, 64)
			index++
		}
	}
}

func (b *board) sumUnmarkedCells() int64 {
	var val int64 = 0
	for i := 0; i < grid_width; i++ {
		for j := 0; j < grid_length; j++ {
			cell := b.grid[i][j]
			if cell.marked == false {
				val += cell.value
			}
		}
	}
	return val
}

// return true if bingo, false otherwise
func (b *board) playMove(number int64) (bool, int64) {
	var bingo bool = false
	var val int64 = 0

	for i := 0; i < grid_length; i++ {
		for j := 0; j < grid_width; j++ {
			if b.grid[i][j].value == number {
				b.grid[i][j].marked = true
				b.horizontal_score[i]++
				b.vertical_score[j]++

				if b.vertical_score[j] == grid_width || b.horizontal_score[i] == grid_length {
					val = b.sumUnmarkedCells()
					val *= number
					bingo = true
					break
				}
			}
		}
	}

	return bingo, val
}

// comma delimited like so: 91,17,64,45,8,13,
func readGameInput(line string) []string {
	return strings.Split(line, ",")
}

func playTurns(game []string, boards []board, m mode) (board, int64) {
	var last_winning_board board
	var last_winning_val int64 = 10

	for i := 0; i < len(game); i++ {
		number, _ := strconv.ParseInt(game[i], 10, 64)
		for e := 0; e < len(boards); e++ {
			b := &boards[e]
			bingo, val := b.playMove(number)

			if bingo == true {
				last_winning_board = *b
				last_winning_val = val
				if m == FirstWins {
					return last_winning_board, last_winning_val
				} else {
					// reset to continue finding the last board
					boards = append(boards[:e], boards[e+1:]...)
					e--
				}
			}

		}
	}
	return last_winning_board, last_winning_val
}

func playGame(inputFile string, m mode) (board, int64) {
	var game []string
	var boards = []board{}

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	readOrder := false
	row_number := 0
	current_board := board{}

	for scanner.Scan() {

		line := scanner.Text()

		if readOrder == false {
			readOrder = true
			game = readGameInput(line)
		} else {

			if line == "" {
				// do nothing. This is the delimiter between boards
			} else {
				current_board.addRow(line, row_number)
				row_number++
			}

			if row_number == grid_length {
				boards = append(boards, current_board)
				current_board = board{}
				row_number = 0
			}
		}
	}
	return playTurns(game, boards, m)

}

func main() {
	inputFilePath := flag.String("input", "../input/day_4.txt", "Path of file to be processed")

	b, v := playGame(*inputFilePath, FirstWins)
	b.printBoard()
	fmt.Printf("First winning board val is %d\n", v)

	b, v = playGame(*inputFilePath, LastWins)
	b.printBoard()
	fmt.Printf("Last winning board val is %d\n", v)
}
