package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	intPtr := flag.Int("part", 1, "Part")

	flag.Parse()

	inputFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer inputFile.Close()

	if *intPtr == 1 {
		fmt.Println("Part 1", part1(inputFile))
	} else {
		fmt.Println("Part 2", part2(inputFile))
	}

}

var availableCubes = map[string]byte{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func part1(input io.Reader) (sum int) {
	s := bufio.NewScanner(input)

	sum = 0
	for i := 1; s.Scan(); i++ {
		game := s.Text()
		// fmt.Println(game)

		sets := strings.Split(game, ": ")[1]

		isValid := true
		for _, set := range strings.Split(sets, "; ") {
			// fmt.Println("--> Set", set)

			for _, cubeQty := range strings.Split(set, ", ") {
				parts := strings.Split(cubeQty, " ")

				qty, _ := strconv.ParseInt(parts[0], 10, 64)
				color := parts[1]

				if availableCubes[color] < byte(qty) {
					// fmt.Println(game)
					// fmt.Printf("xxx Seen %d %s cubes, but there where %d\n", qty, color, availableCubes[color])

					isValid = false
					break
				}
			}

			if !isValid {
				break
			}
		}

		if isValid {
			sum += i
			// fmt.Printf("--> Add %d\n", i)
		}
	}

	return sum
}

func part2(input io.Reader) (power int64) {
	s := bufio.NewScanner(input)

	power = 0
	for i := 1; s.Scan(); i++ {
		game := s.Text()
		// fmt.Println(game)

		sets := strings.Split(game, ": ")[1]

		minimumCubes := map[string]int64{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, set := range strings.Split(sets, "; ") {
			// fmt.Println("--> Set", set)

			for _, cubeQty := range strings.Split(set, ", ") {
				parts := strings.Split(cubeQty, " ")

				qty, _ := strconv.ParseInt(parts[0], 10, 64)
				color := parts[1]

				if minimumCubes[color] < qty {
					minimumCubes[color] = qty
				}
			}
		}

		power = power + (minimumCubes["red"] * minimumCubes["green"] * minimumCubes["blue"])
	}

	return power
}
