package numtowords

import (
	"strings"
)

var numLessThanTwenty = map[int64]string{
	1: "one", 2: "two", 3: "three", 4: "four",
	5: "five", 6: "six", 7: "seven", 8: "eight",
	9: "nine", 10: "ten", 11: "eleven", 12: "twelve",
	13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen",
	17: "seventeen", 18: "eighteen", 19: "nineteen", 20: "twenty",
	30: "thirty", 40: "forty", 50: "fifty", 60: "sixty",
	70: "seventy", 80: "eighty", 90: "ninety",
}

var numberGroups = map[int]string{
	1: "", 2: "thousand", 3: "million", 4: "billion",
	5: "trillion", 6: "quadrillion", 7: "quintillion",
}

func returnMatchingString(remainder int64, quotient int64, groups int, appendHundred bool) string {
	var output strings.Builder

	if quotient != 0 {
		if v, found := numLessThanTwenty[quotient]; found {
			numToWrite := v

			if appendHundred {
				numToWrite = " " + numToWrite + " hundred"
			}

			output.WriteString(numToWrite)
		}
	}

	if remainder != 0 {
		if v, found := numLessThanTwenty[remainder]; found {
			numToWrite := " " + v + " " + numberGroups[groups]
			output.WriteString(numToWrite)
		} else {
			remainder, quotient := divideByCertainNumber(remainder, 10)
			quotient = quotient * 10
			result := returnMatchingString(remainder, quotient, groups, false)
			output.WriteString(" " + result + " ")
		}
	}

	if quotient != 0 && remainder == 0 {
		numToWrite := " " + numberGroups[groups]
		output.WriteString(numToWrite)
	}

	return output.String()
}

func divideByCertainNumber(input int64, divideBy int64) (int64, int64) {
	remainder := input % divideBy
	quotient := input / divideBy

	return remainder, quotient
}

//GroupIntoThree Chunk the number into 3 digit groups
func GroupIntoThree(num int64) map[int]string {
	groups := make(map[int]string)
	initial := 1

	for num > 0 {
		splitNum := num % 1000
		remainder, quotient := divideByCertainNumber(splitNum, 100)
		inWords := returnMatchingString(remainder, quotient, initial, true)
		groups[initial] = inWords
		num /= 1000
		initial = initial + 1
	}

	return groups
}

//Transform Transform number to words
func Transform(num int64) string {
	var output strings.Builder
	numGroups := GroupIntoThree(num)

	for i := len(numGroups); i > 0; i-- {
		output.WriteString(numGroups[i])
	}

	if output.String() == "" {
		return "zero"
	}

	return strings.TrimSpace(strings.Replace(output.String(), "  ", " ", -1))
}
