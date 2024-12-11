package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var input string
	fmt.Println("Enter an expression: ")
	for scanner.Scan() {
		input = scanner.Text()
		if input == "exit" {
			break
		}
		arr := strings.Split(input, " ")
		num1, _ := strconv.ParseFloat(arr[0], 64)
		num2, _ := strconv.ParseFloat(arr[2], 64)
		switch arr[1] {
			case "+":
				fmt.Println(num1 + num2)
			case "-":
				fmt.Println(num1 - num2)
			case "*":
				fmt.Println(num1 * num2)
			case "/":
				fmt.Println(num1 / num2)
		}
		fmt.Println("Enter an expression: ")
	}
}