package main

import "strings"

type keypad struct {
	coordinates
}

var pad = map[coordinates]string{
	{-1, 1}:  "1",
	{0, 1}:   "2",
	{1, 1}:   "3",
	{-1, 0}:  "4",
	{0, 0}:   "5",
	{1, 0}:   "6",
	{-1, -1}: "7",
	{0, -1}:  "8",
	{1, -1}:  "9",
}

func BathroomCode(cmds string) string {
	var k keypad

	code := ""

	for _, line := range strings.Split(cmds, "\n") {
		for _, c := range line {
			k.move(c)
		}
		code += pad[k.coordinates]
	}

	return code
}

func (k *keypad) move(direction rune) {
	switch direction {
	case 'U':
		if k.y < 1 {
			k.y += 1
		}
	case 'D':
		if k.y > -1 {
			k.y -= 1
		}
	case 'R':
		if k.x < 1 {
			k.x += 1
		}
	case 'L':
		if k.x > -1 {
			k.x -= 1
		}
	}
}
