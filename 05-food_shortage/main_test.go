package main

import (
	"os"
	"testing"
)

func TestPart1_TestInput(t *testing.T) {
	var expected int64

	inputFile, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected = 35
	if r := part1(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}

func TestPart1_Input(t *testing.T) {
	var expected int64

	inputFile, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected = 1181555926
	if r := part1(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}

func TestPart2_TestInput(t *testing.T) {
	var expected int64

	inputFile, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected = 46
	if r := part2(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}

func TestPart2_Input(t *testing.T) {
	var expected int64

	inputFile, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected = 37806486
	if r := part2(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}
