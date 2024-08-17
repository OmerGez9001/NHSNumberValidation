package main

import (
	"fmt"
)

func CheckValidNHSNumber(NHSNumber int) bool {
	sum := 0                     // total sum of multiplications
	checkDigit := NHSNumber % 10 // check digit is the last digit, save it into a variable.
	NHSNumber /= 10              // dividing by 10 each time we move to a new digit

	// start loop with the rest of the 9 digits
	for i := 0; i < 9; i++ {
		digit := NHSNumber % 10 // get digit (from the right)
		sum += digit * (i + 2)  // multiply digit according to the ruling of the given table
		NHSNumber /= 10
	}
	sum += checkDigit // in accordance to the algorithm, add last digit as it is to the total
	remainder := sum % 11

	// it's dividable by 11 = check digit of 0 is used: NHS number is valid
	if remainder == 0 {
		return true
	}

	// (specific case) remainder = 1: invalid and unused
	if 11-remainder == 10 {
		return false
	}

	// general case: Check if the remainder equals the check digit
	return remainder == checkDigit
}

func main() {
	NHSNumbers := []int{5990128088, 1275988113, 4536026665, 5990128087, 4536016660}
	for _, NHSNumber := range NHSNumbers {
		if CheckValidNHSNumber(NHSNumber) {
			fmt.Println(NHSNumber, ": Valid NHS Number")
		} else {
			fmt.Println(NHSNumber, ": Invalid NHS Number")
		}
	}
}
