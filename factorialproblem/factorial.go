package factorialproblem

func factorial(input int) int {
	if input < 0 {
		return 0
	}
	if input == 0 || input == 1 {
		return 1
	}
	return input * factorial(input-1)
}
