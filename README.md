# Advent of Code 2023 in Go

Solving Advent of Code for **fun**.

Expect messy code and ocasional refactors.

## Structure

Every day has a numbered folder followed by a hyphen and a more descriptive name.
Tests inputs and the main input are stored in text files.
There is a `main.go` file that has a CLI for computing the outputs, and each part of the problem is _usually_ in its own file inside a function just like this:

```go
func part1(input io.Reader) (output int32) {
    // output type varies
}
```

Along side any other helper functions.
This makes them easier to test with Go's builtin support for unit tests.
