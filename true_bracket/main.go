package main

import "fmt"

func isBracket(str string) bool {
	var stack []string
	strMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	for _, v := range str {
		s := string(v)
		_, ok := strMap[str]
		if ok {
			if len(stack) == 0 {
				return false
			}
			if strMap[str] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s)
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	a := "hello world"
	fmt.Printf("%T %s %d", a, a, a)
}
