package main

import (
	"fmt"
	"strconv"
	"strings"
)

type triangle struct {
	a, b, c int
}

func ValidTriangles(inputs string) (int, error) {
	count := 0

	for _, line := range strings.Split(inputs, "\n") {
		var t triangle
		for i, val := range strings.Fields(line) {
			n, err := strconv.Atoi(val)
			if err != nil {
				return 0, fmt.Errorf("invalid input %s", line)
			}
			switch i {
			case 0:
				t.a = n
			case 1:
				t.b = n
			case 2:
				t.c = n
			default:
				return 0, fmt.Errorf("invalid input %s", line)
			}
		}

		if t.Valid() {
			count++
		}
	}

	return count, nil
}

func ValidTrianglesColumns(inputs string) (int, error) {
	count := 0

	cols := [3][]int{}

	for _, line := range strings.Split(inputs, "\n") {

		for i, val := range strings.Fields(line) {
			n, err := strconv.Atoi(val)
			if err != nil {
				return 0, fmt.Errorf("invalid input %s", line)
			}
			cols[i%3] = append(cols[i%3], n)
		}

	}

	list := append(append(cols[0], cols[1]...), cols[2]...)

	for i := 0; i < len(list); i += 3 {
		t := triangle{a: list[i], b: list[i+1], c: list[i+2]}
		if t.Valid() {
			count++
		}
	}

	return count, nil
}
func (t triangle) Valid() bool {
	return t.a+t.b > t.c &&
		t.b+t.c > t.a &&
		t.c+t.a > t.b
}
