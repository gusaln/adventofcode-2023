package main

import (
	"os"
	"strings"
	"testing"
)

func TestPart1_Line(t *testing.T) {
	var expected int

	expected = 8
	if r := part1(strings.NewReader("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}

func TestPart1_TestInput(t *testing.T) {
	var expected int

	inputFile, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected = 13
	if r := part1(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}

func TestPart1_Input(t *testing.T) {
	var expected int

	inputFile, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected = 27845
	if r := part1(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}

func TestPart2_TestInput(t *testing.T) {
	var expected int

	inputFile, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected = 30
	if r := part2(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}

func TestPart2_Input(t *testing.T) {
	var expected int

	inputFile, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected = 9496801
	if r := part2(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}
