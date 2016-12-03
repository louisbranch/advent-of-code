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
	var p position

	cmds := strings.Split(command, ", ")
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
		p.walk(t)
	}

	return p.distance(), nil
}

func (p *position) walk(times int) {
	switch p.angle {
	case 0:
		p.y += times
	case 90:
		p.x += times
	case 180:
		p.y -= times
	case 270:
		p.x -= times
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
	d := p.x + p.y
	if d < 0 {
		d *= -1
	}
	return d
}
