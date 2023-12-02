package main

import (
	"bufio"
	"io"
)

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func part1(input io.Reader) (sum int32) {
	s := bufio.NewScanner(input)

	for s.Scan() {
		line := []rune(s.Text())

		var first, last rune
		l := len(line)
		for i := 0; i < l; i++ {
			c := line[i]
			if isDigit(c) {
				first = c - '0'
				break
			}
		}

		for i := l - 1; i >= 0; i-- {
			c := line[i]
			if isDigit(c) {
				last = c - '0'
				break
			}
		}

		sum += int32(first*10 + last)
	}

	return sum
}
