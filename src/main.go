package main

import (
	"fmt"
	"time"

	clockdigits "github.com/fifthrace/goclock/pkg/clockdigits"
	"github.com/inancgumus/screen"
)

func main() {

	screen.Clear()

	for {
		screen.Clear()
		screen.MoveTopLeft()
		printDigitsLefttoRight(getClockArray(time.Now()))

		time.Sleep(time.Second)
	}
}

func getClockArray(t time.Time) [8]byte {
	var ret [8]byte

	h := t.Hour()
	m := t.Minute()
	s := t.Second()

	ret[0] = byte(h / 10)
	ret[1] = byte(h % 10)
	ret[2] = clockdigits.Colon
	ret[3] = byte(m / 10)
	ret[4] = byte(m % 10)
	ret[5] = clockdigits.Colon
	ret[6] = byte(s / 10)
	ret[7] = byte(s % 10)

	return ret
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

	//figure out how many lines total (should be 5)
	for i := 0; i < len(clock[0]); i++ {
		for _, digit := range clock {
			for _, pixel := range digit[i] {
				if pixel {
					fmt.Print("█")
				} else {
					fmt.Print("░")
				}
			}
			fmt.Print("░")
		}
		fmt.Println()
	}
}

/*
	if pixel {
		fmt.Print("█")
	} else {
		fmt.Print("░")
	}

*/
