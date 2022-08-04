package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// size of the board, always a square, ex: 4 means 4x4
const BOARD_SIZE int = 4

// store the information for a tile in a simple struct
// lmao i'm not using it right now, maybe if i need to store other info later?

// type tile struct {
// 	x, y, val int
// }

func main() {
	// randomize the seed because it isn't random by default
	seed := time.Now().Unix()
	rand.Seed(seed)

	fmt.Println("Welcome to GO 2048!")

	game_board := [BOARD_SIZE][BOARD_SIZE]int{}

	// game loop omg like python games course 😱
	for {
		if add_tile(&game_board) == false {
			fmt.Printf("\nyou lost dumbass 🤡\n")
			return
		}

		printBoard(&game_board)

		input := get_input()
		move(input, &game_board)
	}
}

func printBoard(b *[BOARD_SIZE][BOARD_SIZE]int) {
	// prints the board i dont know what else to say
	fmt.Printf("\n")
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			fmt.Printf("%d ", b[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func get_input() string {
	// gets input from the user, either u, d, l, or r
	var user_input string

	for !is_valid(&user_input) {
		fmt.Printf("Enter an input (u/d/l/r): ")
		fmt.Scanf("%s", &user_input)
	}

	return user_input
}

func is_valid(s *string) bool {
	// takes in a string s and checks if it is a valid input, returns a bool
	// valid inputs start with the letter u, d, l, or r (case insensitive)
	valid_inputs := [4]string{"u", "d", "l", "r"}

	if len(*s) == 0 {
		fmt.Println("Please enter an input")
		return false
	}
	*s = strings.ToLower(*s)
	*s = (*s)[:1]
	for i := 0; i < len(valid_inputs); i++ {
		if *s == valid_inputs[i] {
			return true
		}
	}
	fmt.Println("Invalid input entered")
	return false
}

func add_tile(b *[BOARD_SIZE][BOARD_SIZE]int) bool {
	// takes in the board and returns true if it adds a tile successfully
	// returns false if unsuccessful aka the board is full
	var val int

	// create a slice of all of the empty tiles (tiles with a value of 0)
	empty_tiles := make([]*int, 0, 16)
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if (*b)[i][j] == 0 {
				empty_tiles = append(empty_tiles, &(*b)[i][j])
			}
		}
	}

	if len(empty_tiles) == 0 {
		return false
	}

	// pick a random tile from the empty tiles
	tile := empty_tiles[rand.Intn(len(empty_tiles))]

	// pick a value, mostly 2, 10% chance it is a 4
	percent := rand.Intn(10)
	if percent == 0 {
		val = 4
	} else {
		val = 2
	}

	// assign the value to the tile
	*tile = val

	return true
}

// move code below to separate file eventually, or not idk
func move(input string, b *[BOARD_SIZE][BOARD_SIZE]int) {
	// how tf am i going to make the tiles move 😭
	switch input {
	case "u":
		fmt.Printf("TODO LMAO\n")
	case "d":
		fmt.Printf("TODO LMAO\n")
	case "l":
		move_left(b)
	case "r":
		move_right(b)
	}
}

func move_left(b *[BOARD_SIZE][BOARD_SIZE]int) {
	// don't ask me how it works cuz idk
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if (*b)[i][j] != 0 {
				for pos := j; pos != 0; pos-- {
					curr_tile := &(*b)[i][pos]
					dest_tile := &(*b)[i][pos-1]

					if *dest_tile == 0 { // slide
						*dest_tile = *curr_tile
						*curr_tile = 0
					} else if *dest_tile == *curr_tile { // merge
						*dest_tile = *curr_tile * 2
						*curr_tile = 0
					}
				}
			}
		}
	}
}

func move_right(b *[BOARD_SIZE][BOARD_SIZE]int) {
	// don't ask me how it works cuz idk
	for i := 0; i < BOARD_SIZE; i++ {
		for j := BOARD_SIZE - 1; j >= 0; j-- {
			if (*b)[i][j] != 0 {
				for pos := j; pos != BOARD_SIZE-1; pos++ {
					curr_tile := &(*b)[i][pos]
					dest_tile := &(*b)[i][pos+1]

					if *dest_tile == 0 { // slide
						*dest_tile = *curr_tile
						*curr_tile = 0
					} else if *dest_tile == *curr_tile { // merge
						*dest_tile = *curr_tile * 2
						*curr_tile = 0
					}
				}
			}
		}
	}
}
