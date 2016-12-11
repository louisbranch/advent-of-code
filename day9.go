package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func BulkDecompress(input string) (v1, v2 int) {
	for _, line := range strings.Split(input, "\n") {
		v1 += len(Decompress(line))
		v2 += RecursiveDecompress(line)
	}
	return
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

func RecursiveDecompress(input string) (length int) {
	re := regexp.MustCompile(`^\((\d+)x(\d+)\)`)

	for len(input) > 0 {
		if matches := re.FindStringSubmatch(input); len(matches) == 3 {
			size := len(matches[0])

			offset, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatal(err)
			}

			times, err := strconv.Atoi(matches[2])
			if err != nil {
				log.Fatal(err)
			}
			tmp := string(input[size : size+offset])

			input = strings.Repeat(tmp, times) + string(input[size+offset:])
		} else {
			length++
			input = input[1:]
		}
	}

	return length
}
