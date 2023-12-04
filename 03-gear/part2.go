package main

import (
	"fmt"
	"io"
	"strings"
)

func getCoord(x, y int) string {
	return fmt.Sprintf("%d.%d", x, y)
}

var present = struct{}{}

func part2(raw io.Reader) (sum int) {
	inputBytes, _ := io.ReadAll(raw)

	input := string(inputBytes)
	lines := strings.Split(input, "\n")
	linesCount := len(lines)
	lineLen := len(lines[0])

	sum = 0
	possibleGearNumbers := map[string]int32{}
	for y, line := range lines {

		// isPartNumber := false
		var number int32 = -1
		var gearCoords = map[string]struct{}{}
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

				for coord := range gearCoords {
					if other, exists := possibleGearNumbers[coord]; exists {
						sum += int(other * number)

						delete(possibleGearNumbers, coord)
						break
					}

					possibleGearNumbers[coord] = number
				}

				// Reset the values
				number = -1
				gearCoords = map[string]struct{}{}
			}

			// Check for '*'
			if number >= 0 {
				for _, d := range directions {
					yD := y + d[0]
					xD := x + d[1]

					if yD < 0 || yD >= linesCount || xD < 0 || xD >= lineLen {
						continue
					}

					if rune(lines[yD][xD]) == '*' {
						gearCoords[getCoord(xD, yD)] = present

						break
					}
				}
			}
		}

		for coord := range gearCoords {
			if other, exists := possibleGearNumbers[coord]; exists {
				sum += int(other * number)

				delete(possibleGearNumbers, coord)
				break
			}

			possibleGearNumbers[coord] = number
		}
	}

	return sum
}
