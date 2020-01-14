package clockdigits

//Digit type stores a 5 x 3 array of string to represent a digital number
type Digit [5][3]string

//GetDigit given an integer between 0 and 10 return a Digit
func GetDigit(n byte) Digit {
	var dig Digit

	switch n {
	case 0:
		dig = [5][3]string{
			[3]string{"█", "█", "█"},
			[3]string{"█", "░", "█"},
			[3]string{"█", "░", "█"},
			[3]string{"█", "░", "█"},
			[3]string{"█", "█", "█"},
		}
	case 1:
		dig = [5][3]string{
			[3]string{"█", "█", "░"},
			[3]string{"░", "█", "░"},
			[3]string{"░", "█", "░"},
			[3]string{"░", "█", "░"},
			[3]string{"█", "█", "█"},
		}
	case 2:
		dig = [5][3]string{
			[3]string{"█", "█", "█"},
			[3]string{"░", "░", "█"},
			[3]string{"█", "█", "█"},
			[3]string{"█", "░", "░"},
			[3]string{"█", "█", "█"},
		}
	case 3:
		dig = [5][3]string{
			[3]string{"█", "█", "█"},
			[3]string{"░", "░", "█"},
			[3]string{"█", "█", "█"},
			[3]string{"░", "░", "█"},
			[3]string{"█", "█", "█"},
		}
	case 4:
		dig = [5][3]string{
			[3]string{"█", "░", "█"},
			[3]string{"█", "░", "█"},
			[3]string{"█", "█", "█"},
			[3]string{"░", "░", "█"},
			[3]string{"░", "░", "█"},
		}
	case 5:
		dig = [5][3]string{
			[3]string{"█", "█", "█"},
			[3]string{"█", "░", "░"},
			[3]string{"█", "█", "█"},
			[3]string{"░", "░", "█"},
			[3]string{"█", "█", "█"},
		}
	case 6:
		dig = [5][3]string{
			[3]string{"█", "█", "█"},
			[3]string{"█", "░", "░"},
			[3]string{"█", "█", "█"},
			[3]string{"█", "░", "█"},
			[3]string{"█", "█", "█"},
		}
	case 7:
		dig = [5][3]string{
			[3]string{"█", "█", "█"},
			[3]string{"░", "░", "█"},
			[3]string{"░", "░", "█"},
			[3]string{"░", "░", "█"},
			[3]string{"░", "░", "█"},
		}
	case 8:
		dig = [5][3]string{
			[3]string{"█", "█", "█"},
			[3]string{"█", "░", "█"},
			[3]string{"█", "█", "█"},
			[3]string{"█", "░", "█"},
			[3]string{"█", "█", "█"},
		}
	case 9:
		dig = [5][3]string{
			[3]string{"█", "█", "█"},
			[3]string{"█", "░", "█"},
			[3]string{"█", "█", "█"},
			[3]string{"░", "░", "█"},
			[3]string{"░", "░", "█"},
		}
	default:
		dig = [5][3]string{
			[3]string{"░", "░", "░"},
			[3]string{"░", "░", "░"},
			[3]string{"░", "░", "░"},
			[3]string{"░", "░", "░"},
			[3]string{"░", "░", "░"},
		}
	}

	return dig
}
