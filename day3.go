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

	//"810  679   10"

	return count, nil
}

func (t triangle) Valid() bool {
	return t.a+t.b > t.c &&
		t.b+t.c > t.a &&
		t.c+t.a > t.b
}
