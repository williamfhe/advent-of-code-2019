package main

import "fmt"

const (
	rangeStart = 136760
	rangeStop  = 595730
)

func validateNumber(number int) bool {
	// default value > 10 (arbitrarily chosen)
	lastDigit := 42

	hasDigitPair := false
	isFirstValidPair := false
	groupSize := 1

	for number > 0 {
		currentDigit := number % 10
		number /= 10

		if currentDigit > lastDigit {
			return false
		}

		if currentDigit == lastDigit && !hasDigitPair {
			groupSize++

			if groupSize == 2 {
				isFirstValidPair = true
			}

			if groupSize > 2 && isFirstValidPair {
				isFirstValidPair = false
			}

		} else {
			if groupSize == 2 && isFirstValidPair {
				hasDigitPair = true
			}
			groupSize = 1
		}

		lastDigit = currentDigit
	}

	return hasDigitPair || isFirstValidPair
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
