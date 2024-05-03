package main

import (
	"fmt"
)

func sBox(b byte) (byte, error) {
	// build a lookup table 
	// convert the relevant byte to the index the row and column 
	box := [][]byte{
		{0, 2, 1, 3},
		{2, 0, 3, 1},
		{1, 3, 0, 2},
		{3, 1, 2, 0},
	}

	row := (b >> 2) & 3
	column := b & 3
	return box[row][column], nil
}

func main() {
	for i := 0; i <= 16; i++ {
		b := byte(i)
		subbed, err := sBox(b)
		if err != nil {
			fmt.Printf("Error with input %04b: %v\n", i, err)
			continue
		}
		fmt.Printf("%04b -> %02b\n", i, subbed)
	}
}
