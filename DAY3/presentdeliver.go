package main

import (
	"fmt"
	"os"
)

func PresentDeliver(str string) int {
	var noOfHouses int
	x, y := 0, 0
	housesVisited := make(map[string]bool)
	housesVisited[fmt.Sprintf("%d,%d", x, y)] = true

	for _, char := range str {
		switch string(char) {
		case ">":
			x++
		case "<":
			x--
		case "^":
			y++
		case "v":
			y--
		}
		housesVisited[fmt.Sprintf("%d,%d", x, y)] = true

	}

	noOfHouses = len(housesVisited)
	return noOfHouses
}

func SantaAndRobo(str string)int{
	var noOfHouses int
	x, y := 0, 0
	x2, y2 := 0, 0
	Visits := make(map[string]bool)
	//initialize  to visit the first home
	Visits[fmt.Sprintf("%d,%d", x,y)] = true
	for i, char := range str {
		if i %2 == 0{
			switch string(char) {
			case ">":
				x++
			case "<":
				x--
			case "^":
				y++
			case "v":
				y--
			}
			Visits[fmt.Sprintf("%d,%d", x,y)] = true
		}else {
			switch string(char) {
			case ">":
				x2++
			case "<":
				x2--
			case "^":
				y2++
			case "v":
				y2--
			}
			Visits[fmt.Sprintf("%d,%d", x2,y2)] = true
		}
		
	}
	noOfHouses = len(Visits)
	return noOfHouses
}

func main() {
	str, err := os.ReadFile("input.txt")
	fmt.Println("length of str =>", len(str))
	if err == nil{
		noOfHouses := PresentDeliver(string(str))
		fmt.Println(noOfHouses)
		visits := SantaAndRobo(string(str))
		fmt.Println("unique visits",visits)
	}else {
		fmt.Println("Error \n", err)
		return
	}
	
}
