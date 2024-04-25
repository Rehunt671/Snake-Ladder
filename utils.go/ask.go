package utils

import (
	"bufio"
	"fmt"
	"os"
)

func AskNumber(question string, max int) int {
	var num int
	fmt.Println(question)
	fmt.Scan(&num)
	for num < 0 || num > max {
		fmt.Println(question)
		num = AskNumber(question, max)
	}
	bufio.NewReader(os.Stdin).ReadString('\n')
	return num
}

func AskString(message string) string {
	var text string
	fmt.Println(message)
	fmt.Scan(&text)
	bufio.NewReader(os.Stdin).ReadString('\n')
	return text
}
