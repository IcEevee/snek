package main

import (
	"fmt"
	// "github.com/inancgumus/screen"

	"math/rand"
	"strconv"
	"time"

	"github.com/nathan-fiscaletti/consolesize-go"
)

// https://gist.github.com/ik5/d8ecde700972d4378d87
var (
	Red     = "\033[1;31m%s\033[0m"
	Green   = "\033[1;32m%s\033[0m"
	Yellow  = "\033[1;33m%s\033[0m"
	Purple  = "\033[1;34m%s\033[0m"
	Magenta = "\033[1;35m%s\033[0m"
	Teal    = "\033[1;36m%s\033[0m"
	White   = "\033[1;37m%s\033[0m"

	Red_bg     = "\u001b[41m"
	Green_bg   = "\u001b[42m"
	Yellow_bg  = "\u001b[43m"
	Purple_bg  = "\u001b[44m"
	Magenta_bg = "\u001b[45m"
	Teal_bg    = "\u001b[46m"
	White_bg   = "\u001b[47m"

	Reset = "\u001b[0m"
)

var score = 0

// lines, cols
var snake_bodys = [][]int{
	{2, 2},
	{2, 3},
	{2, 4},
}

var red_color = true

var apple_pos = gen_apple_pos()

func print_field() {
	TERM_COLS, TERM_LINES := consolesize.GetConsoleSize()

	for line := 1; line <= TERM_LINES; line++ {
		if line != TERM_LINES {
			for col := 1; col <= TERM_COLS; col++ {
				if line == TERM_LINES-1 {

					// Status bar

					switch col {
					case 1, 2, 3:
						fmt.Print(Green_bg, " ")

					case 4:
						fmt.Print(Green_bg, "S")

					case 5:
						fmt.Print(Green_bg, "N")

					case 6:
						fmt.Print(Green_bg, "E")

					case 7:
						fmt.Print(Green_bg, "K")

					case 8, 9, 10:
						fmt.Print(Green_bg, " ")

					case TERM_COLS, TERM_COLS - 1, TERM_COLS - 2, TERM_COLS - 7, TERM_COLS - 8, TERM_COLS - 9:
						fmt.Print(Green_bg, " ")

					case TERM_COLS - 3:
						fmt.Print(Green_bg, "E")

					case TERM_COLS - 4:
						fmt.Print(Green_bg, "E")

					case TERM_COLS - 5:
						fmt.Print(Green_bg, "c")

					case TERM_COLS - 6:
						fmt.Print(Green_bg, "I")

					default:
						fmt.Print(Reset, " ")

					}

				} else {
					// Snake Body, Apple pos and Field printing
					printing_pos := []int{line, col}

					for _, snake_body := range snake_bodys {
						if snake_body[0] == printing_pos[0] && snake_body[1] == printing_pos[1] {
							if red_color {
								fmt.Print(Red_bg, " ", Reset)
								red_color = false
							} else {
								fmt.Print(Yellow_bg, " ", Reset)
								red_color = true
							}
						}

					}

					if apple_pos[0] == printing_pos[0] && apple_pos[1] == printing_pos[1] {
						fmt.Print(Teal_bg, " ", Reset)
					}

				}

				if col == TERM_COLS {
					fmt.Print("\n")
				}
			}
		} else {
			score_text := "Score: " + strconv.Itoa(score)
			score_text_length := len(score_text)
			fmt.Print(Reset, score_text)
			for i := 1; i <= TERM_COLS-score_text_length; i++ {
				fmt.Print(Reset, " ")
			}
		}
	}
}

func gen_apple_pos() []int {
	rand.Seed(time.Now().UnixNano())

	TERM_COLS, TERM_LINES := consolesize.GetConsoleSize()

	new_apple_pos := []int{
		rand.Intn(TERM_LINES - 2),
		rand.Intn(TERM_COLS),
	}
	for _, snake_body := range snake_bodys {
		if new_apple_pos[0] == snake_body[0] || new_apple_pos[1] == snake_body[1] {
			gen_apple_pos()
		}
	}

	return new_apple_pos
}

func main() {
	print_field()
	fmt.Print(gen_apple_pos())
}
