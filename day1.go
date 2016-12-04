package main

import (
	"fmt"
	"strconv"
	"strings"
)

type coordinates struct {
	x int
	y int
}

type position struct {
	coordinates
	angle int
}

func Move(command string) (int, error) {
	visited := make(map[coordinates]bool)
	p := &position{}

	cmds := strings.Split(command, ", ")
LOOP:
	for _, cmd := range cmds {
		if len(cmd) < 2 {
			return 0, fmt.Errorf("wrong command %s", command)
		}
		d := cmd[0]

		t, err := strconv.Atoi(cmd[1:])
		if err != nil {
			return 0, fmt.Errorf("wrong command %s - %s", command, err)
		}

		p.rotate(d)

		for i := 0; i < t; i++ {
			p.walk()
			if visited[p.coordinates] {
				break LOOP
			}
			visited[p.coordinates] = true
		}
	}

	return p.distance(), nil
}

func (p *position) walk() {
	switch p.angle {
	case 0:
		p.y += 1
	case 90:
		p.x += 1
	case 180:
		p.y -= 1
	case 270:
		p.x -= 1
	}
}

func (p *position) rotate(direction byte) {
	switch direction {
	case 'L':
		if p.angle == 0 {
			p.angle = 270
		} else {
			p.angle -= 90
		}
	case 'R':
		if p.angle == 270 {
			p.angle = 0
		} else {
			p.angle += 90
		}
	}
}

func (p *position) distance() int {
	x := p.x
	y := p.y

	if x < 0 {
		x *= -1
	}

	if y < 0 {
		y *= -1
	}

	return x + y
}
