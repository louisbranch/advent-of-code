package main

import "strings"

func Noise(input string) (secret string) {
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

		var most byte
		n := 0

		for k, v := range m {
			if v > n {
				n = v
				most = k
			}
		}

		secret += string(most)
	}

	return
}
