package main

import "fmt"

const (
	rangeStart = 136760
	rangeStop  = 595730
)

func validateNumber(number int) bool {
	var hasDigitPair bool

	// default value > 10 (arbitrarily chosen)
	lastDigit := 42

	for number > 0 {
		currentDigit := number % 10
		number /= 10

		if currentDigit > lastDigit {
			return false
		}

		if currentDigit == lastDigit {
			hasDigitPair = true
		}

		lastDigit = currentDigit
	}

	return hasDigitPair
}

func main() {

	numbersMatchingCriteria := 0

	for i := rangeStart; i <= rangeStop; i++ {
		if validateNumber(i) {
			numbersMatchingCriteria++
		}
	}

	fmt.Println(numbersMatchingCriteria)
}
