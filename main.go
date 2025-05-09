package main

import (
	"bufio"
	"fmt"
	"github.com/Qwertymart/Roman_Calculator/evaluation"
	"os"
	"strings"
)

func main() {
	tests := []string{
		"IX + I",         // 9 + 1 = 10 => X
		"X - III",        // 10 - 3 = 7 => VII
		"VI * II",        // 6 * 2 = 12 => XII
		"XII / III",      // 12 / 3 = 4 => IV
		"(X + II) * II",  // (10 + 2) * 2 = 24 => XXIV
		"X + (VI / III)", // 10 + (6 / 3) = 12 => XII
		"X / (V - V)",    // ошибка
		"IV + BAD",       // недопустимый символ => ошибка
	}

	for _, expr := range tests {
		fmt.Printf("Evaluating: %s\n", expr)
		result, err := evaluation.Evaluate(expr)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Result: %s (%d)", result.Roman, result.Value)
		}
		fmt.Println("----------------------------")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter your expression (or type 'exit' to quit):")
		if !scanner.Scan() {
			break // EOF
		}
		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "exit" {
			break
		}
		result, err := evaluation.Evaluate(userInput)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Result: %s (%d)\n", result.Roman, result.Value)
		}
	}
}
