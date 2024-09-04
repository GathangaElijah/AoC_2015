package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// grid := createGrid()
	// str1 := "turn on 0,0 through 9,9"
	// str2 := "toggle 0,0 through 9,0"
	// str3 := "turn off 4,9 through 0,5"
	// fmt.Println("On", Lights(str1,grid))
	// fmt.Println("Toggle", Lights(str2, grid))
	// fmt.Println("Off", Lights(str3, grid))
	// fmt.Println("Final output", Output())
	// Brightness(str1, grid)
	// Brightness("toggle 0,0 through 999,999", grid)
	// fmt.Println("Brightness 1 =>", Brightness(str1, grid))
	// fmt.Println("Brightness 2 =>", Brightness("toggle 0,0 through 999,999", grid))
	fmt.Println("Final brightness =>", Output())
}

// Create the grid
func createGrid() [][]int {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}
	return grid
}

func Lights(s string, grid [][]int) int {
	// Create the grid
	noOfLights := 0

	// Split the string with a space to get the parts
	parts := strings.Split(s, " ")

	// Get the starting and ending coordinates
	var startCoord, endCoord []string
	if parts[0] == "toggle" {
		startCoord = strings.Split(parts[1], ",")
		endCoord = strings.Split(parts[3], ",")
	} else {
		startCoord = strings.Split(parts[2], ",")
		endCoord = strings.Split(parts[4], ",")
	}

	// Covert the coordinates to int
	x1, _ := strconv.Atoi(startCoord[0])
	y1, _ := strconv.Atoi(startCoord[1])
	x2, _ := strconv.Atoi(endCoord[0])
	y2, _ := strconv.Atoi(endCoord[1])

	// Turning on the lights
	for row := x1; row <= x2; row++ {
		for col := y1; col <= y2; col++ {
			if parts[0] == "toggle" {
				if grid[row][col] == 1 {
					grid[row][col] = 0
					noOfLights--
				} else {
					grid[row][col] = 1
					noOfLights++
				}
			} else if parts[1] == "on" {
				if grid[row][col] == 0 {
					grid[row][col] = 1
					noOfLights++
				}
			} else if parts[1] == "off" {
				if grid[row][col] == 1 {
					grid[row][col] = 0
					noOfLights--
				}
			}
		}
	}
	return noOfLights
}

// Changing with brightness
func Brightness(s string, grid [][]int) {
	// Split the string with a space to get the parts
	parts := strings.Split(s, " ")

	// Get the starting and ending coordinates
	// Get the starting and ending coordinates
	var startCoord, endCoord []string
	if parts[0] == "toggle" {
		startCoord = strings.Split(parts[1], ",")
		endCoord = strings.Split(parts[3], ",")
	} else {
		startCoord = strings.Split(parts[2], ",")
		endCoord = strings.Split(parts[4], ",")
	}

	// Convert the coordinates to int
	x1, _ := strconv.Atoi(startCoord[0])
	y1, _ := strconv.Atoi(startCoord[1])
	x2, _ := strconv.Atoi(endCoord[0])
	y2, _ := strconv.Atoi(endCoord[1])

		// Update brightness based on the command
	for row := x1; row <= x2; row++ {
		for col := y1; col <= y2; col++ {
			if parts[0] == "toggle" {
				grid[row][col] += 2
			} else if parts[1] == "on" {
				grid[row][col] += 1
			} else if parts[1] == "off" {
				grid[row][col] -= 1
				if grid[row][col] < 0 {
					grid[row][col] = 0
				}
			}
		}
	}

	// Covert the coordinates to int

	// Check if coordinates are within grid bounds
	// if x1 < 0 || x1 >= 1000 || y1 < 0 || y1 >= 1000 || x2 < 0 || x2 >= 1000 || y2 < 0 || y2 >= 1000 {
	//     fmt.Println("Error: Coordinates out of bounds!")
	//     return
	// }

	// Controlling the brightness
	// fmt.Print(grid)
}

func calcTotalBright(grid [][]int) int {
	// Calculating the total brightness
	var totalBrightness int
	for _, row := range grid {
		for _, brtValue := range row {
			if brtValue >= 0 {
				totalBrightness += brtValue
			}
		}
	}
	fmt.Println(totalBrightness)
	return totalBrightness
}

func Output() int {
	grid := createGrid()
	file, _ := os.Open("input.txt")
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		// Brightness
		Brightness(line, grid)
	}
	totalBrightness := calcTotalBright(grid)
	return totalBrightness
}
