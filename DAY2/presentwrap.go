package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PresentWrap(l, w, h int) int {
	var totalArea int
	surfaceArea := (2 * l * w) + (2 * w * h) + (2 * h * l)
	side1 := l*w
	side2 := w*h
	side3 := l*h

	slack := side1
	
	if side2 < slack{
		slack = side2
	}
	if side3 < slack{
		slack = side3
	}
	totalArea = surfaceArea + slack

	return totalArea
}

func fileReader() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file, \n", err)
	}
	defer file.Close()

	all := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var digitSlice []int
		line := scanner.Text()
		chars := strings.Split(line, "x")
		for _, char := range chars {
			digit, _ := strconv.Atoi(char)
			digitSlice = append(digitSlice, digit)
		}
		all = append(all, digitSlice)
	}
	return all
}

func main() {
	area := PresentWrap(2, 3, 4)
	fmt.Println("This is area", area)

	var totalArea int
	var l, w, h int
	Digits := fileReader()
	// fmt.Println(Digits)
	for _, digits := range Digits {
		l = digits[0]
		w = digits[1]
		h = digits[2]
		area := PresentWrap(l, w, h)
		fmt.Printf("Digits %v, Area =>%v \n", digits, area)
		totalArea += area
	}
	fmt.Println(totalArea)
}
