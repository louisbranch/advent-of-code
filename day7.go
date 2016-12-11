package main

import "strings"

func IpAbba(input string) int {
	var c []rune
	count := 0

	for _, line := range strings.Split(input, "\n") {
		inside := false
		include := false

		for _, char := range line {
			switch char {
			case '[':
				c = c[:0]
				inside = true
				continue
			case ']':
				c = c[:0]
				inside = false
				continue
			}

			c = append(c, char)

			if len(c) < 4 {
				continue
			}

			if len(c) > 4 {
				c = c[1:]
			}

			if c[0] != c[1] && c[2] != c[3] && c[0] == c[3] && c[1] == c[2] {
				if inside {
					include = false
					break
				} else {
					include = true
				}
			}
		}

		if include {
			count++
		}

		c = c[:0]
	}

	return count
}
