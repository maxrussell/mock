package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	message := strings.Join(os.Args[1:], " ")
	newMessage := strings.Builder{}

	firstRune, _ := utf8.DecodeRune([]byte(message))
	nextShouldBeCapital := unicode.IsUpper(firstRune)
	for _, ch := range message {
		if !unicode.IsLetter(ch) {
			newMessage.WriteRune(ch)
		} else {
			if nextShouldBeCapital {
				newMessage.WriteRune(unicode.ToUpper(ch))
			} else {
				newMessage.WriteRune(unicode.ToLower(ch))
			}

			nextShouldBeCapital = !nextShouldBeCapital
		}
	}

	fmt.Println(newMessage.String())
}
