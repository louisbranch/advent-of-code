package main

import "strings"

func NoiseMax(input string) (secret string) {
	compare := func(prev, next int) bool {
		return prev < next
	}

	return noise(input, compare)
}

func NoiseMin(input string) (secret string) {
	compare := func(prev, next int) bool {
		if prev == 0 {
			return true
		}

		return next < prev
	}

	return noise(input, compare)
}

func noise(input string, compare func(prev, next int) bool) (secret string) {
	lines := strings.Split(input, "\n")

	if len(lines) == 0 {
		return
	}

	length := len(lines[0])

	for i := 0; i < length; i++ {
		m := make(map[byte]int)

		for _, line := range lines {
			c := line[i]
			m[c] += 1
		}

		var code byte
		prev := 0

		for char, next := range m {
			if compare(prev, next) {
				prev = next
				code = char
			}
		}

		secret += string(code)
	}

	return
}
