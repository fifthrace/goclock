package main

import (
	"fmt"

	clockdigits "github.com/fifthrace/goclock/pkg/clockdigits"
	"github.com/inancgumus/screen"
)

func main() {
	newClock := [8]byte{1, 0, clockdigits.Colon, 2, 3, clockdigits.Colon, 5, 9}

	screen.Clear()

	screen.MoveTopLeft()
	printDigitsLefttoRight(newClock)
}

func printSingleDigit(n byte) {
	d := clockdigits.GetDigit(n)

	for _, b := range d {
		for _, c := range b {
			if c {
				fmt.Print("█")
			} else {
				fmt.Print("░")
			}
		}
		fmt.Println("")
	}
}

func printDigitsLefttoRight(d [8]byte) {
	var clock [8]clockdigits.Digit

	//build my clock array
	for i, digit := range d {
		clock[i] = clockdigits.GetDigit(digit)
	}

	//need to print each line from left to right, so we'll loop over lines then the
	//digits representation per line
	for linenum := range clock[0] {
		for _, digit := range clock[linenum] {
			for _, pixel := range digit {
				if pixel {
					fmt.Print("█")
				} else {
					fmt.Print("░")
				}
			}
		}
		fmt.Println()
	}
}
