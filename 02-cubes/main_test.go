package main

import (
	"os"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	var expected int

	expected = 1
	if r := part1(strings.NewReader("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}

	expected = 0
	if r := part1(strings.NewReader("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}

	inputFile, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected = 8
	if r := part1(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}

func TestPart1Input(t *testing.T) {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected := 2810
	if r := part1(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}

func TestPart2(t *testing.T) {
	var expected int64

	expected = 48
	if r := part2(strings.NewReader("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}

	expected = 12
	if r := part2(strings.NewReader("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}

	expected = 1560
	if r := part2(strings.NewReader("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}

	expected = 630
	if r := part2(strings.NewReader("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red")); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}

	expected = 36
	if r := part2(strings.NewReader("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}

	inputFile, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected = 2286
	if r := part2(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}

func TestPart2Input(t *testing.T) {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected := int64(69110)
	if r := part2(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}
