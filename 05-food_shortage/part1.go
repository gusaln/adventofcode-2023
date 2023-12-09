package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const EMPTY = ""

func part1(raw io.Reader) int64 {
	s := bufio.NewScanner(raw)

	// We need to scan to get the first line
	s.Scan()
	seeds := parseSeeds(s.Text())
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

			mapper.addTransformation(rawRange)
		}

		for i := 0; i < len(seeds); i++ {
			before := seeds[i]
			after := mapper.convert(before)

			// fmt.Println("from", before, "to", after)

			seeds[i] = after
		}
	}

	fmt.Println("Locations: ", seeds)

	var min int64 = seeds[0]
	for _, l := range seeds[0:] {
		if l < min {
			min = l
		}
	}

	return min
}

func parseSeeds(raw string) []int64 {
	_, raw, _ = strings.Cut(raw, ": ")

	a := strings.Split(raw, " ")
	seeds := make([]int64, len(a))
	for i, s := range a {
		n, _ := strconv.ParseInt(s, 10, 64)

		seeds[i] = n
	}

	return seeds
}

type transformation struct {
	dstStart int64
	dstEnd   int64
	srcStart int64
	srcEnd   int64
	delta    int64
}

func parseTransformation(raw string) transformation {
	parts := strings.Split(raw, " ")

	dst, _ := strconv.ParseInt(parts[0], 10, 64)
	src, _ := strconv.ParseInt(parts[1], 10, 64)
	length, _ := strconv.ParseInt(parts[2], 10, 64)

	return transformation{
		srcStart: src,
		dstStart: dst,
		srcEnd:   src + length - 1,
		dstEnd:   dst + length - 1,
		delta:    dst - src,
	}
}

func (r transformation) String() string {
	// return fmt.Sprintf("<Mapping [%d, %d] -> [%d, %d]>", r.srcStart, r.srcEnd, r.dstStart, r.dstEnd)
	return fmt.Sprintf("<Mapping [%d, %d] ; %+d>", r.srcStart, r.srcEnd, r.delta)
}

func (r transformation) contains(n int64) bool {
	return n >= r.srcStart && n <= r.srcEnd
}

func (r transformation) convert(n int64) int64 {
	if r.contains(n) {
		// fmt.Println("transf.", n, " , using:", r)

		return n + r.delta
	}

	return n
}

type mapper struct {
	transformations []transformation
}

func newMapper() mapper {
	return mapper{
		transformations: []transformation{},
	}
}

func (m *mapper) addTransformation(raw string) {
	m.transformations = append(m.transformations, parseTransformation(raw))
}

func (m mapper) findRangeIndex(n int64) int {
	for i, r := range m.transformations {
		if r.contains(n) {
			return i
		}
	}

	return -1
}

func (m mapper) convert(n int64) int64 {
	if i := m.findRangeIndex(n); i >= 0 {
		return m.transformations[i].convert(n)
	}

	return n
}
