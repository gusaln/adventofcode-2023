package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

func part2(raw io.Reader) int64 {
	s := bufio.NewScanner(raw)

	// We need to scan to get the first line
	s.Scan()
	seeds := parseSeedRanges(s.Text())
	// log.Println("Seeds: ", seeds)

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
		// log.Println("--- New mapper ---")

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
			result := mapper.convertRange(seeds[i])

			seeds[i] = result[0]
			if len(result) > 1 {
				seeds = append(seeds, result[1:]...)
			}
		}
	}

	// log.Println("Locations: ", seeds)

	var min int64 = math.MaxInt64
	for _, s := range seeds {
		if s.start < min {
			min = s.start
		}
	}

	return min
}

type numberRange struct {
	start, end int64
}

func newSeedRange(start, length int64) numberRange {
	return numberRange{
		start: start,
		end:   start + length - 1,
	}
}

func parseSeedRanges(raw string) []numberRange {
	_, raw, _ = strings.Cut(raw, ": ")

	r := []numberRange{}
	var start int64 = -1
	for _, nStr := range strings.Split(raw, " ") {
		n, _ := strconv.ParseInt(nStr, 10, 64)
		if start < 0 {
			start = n
		} else {
			r = append(r, newSeedRange(start, n))
			start = -1
		}

	}

	return r
}

func (sR numberRange) String() string {
	return fmt.Sprintf("<Range [%d, %d]>", sR.start, sR.end)
}

// func (sR seedRange) length() int64 {
// 	return sR.end - sR.start + 1
// }

func (sR numberRange) shift(offset int64) numberRange {
	return numberRange{sR.start + offset, sR.end + offset}
}

func (t transformation) containsRange(seedRange numberRange) bool {
	return seedRange.start >= t.srcStart && seedRange.end <= t.srcEnd
}

func (t transformation) isContainedInRange(seedRange numberRange) bool {
	return t.srcStart >= seedRange.start && t.srcEnd <= seedRange.end
}

func (t transformation) overlapsRange(seedRange numberRange) bool {
	return seedRange.start <= t.srcEnd && seedRange.end >= t.srcStart
}

func (t transformation) convertRange(sR numberRange) []numberRange {
	if !t.overlapsRange(sR) {
		return []numberRange{sR}
	}

	if t.containsRange(sR) {
		a := sR.shift(t.delta)
		// log.Println(t, "=> from", sR, "to ->", a)
		return []numberRange{a}
	}

	if t.isContainedInRange(sR) {
		var pre, in, post numberRange

		pre = numberRange{start: sR.start, end: t.srcStart - 1}
		in = numberRange{start: t.srcStart, end: t.srcEnd}.shift(t.delta)
		post = numberRange{start: t.srcEnd + 1, end: sR.end}

		// log.Println(t, "=> from", sR, "3-split to ->", in, pre, post)

		return []numberRange{in, pre, post}
	}

	var in, out numberRange
	if sR.end <= t.srcEnd {
		out = numberRange{start: sR.start, end: t.srcStart - 1}
		in = numberRange{start: t.srcStart, end: sR.end}.shift(t.delta)
	} else {
		in = numberRange{start: sR.start, end: t.srcEnd}.shift(t.delta)
		out = numberRange{start: t.srcEnd + 1, end: sR.end}
	}

	// log.Println(t, "=> from", sR, "2-split to ->", in, out)

	return []numberRange{in, out}
}

func (m mapper) convertRange(sR numberRange) []numberRange {
	for _, r := range m.transformations {
		if r.overlapsRange(sR) {
			return r.convertRange(sR)
		}
	}

	return []numberRange{sR}
}
