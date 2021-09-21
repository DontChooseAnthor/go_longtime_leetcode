package main

import (
	"fmt"
	"strings"
)

func ScanLine() string {
	var (
		c   byte
		err error
		b   []byte
	)
	for err == nil {
		_, err = fmt.Scanf("%c", &c)
		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}
	return string(b)
}

func main() {
	// 1. show word in cli, guide input
	fmt.Println("please input your words, enter is confirm")
	line := ScanLine()
	wordArr := strings.Fields(line)
	lastWord := wordArr[len(wordArr)-1]
	fmt.Println("last word length is: ", len(lastWord))
}
