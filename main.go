package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	names := []string{
		"😀",
		"$",
		" ",
		"การ",
		"ABCD",
		"Raดิส",
		"กา",
		"hello worldsdkjfksj",
		"          ssdj   ",
		"😀",
		"i4235r29348@#$@#$",
		"",
		"askd12  3123123การ",
		"askd12 😀 312312",
		"การddfd",
		"nดส่าบ",
		"\u0089",
		"AbaAZ12    ส่า่่ไไไหฟห",
		"ab          cd",
		"ab    \u01C3     cd",
		"  \u01C3 ",
		" \u01c1 ",
	}
	for _, name := range names {
		name = strings.TrimSpace(name)
		l := utf8.RuneCountInString(name)
		if l < 1 || l > 20 {
			continue
		}
		if !regexExclName.MatchString(name) {
			continue
		}
		fmt.Printf(">> %d %q\n", l, name)
	}
}

var regexExclName = regexp.MustCompile(`^[\pL\pM\pN\pZ]+$`)
