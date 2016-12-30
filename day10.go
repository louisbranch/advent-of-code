package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type bot struct {
	low  int
	high int
}

type output []int

type receiver interface {
	get(int)
}

type step struct {
	label     string
	bot       int
	chips     [2]int
	receivers [2]receiver
}

var input = `value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2`

func main() {
	gives := regexp.MustCompile(`bot (\d+) gives low to (\w+) (\d+) and high to (\w+) (\d+)`)
	gets := regexp.MustCompile(`value (\d+) goes to bot (\d+)`)

	var steps []*step

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		inst := &step{}

		if matches := gives.FindStringSubmatch(line); len(matches) == 6 {
			inst.label = "give"

			n, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatal(err)
			}

			inst.bot = n

			n, err = strconv.Atoi(matches[2])
			if err != nil {
				log.Fatal(err)
			}

			inst.chips[0] = n

			n, err = strconv.Atoi(matches[3])
			if err != nil {
				log.Fatal(err)
			}

			inst.chips[1] = n

			steps = append(steps, inst)
			continue
		}

		if matches := gets.FindStringSubmatch(line); len(matches) == 3 {
			inst.label = "get"

			n, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatal(err)
			}

			inst.chips[0] = n

			n, err = strconv.Atoi(matches[2])
			if err != nil {
				log.Fatal(err)
			}

			inst.bot = n

			steps = append(steps, inst)
			continue
		}

		log.Fatalf("unknown instruction %s", line)
	}

	processInstructions(steps)
}

func processInstructions(steps []*step) {
	var bots []*bot
	var b *bot
	for i := 0; i < len(steps); i++ {
		step := steps[i]
		if step == nil {
			continue
		}
		switch step.label {
		case "get":
			n := step.chips[0]
			b, bots = findBot(step.bot, bots)
			b.get(n)
			steps[i] = nil
			i = 0
		case "give":
			b, bots = findBot(step.bot, bots)
			lower, bots := findBot(step.chips[0], bots)
			lower.get(b.low)
			higher, bots := findBot(step.chips[1], bots)
			higher.get(b.high)
			b.low = 0
			b.high = 0
		}
	}
}

func (b *bot) get(n int) {
	if b.low == 0 {
		b.low = n
	} else {
		if n < b.low {
			b.high = b.low
			b.low = n
		}
	}
}

func (o output) get(n int) {
	o = append(o, n)
}

func findBot(n int, bots []*bot) (*bot, []*bot) {
	var b *bot

	if len(bots) > n {
		return bots[n], bots
	}

	for i := len(bots); i <= n; i++ {
		b := &bot{}
		bots = append(bots, b)
	}

	return b, bots
}
