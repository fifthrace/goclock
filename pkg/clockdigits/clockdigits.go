package clockdigits

//Digit type stores a 5 x 3 array of string to represent a digital number
type Digit [5][3]bool

//Colon is representing a : character in this context
const Colon byte = 10

//GetDigit given an integer between 0 and 10 return a Digit
func GetDigit(n byte) Digit {
	var dig Digit

	switch n {
	case 0:
		dig = [5][3]bool{
			[3]bool{true, true, true},
			[3]bool{true, false, true},
			[3]bool{true, false, true},
			[3]bool{true, false, true},
			[3]bool{true, true, true},
		}
	case 1:
		dig = [5][3]bool{
			[3]bool{true, true, false},
			[3]bool{false, true, false},
			[3]bool{false, true, false},
			[3]bool{false, true, false},
			[3]bool{true, true, true},
		}
	case 2:
		dig = [5][3]bool{
			[3]bool{true, true, true},
			[3]bool{false, false, true},
			[3]bool{true, true, true},
			[3]bool{true, false, false},
			[3]bool{true, true, true},
		}
	case 3:
		dig = [5][3]bool{
			[3]bool{true, true, true},
			[3]bool{false, false, true},
			[3]bool{true, true, true},
			[3]bool{false, false, true},
			[3]bool{true, true, true},
		}
	case 4:
		dig = [5][3]bool{
			[3]bool{true, false, true},
			[3]bool{true, false, true},
			[3]bool{true, true, true},
			[3]bool{false, false, true},
			[3]bool{false, false, true},
		}
	case 5:
		dig = [5][3]bool{
			[3]bool{true, true, true},
			[3]bool{true, false, false},
			[3]bool{true, true, true},
			[3]bool{false, false, true},
			[3]bool{true, true, true},
		}
	case 6:
		dig = [5][3]bool{
			[3]bool{true, true, true},
			[3]bool{true, false, false},
			[3]bool{true, true, true},
			[3]bool{true, false, true},
			[3]bool{true, true, true},
		}
	case 7:
		dig = [5][3]bool{
			[3]bool{true, true, true},
			[3]bool{false, false, true},
			[3]bool{false, false, true},
			[3]bool{false, false, true},
			[3]bool{false, false, true},
		}
	case 8:
		dig = [5][3]bool{
			[3]bool{true, true, true},
			[3]bool{true, false, true},
			[3]bool{true, true, true},
			[3]bool{true, false, true},
			[3]bool{true, true, true},
		}
	case 9:
		dig = [5][3]bool{
			[3]bool{true, true, true},
			[3]bool{true, false, true},
			[3]bool{true, true, true},
			[3]bool{false, false, true},
			[3]bool{false, false, true},
		}
	case Colon:
		dig = [5][3]bool{
			[3]bool{false, false, false},
			[3]bool{false, true, false},
			[3]bool{false, false, false},
			[3]bool{false, true, false},
			[3]bool{false, false, false},
		}
	}

	return dig
}
