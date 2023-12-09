package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	intPtr := flag.Int("part", 1, "Part")
	testPtr := flag.Bool("test", false, "Test")

	flag.Parse()

	file := "input.txt"
	if *testPtr {
		file = "test.txt"
	}
	inputFile, err := os.Open(file)
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
