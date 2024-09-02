package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Nice strings with the original rules
func Nice(str string) bool {
	vowels := "aeiou"
	count := 0
	hasDouble := false

	for i := 0; i < len(str); i++ {

		// remove the bad pairs
		if i > 0 {
			if (i > 0 && str[i] == 'b' && str[i-1] == 'a') ||
				(str[i] == 'd' && str[i-1] == 'c') ||
				(str[i] == 'q' && str[i-1] == 'p') ||
				(str[i] == 'y' && str[i-1] == 'x') {
				return false
			}
		}
		// Counting vowels
		if isVowel(str[i], vowels) {
			count++
		}

		// Check the consecutive chars
		if i > 0 && str[i] == str[i-1] {
			hasDouble = true
		}

	}
	return count >= 3 && hasDouble
}

func isVowel(s byte, vowels string) bool {
	for i := 0; i < len(vowels); i++ {
		if s == vowels[i] {
			return true
		}
	}
	return false
}
func repeats(s string) bool {
	// i != i + 1 && i == i + 2
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			fmt.Println(string(s[i]), string(s[i+1]), string(s[i+2]))
			return true
		}
	}
	return false
}

func Nice2(s string) bool {
	if !repeats(s){
		return false
	}
	for i := 0; i < len(s)-2; i++ {
		pair := s[i: i+2]
		if strings.Contains(s[i+2:], pair){
			return true
		}
	}
	return false
}


func NiceStrings() int {
	count := 0
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error Occured %v\n", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if Nice2(line) {

			fmt.Println(line)
			count++
		}
	}

	return count
}

func main() {
	// s := "ieodomkazucvgmuy"
	// fmt.Println(Nice(s))
	// fmt.Println("Final =>", NiceStrings())


	// fmt.Println("Nice2", Nice2(s))

	fmt.Println(NiceStrings())
}
