package main

import (
	"fmt"
	"os"
)

func FloorCalc(str string) int {
	floor := 0
	for i := 0; i < len(str); i++ {
		bracket := string(str[i])
		if bracket == "(" {
			floor++
		}
		if bracket == ")" {
			floor--
		}
	}

	return floor
}

func main() {
	str, err := os.ReadFile("input.txt")
	if err == nil {
		floor := FloorCalc(string(str))
		fmt.Println(floor)
	}else {
		fmt.Println(err)
	}
}
