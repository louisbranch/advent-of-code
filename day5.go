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
