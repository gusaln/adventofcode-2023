package main

import (
	"os"
	"strings"
	"testing"
)

func TestPart1_TestInput(t *testing.T) {
	var expected int

	inputFile, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected = 4361
	if r := part1(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}

func TestPart1_ParsesDirections(t *testing.T) {
	type testCase struct {
		dir   string
		input string
	}

	var cases = []testCase{
		{dir: "ne", input: "*..\n.1.\n..."},
		{dir: "n", input: ".*.\n.1.\n..."},
		{dir: "nw", input: "..*\n.1.\n..."},
		{dir: "e", input: "...\n*1.\n..."},
		{dir: "se", input: "...\n.1.\n*.."},
		{dir: "s", input: "...\n.1.\n.*."},
		{dir: "sw", input: "...\n.1.\n..*"},
	}

	for _, d := range cases {
		if r := part1(strings.NewReader(d.input)); r != 1 {
			t.Errorf("got %d for direction %s, expected 1", r, d.dir)
		}
	}
}

func TestPart1_ParsesEndCorrectly(t *testing.T) {
	var expected int

	s := "..975\n" +
		"....#"

	expected = 975
	if r := part1(strings.NewReader(s)); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}

	s = "..975\n" +
		"....."

	expected = 0
	if r := part1(strings.NewReader(s)); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}

func TestPart1_Input(t *testing.T) {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected := 520019
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

	expected = 467835
	if r := part2(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}

func TestPart2_Input(t *testing.T) {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("got error %s", err.Error())
	}
	defer inputFile.Close()

	expected := 75519888
	if r := part2(inputFile); r != expected {
		t.Errorf("got %d, expected %d", r, expected)
	}
}
