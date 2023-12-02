package main

import (
	"os"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	if r := part1(strings.NewReader("1abc2")); r != 12 {
		t.Errorf("got %d, expected 12", r)
	}

	inputFile, err := os.Open("test1.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()
	if r := part1(inputFile); r != 142 {
		t.Errorf("got %d, expected 142", r)
	}
}

func TestPart1Input(t *testing.T) {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()
	if r := part1(inputFile); r != 54940 {
		t.Errorf("got %d, expected 54940", r)
	}
}

func TestPart2(t *testing.T) {
	if r := part2(strings.NewReader("two1nine")); r != 29 {
		t.Errorf("got %d, expected 29", r)
	}

	inputFile, err := os.Open("test2.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()
	if r := part2(inputFile); r != 281 {
		t.Errorf("got %d, expected 281", r)
	}
}

func TestPart2Input(t *testing.T) {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()
	if r := part2(inputFile); r != 54208 {
		t.Errorf("got %d, expected 54208", r)
	}
}
