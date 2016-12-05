package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/bradfitz/slice"
)

type room struct {
	id       int
	code     []string
	checksum string
}

func ValidRooms(input string) (int, error) {
	count := 0

	for _, line := range strings.Split(input, "\n") {
		r, err := NewRoom(line)

		if err != nil {
			return 0, err
		}

		if r.Valid() {
			count += r.id
		}

	}

	return count, nil
}

func NewRoom(code string) (*room, error) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	tokens := strings.FieldsFunc(code, f)

	checksum, tokens := last(tokens)

	n, tokens := last(tokens)

	id, err := strconv.Atoi(n)

	if err != nil {
		return nil, fmt.Errorf("invalid id %s", n)
	}

	r := &room{
		id:       id,
		code:     tokens,
		checksum: checksum,
	}

	return r, nil
}

func last(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[len(s)-1], s[:len(s)-1]
}

func (r *room) Valid() bool {
	type letter struct {
		value string
		count int
	}

	mapping := make(map[string]int)

	for _, c := range r.code {
		for _, l := range c {
			mapping[string(l)] += 1
		}
	}

	var letters []letter

	for k, v := range mapping {
		letters = append(letters, letter{value: k, count: v})
	}

	slice.Sort(letters, func(i, j int) bool {
		if letters[i].count == letters[j].count {
			return strings.Compare(letters[i].value, letters[j].value) == -1
		} else {
			return letters[i].count > letters[j].count
		}
	})

	if len(letters) < 5 {
		return false
	}

	checksum := ""

	for i := 0; i < 5; i++ {
		checksum += letters[i].value
	}

	return r.checksum == checksum
}

func (r *room) EncryptedName() string {

	return ""
}
