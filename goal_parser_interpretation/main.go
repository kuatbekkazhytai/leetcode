package main

import "fmt"

func main() {
	fmt.Println(interpret("(al)G(al)()()G"))
}

func interpret(command string) string {
	str := ""
	i := 0
	for i < len(command) {
		if command[i] == 'G' {
			str += "G"
			i++
		} else if command[i] == '(' {
			if command[i+1] == ')' {
				str += "o"
				i += 2
			} else {
				str += "al"
				i += 4
			}
		}
	}

	return str
}
