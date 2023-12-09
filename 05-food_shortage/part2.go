package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func parseSeedRanges(raw string) []int64 {
	_, raw, _ = strings.Cut(raw, ": ")

	seeds := make([]int64, 0, 512)
	var start int64 = -1
	for _, nStr := range strings.Split(raw, " ") {
		n, _ := strconv.ParseInt(nStr, 10, 64)
		if start < 0 {
			start = n
		} else {
			for i := 0; i < int(n); i++ {
				seeds = append(seeds, start+int64(i))
			}
			start = -1
		}

	}

	return seeds
}

func part2(raw io.Reader) int64 {
	s := bufio.NewScanner(raw)

	// We need to scan to get the first line
	s.Scan()
	seeds := parseSeedRanges(s.Text())
	fmt.Println("Seeds: ", seeds)

	// Ignore empty line
	s.Scan()

	// For each mapping
	for s.Scan() {
		// Advance if we are at a title
		line := s.Text()
		if !strings.HasSuffix(line, ":") {
			continue
		}

		mapper := newMapper()
		// fmt.Println("--- New mapper ---")

		// Build mapper
		for s.Scan() {
			// If we find an empty line, we finished the mappings
			rawRange := s.Text()
			if rawRange == EMPTY {
				break
			}

			mapper.addMappingRange(rawRange)
		}

		for i := 0; i < len(seeds); i++ {
			before := seeds[i]
			after := mapper.convert(before)

			// fmt.Println("from", before, "to", after)

			seeds[i] = after
		}
	}

	// Doing extra work
	var min int64 = seeds[0]
	for _, l := range seeds {
		if l < min {
			min = l
		}
	}

	return min
}
