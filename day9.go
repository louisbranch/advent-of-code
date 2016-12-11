package main

import (
	"log"
	"strconv"
	"strings"
)

func BulkDecompress(input string) int {
	size := 0
	for _, line := range strings.Split(input, "\n") {
		dec := Decompress(line)
		size += len(dec)
	}
	return size
}

func Decompress(input string) string {
	dec := ""
	buf := ""
	i := 0

	capture := false
	for i < len(input) {
		char := input[i]
		switch char {
		case '(':
			capture = true
			i++
		case ')':
			nums := strings.SplitN(buf, "x", 2)

			offset, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatal(err)
			}

			times, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatal(err)
			}

			tmp := string(input[i+1 : i+1+offset])

			dec += strings.Repeat(tmp, times)

			buf = ""
			capture = false
			i += offset + 1
		default:
			if capture {
				buf += string(char)
			} else {
				dec += string(char)
			}
			i++
		}
	}

	return dec
}
