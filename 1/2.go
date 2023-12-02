package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var VALUES = map[string]rune{
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

var STARTS = map[rune][]string{
	'z': []string{"zero"},
	'o': []string{"one"},
	't': []string{"two", "three"},
	'f': []string{"four", "five"},
	's': []string{"six", "seven"},
	'e': []string{"eight"},
	'n': []string{"nine"},
}

var ENDS = map[rune][]string{
	'o': []string{"zero", "two"},
	'e': []string{"one", "three", "five", "nine"},
	'r': []string{"four"},
	'x': []string{"six"},
	'n': []string{"seven"},
	't': []string{"eight"},
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		os.Exit(2)
	}
	inputFile, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer inputFile.Close()

	s := bufio.NewScanner(inputFile)

	sum := 0
	for s.Scan() {
		line := s.Text()
		chars := []rune(line)
		// fmt.Println("line:", line)

		var first, last rune
		l := len(chars)
		for i := 0; i < l; i++ {
			c := chars[i]
			if isDigit(c) {
				first = c
				break
			}

			if digitStrs, isP := STARTS[c]; isP {
				// fmt.Println(">>> Maybe", string(c), digitStrs)

				done := false
				for _, digitStr := range digitStrs {
					dLen := len(digitStr)

					maybeDigit := string(chars[i : i+dLen])
					if _, isPresent := VALUES[maybeDigit]; isPresent {
						first = VALUES[maybeDigit]
						done = true
						// fmt.Println(">>> IT IS", string(first))
						break
					}
				}

				if done {
					break
				}
			}
		}

		for i := l - 1; i >= 0; i-- {
			c := chars[i]
			if isDigit(c) {
				last = c
				break
			}

			if digitStrs, isP := ENDS[c]; isP {
				// fmt.Println(">>> Maybe", string(c), digitStrs)

				done := false
				for _, digitStr := range digitStrs {
					dLen := len(digitStr)
					if i < dLen-1 {
						continue
					}

					maybeDigit := string(chars[i-dLen+1 : i+1])
					// fmt.Println(">>> MAYBE", maybeDigit)

					if _, isPresent := VALUES[maybeDigit]; isPresent {
						last = VALUES[maybeDigit]
						// fmt.Println(">>> IT IS", string(last))
						done = true
						break
					}
				}

				if done {
					break
				}
			}
		}

		calibrationValue, _ := strconv.ParseInt(string([]rune{first, last}), 10, 64)

		// fmt.Println("first", string(first), "last", string(last), "calibration", calibrationValue)
		sum += int(calibrationValue)
	}

	fmt.Println(sum)
}
