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

func Ribbon(l, w, h int)int{
	faceA := (2*l) + (2*w)
	faceB := (2*w) + (2*h)
	faceC := (2*l) + (2*h)

	perimeter := faceA
	if faceB < perimeter{
		perimeter = faceB
	}
	if faceC < perimeter{
		perimeter = faceC
	}
	volume := l * w * h 
	// Total feet required
	totalArea := volume + perimeter
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
	var totalArea int
	var totalRibbonlen int
	var l, w, h int
	Digits := fileReader()
	// fmt.Println(Digits)
	for _, digits := range Digits {
		l = digits[0]
		w = digits[1]
		h = digits[2]
		area := PresentWrap(l, w, h)
		totalArea += area
		ribbonLength := Ribbon(l, w, h)
		totalRibbonlen += ribbonLength
	}
	fmt.Println("This is the total Area needed", totalArea)
	fmt.Println("This is the total ribbon length =>", totalRibbonlen)
}
