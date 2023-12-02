package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

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
		line := []rune(s.Text())
		// fmt.Println("D:", line)

		var first, last rune
		l := len(line)
		for i := 0; i < l; i++ {
			c := line[i]
			if isDigit(c) {
				first = c
				break
			}
		}

		for i := l - 1; i >= 0; i-- {
			c := line[i]
			if isDigit(c) {
				last = c
				break
			}
		}

		calibrationValue, _ := strconv.ParseInt(string([]rune{first, last}), 10, 64)

		// fmt.Println("first", string(first), "last", string(last), "calibration", calibrationValue)
		sum += int(calibrationValue)
	}

	fmt.Println(sum)
}
