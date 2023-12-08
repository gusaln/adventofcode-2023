package main

import (
	"bufio"
	"io"
	"strings"
)

var present = struct{}{}

const EMPTY_NUMBER = ""

func part1(raw io.Reader) (points int) {
	s := bufio.NewScanner(raw)

	for s.Scan() {
		cardPoints := 0

		line := s.Text()
		line = line[strings.Index(line, ": ")+2:]

		parts := strings.Split(line, " | ")
		winning := map[string]struct{}{}
		for _, n := range strings.Split(parts[0], " ") {
			if n != EMPTY_NUMBER {
				winning[n] = present
			}
		}

		own := strings.Split(parts[1], " ")
		for _, n := range own {
			if _, isW := winning[n]; !isW {
				continue
			}

			if cardPoints == 0 {
				cardPoints = 1
			} else {
				cardPoints *= 2
			}
		}

		points += cardPoints
	}

	return points
}
