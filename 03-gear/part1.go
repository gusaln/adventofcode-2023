package main

import (
	"io"
	"strings"
)

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func isValidSymbol(c rune) bool {
	return c != '.' && (c < '0' || c > '9')
}

type direction [2]int

var directions = []direction{
	{-1, -1},
	{-1, 0},
	{-1, +1},
	{0, -1},
	{0, +1},
	{+1, -1},
	{+1, 0},
	{+1, +1},
}

func part1(raw io.Reader) (sum int) {
	inputBytes, _ := io.ReadAll(raw)

	input := string(inputBytes)
	lines := strings.Split(input, "\n")
	linesCount := len(lines)
	lineLen := len(lines[0])

	sum = 0
	for y, line := range lines {

		isPartNumber := false
		var number int32 = -1
		for x := 0; x < lineLen; x++ {
			if number < 0 {
				// IF we DON'T have a number

				if isDigit(rune(line[x])) {
					number = int32(line[x] - '0')
				}

			} else if isDigit(rune(line[x])) {
				// IF we DO have a number AND current position is a digit

				number = (number * 10) + int32(line[x]-'0')

			} else {
				// IF we DO have a number AND current position is NOT a digit

				if isPartNumber {
					// fmt.Println(number, "is a part number")
					sum += int(number)
				} else {
					// fmt.Println(number, "is NOT a part number")
				}

				// Reset the values
				number = -1
				isPartNumber = false
			}

			// Check validity if necessary
			if number >= 0 && !isPartNumber {
				for _, d := range directions {
					yD := y + d[0]
					xD := x + d[1]

					if yD < 0 || yD >= linesCount || xD < 0 || xD >= lineLen {
						continue
					}

					isPartNumber = isValidSymbol(rune(lines[yD][xD]))
					// fmt.Printf("Checking %s for %d (%d,%d) = %v \n", string(lines[yD][xD]), number, x, y, isPartNumber)
					if isPartNumber {
						break
					}
				}
			}
		}

		// If the number is at the end
		if isPartNumber {
			sum += int(number)
		}
	}

	return sum
}
