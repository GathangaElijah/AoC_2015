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
func Basement(str string) int{
	floor := 0
	var position int
	for i := 0; i < len(str); i++{
		bracket := string(str[i])
		if bracket == "("{
			floor ++
			if floor == -1{
				position = i +1
				break
			}
		}
		if bracket == ")"{
			floor --
			if floor == -1{
				position = i +1
				break
			}
		}
		
	}
	return position
}

func main() {
	str, err := os.ReadFile("input.txt")
	if err == nil {
		floor := FloorCalc(string(str))
		fmt.Println(floor)
	}else {
		fmt.Println(err)
	}
	position := Basement(string(str))
	fmt.Println(position)
}
