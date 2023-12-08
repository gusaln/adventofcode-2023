package main

import (
	"bufio"
	"io"
	"strings"
)

func part2(raw io.Reader) (points int) {
	s := bufio.NewScanner(raw)

	points = 0
	extras := map[int]int{}
	for card := 1; s.Scan(); card++ {
		line := s.Text()
		line = line[strings.Index(line, ": ")+2:]

		parts := strings.Split(line, " | ")
		winning := map[string]struct{}{}
		for _, n := range strings.Split(parts[0], " ") {
			if n != EMPTY_NUMBER {
				winning[n] = present
			}
		}

		copiesOfThisCard := 1 + extras[card]
		matchingNumbers := 0
		own := strings.Split(parts[1], " ")
		for _, n := range own {
			if _, isW := winning[n]; !isW {
				continue
			}

			matchingNumbers++
			n, hadCopies := extras[card+matchingNumbers]
			if hadCopies {
				extras[card+matchingNumbers] = n + copiesOfThisCard
			} else {
				extras[card+matchingNumbers] = copiesOfThisCard
			}
		}

		points++
	}

	// fmt.Println(points)

	for _, nCopies := range extras {
		points += nCopies
	}

	// fmt.Println(extras)

	return points
}
