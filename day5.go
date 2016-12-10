package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func DoorPassword() {
	password := ""
	i := 0

	for len(password) < 8 {
		h := md5.New()
		io.WriteString(h, "ugkcyxxp")
		io.WriteString(h, strconv.Itoa(i))
		hex := fmt.Sprintf("%0Xd", h.Sum(nil))

		if strings.Index(hex, "00000") == 0 {
			password += string(hex[5])
		}

		i++
	}

	fmt.Println(password)
}

func FancyDoorPassword() {
	var password [8]*byte
	i := 0

	for {
		h := md5.New()
		io.WriteString(h, "ugkcyxxp")
		io.WriteString(h, strconv.Itoa(i))
		i++
		hex := fmt.Sprintf("%0Xd", h.Sum(nil))

		if strings.Index(hex, "00000") != 0 {
			continue
		}

		p, err := strconv.Atoi(string(hex[5]))
		if err != nil || p < 0 || p > 7 {
			continue
		}

		v := hex[6]
		if password[p] == nil {
			password[p] = &v
		}

		done := true
		for _, b := range password {
			if b != nil {
				fmt.Printf("%s", string(*b))
			} else {
				fmt.Print("_")
				done = false
			}
		}
		fmt.Println()
		if done {
			break
		}
	}
}
