package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	message := strings.Join(os.Args[1:], " ")

	nextShouldBeCapital := isUpper(message[0])
	newMessage := strings.Builder{}
	for i := range message {
		var nextString string

		if !isWordChar(message[i]) {
			nextString = toString(message[i])
		} else {
			if nextShouldBeCapital {
				nextString = strings.ToUpper(toString(message[i]))
			} else {
				nextString = strings.ToLower(toString(message[i]))
			}

			nextShouldBeCapital = !nextShouldBeCapital
		}

		newMessage.WriteString(nextString)
	}

	fmt.Println(newMessage.String())
}

func isWordChar(ch byte) bool {
	return isLower(ch) || isUpper(ch)
}

func isLower(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}

func isUpper(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}

func toString(b byte) string {
	return string([]byte{b})
}
