package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  = "0123456789"
	specialChars = "!@#$%^&*()-_=+,.?/:;{}[]~"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Random Password Generator v1 ===")

	length := askForLength()
	useUpper := askYesNo("Include capital letters? (y/n) ")
	useNumbers := askYesNo("Include numbers? (y/n) ")
	useSpecial := askYesNo("Include special characters? (y/n) ")

	password := generatePassword(length, useUpper, useNumbers, useSpecial)

	fmt.Printf("\nYour generated password: \n%s\n", password)
}

func askForLength() int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Password length (8-64): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		length, err := strconv.Atoi(input)
		if err != nil || length < 8 || length > 64 {
			fmt.Println("Please enter a number between 8 and 64")
			continue
		}

		return length
	}
}

func askYesNo(question string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(question)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "y" || input == "yes" {
			return true
		} else if input == "n" || input == "no" {
			return false
		}

		fmt.Println("Please answer with 'y' (yes) or 'n' (no)")
	}
}

func generatePassword(length int, useUpper, useNumbers, useSpecial bool) string {
	var charPool string

	charPool = lowerChars

	if useUpper {
		charPool += upperChars
	}
	if useNumbers {
		charPool += numberChars
	}
	if useSpecial {
		charPool += specialChars
	}

	if len(charPool) == 0 {
		charPool = lowerChars
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = charPool[rand.Intn(len(charPool))]
	}

	return string(password)
}
