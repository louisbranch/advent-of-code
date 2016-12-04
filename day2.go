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

var fancy = map[coordinates]string{
	{0, 2}:   "1",
	{-1, 1}:  "2",
	{0, 1}:   "3",
	{1, 1}:   "4",
	{-2, 0}:  "5",
	{-1, 0}:  "6",
	{0, 0}:   "7",
	{1, 0}:   "8",
	{2, 0}:   "9",
	{-1, -1}: "A",
	{0, -1}:  "B",
	{1, -1}:  "C",
	{0, -2}:  "D",
}

func BathroomCode(cmds string, complex bool) string {
	var k keypad
	keys := pad

	if complex {
		keys = fancy
		k.coordinates = coordinates{-2, 0}
	}

	code := ""

	for _, line := range strings.Split(cmds, "\n") {
		for _, c := range line {
			k.move(c, keys)
		}
		code += keys[k.coordinates]
	}

	return code
}

func (k *keypad) move(direction rune, keys map[coordinates]string) string {
	c := k.coordinates

	switch direction {
	case 'U':
		c.y += 1
	case 'D':
		c.y -= 1
	case 'R':
		c.x += 1
	case 'L':
		c.x -= 1
	}

	btn, ok := keys[c]

	if ok {
		k.coordinates = c
	}

	return btn
}
