package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	grid := createGrid()
	str1 := "turn on 0,0 through 999,999"
	str2 := "toggle 0,0 through 999,0"
	str3 := "turn off 499,499 through 500,500"
	fmt.Println("On", Lights(str1,grid))
	fmt.Println("Toggle", Lights(str2, grid))
	fmt.Println("Off", Lights(str3, grid))
	fmt.Println("Final output", Output())
}
//Create the grid
func createGrid()[][]int{
	grid := make([][]int, 1000)
	for i := range grid{
		grid[i] = make([]int, 1000)
	}
	return grid
}

func Lights(s string, grid [][]int) int{
	// Create the grid
	// grid := createGrid()
	noOfLights := 0

	// Split the string with a space to get the parts
	parts := strings.Split(s, " ")

	// Get the starting and ending coordinates
	var startCoord ,endCoord []string
	if parts[0] == "toggle"{
		startCoord = strings.Split(parts[1], ",")
		endCoord = strings.Split(parts[3], ",")
	}else {
		startCoord = strings.Split(parts[2], ",")
		endCoord = strings.Split(parts[4], ",")
	}

	// Covert the coordinates to int
	x1, _ := strconv.Atoi(startCoord[0])
	y1, _ := strconv.Atoi(startCoord[1])
	x2, _ := strconv.Atoi(endCoord[0])
	y2, _ := strconv.Atoi(endCoord[1])

	// Turning on the lights

	for row := x1; row <= x2; row ++{
		for col := y1; col <= y2; col++{

			if parts[0] == "toggle"{
				if grid[row][col] == 1{
					grid[row][col] = 0
					noOfLights --
				}else {
					grid[row][col] = 1
					noOfLights ++
				}
			}else if parts[1] == "on"{
				if grid[row][col] == 0{
					grid[row][col] = 1
					noOfLights ++
				}
				
			} else if parts[1] == "off"{
				if grid[row][col] == 1{
					grid[row][col] = 0
					noOfLights --
				}
				
			}
		}
	}
	return noOfLights
}

func Output() int{
	grid := createGrid()
	file, _:= os.Open("input.txt")
	defer file.Close()
	lit := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		no := Lights(line,grid)
		lit += no
	}
	return lit
}