package main

import (
	"bufio"
	"io"
)

var values = map[string]rune{
	"zero":  '0',
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

var byFirst = map[rune][]string{
	'z': {"zero"},
	'o': {"one"},
	't': {"two", "three"},
	'f': {"four", "five"},
	's': {"six", "seven"},
	'e': {"eight"},
	'n': {"nine"},
}

var byLast = map[rune][]string{
	'o': {"zero", "two"},
	'e': {"one", "three", "five", "nine"},
	'r': {"four"},
	'x': {"six"},
	'n': {"seven"},
	't': {"eight"},
}

func part2(input io.Reader) (sum int32) {
	s := bufio.NewScanner(input)

	for s.Scan() {
		line := s.Text()
		chars := []rune(line)
		// fmt.Println("line:", line)

		var first, last rune = -1, -1
		l := len(chars)
		for i := 0; i < l; i++ {
			c := chars[i]
			if isDigit(c) {
				first = c - '0'
				break
			}

			if digits, isP := byFirst[c]; isP {
				// fmt.Println(">>> Maybe", string(c), digits)

				for _, digit := range digits {
					dLen := len(digit)

					maybeDigit := string(chars[i : i+dLen])
					if _, isPresent := values[maybeDigit]; isPresent {
						first = values[maybeDigit] - '0'
						// fmt.Println(">>> IT IS", string(first))
						break
					}
				}

				if first > -1 {
					break
				}
			}
		}

		for i := l - 1; i >= 0; i-- {
			c := chars[i]
			if isDigit(c) {
				last = c - '0'
				break
			}

			if digits, isP := byLast[c]; isP {
				// fmt.Println(">>> Maybe", string(c), digits)

				for _, digit := range digits {
					dLen := len(digit)
					if i < dLen-1 {
						continue
					}

					maybeDigit := string(chars[i-dLen+1 : i+1])
					// fmt.Println(">>> MAYBE", maybeDigit)

					if _, isPresent := values[maybeDigit]; isPresent {
						last = values[maybeDigit] - '0'
						// fmt.Println(">>> IT IS", string(last))
						break
					}
				}

				if last > -1 {
					break
				}
			}
		}

		// fmt.Printf("first=%v last=%v\n", first, last)
		sum += int32(first*10 + last)
	}

	return sum
}
