package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type screen [][]bool

func TwoFactor(rows, cols int, input string) int {
	rect := regexp.MustCompile(`rect (\d+)x(\d)+`)
	rrow := regexp.MustCompile(`rotate row y=(\d+) by (\d+)`)
	rcol := regexp.MustCompile(`rotate column x=(\d+) by (\d+)`)

	s := screen{}

	for i := 0; i < rows; i++ {
		s = append(s, make([]bool, cols))
	}

	for _, line := range strings.Split(input, "\n") {
		fmt.Print(s)
		fmt.Print("----\n")
		fmt.Println(line)
		if matches := rect.FindStringSubmatch(line); len(matches) == 3 {
			x, y := coord(matches)
			rectangle(x, y, s)
			continue
		}

		if matches := rrow.FindStringSubmatch(line); len(matches) == 3 {
			x, y := coord(matches)
			rotateRow(x, y, s)
			continue
		}

		if matches := rcol.FindStringSubmatch(line); len(matches) == 3 {
			x, y := coord(matches)
			rotateCol(x, y, s)
			continue
		}

		log.Fatalf("unknown instruction %s", line)
	}

	count := 0
	for _, row := range s {
		for _, col := range row {
			if col {
				count++
			}
		}
	}
	fmt.Print(s)

	return count
}

func coord(s []string) (x, y int) {
	x, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}

	y, err = strconv.Atoi(s[2])
	if err != nil {
		log.Fatal(err)
	}

	return x, y
}

func rectangle(x, y int, s screen) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			s[j][i] = true
		}
	}
}

func rotateRow(y, n int, s screen) {
	s[y] = rotate(s[y], n)
}

func rotateCol(x, n int, s screen) {
	var tmp []bool

	for _, row := range s {
		tmp = append(tmp, row[x])
	}

	tmp = rotate(tmp, n)

	for i, row := range s {
		row[x] = tmp[i]
	}
}

func rotate(row []bool, n int) []bool {
	l := len(row)
	n = l - (n % l)
	row = append(row[n:], row[:n]...)
	return row
}

func (s screen) String() string {
	result := ""

	for _, rows := range s {
		for _, col := range rows {
			if col {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}

	return result
}
