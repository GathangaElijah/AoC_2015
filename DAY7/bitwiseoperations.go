package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	// data := FileReader()
	// fmt.Println(data[:10])
	var signalData [][]string
	signalData = [][]string{
		{"123", "->" ,"x"},
		{"456", "->", "y"},
		{"x", "AND", "y", "->", "d"},
		{"x", "OR", "y", "->", "e"},
		{"x", "LSHIFT", "2", "->", "f"},
		{"y", "RSHIFT", "2", "->", "g"},
		{"NOT", "x", "->", "h"},
		{"NOT", "y", "->", "i"},
	}
	signalMap := CircuitCommands(signalData)
	fmt.Println(signalMap)

	realData := FileReader()
	realMap := CircuitCommands(realData)
	fmt.Println("This is the realmap: ", realMap)
	if valueA, ok := realMap["a"]; ok{
		fmt.Println("This is the value of A: ", valueA)
	} else {
		fmt.Println("a is not in the map")
		return
	}
}

func FileReader()[][]string{
	file, err := os.Open("input.txt")
	if err != nil{
		fmt.Println("Error opening file!\n", err)
		return nil
	}
	defer file.Close()

	var circuitData [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		lineData := strings.Split(line, " ")
		circuitData = append(circuitData, lineData)
	}
	return circuitData
}

func CircuitCommands(circuitData [][]string) map[string]uint16{
	signalMap := make(map[string]uint16)
    unresolved := make(map[string][]string)

    // Parse the instructions into the unresolved map
    for _, signalInfo := range circuitData {
        if len(signalInfo) == 3 {
            // Direct signal assignment (e.g., 123 -> x)
            unresolved[signalInfo[2]] = signalInfo
        } else if len(signalInfo) == 4 {
            // NOT gate (e.g., NOT x -> h)
            unresolved[signalInfo[3]] = signalInfo
        } else if len(signalInfo) == 5 {
            // Gate (e.g., x AND y -> d)
            unresolved[signalInfo[4]] = signalInfo
        }
    }

    // Function to recursively resolve the signal for a wire
    var resolveSignal func(wire string) uint16
    resolveSignal = func(wire string) uint16 {
        // If the wire is already in the signalMap, return its value
        if val, ok := signalMap[wire]; ok {
            return val
        }

        // Check if the wire is a number
        if n, err := strconv.Atoi(wire); err == nil {
            return uint16(n)
        }

        // Fetch the unresolved instruction for the wire
        instruction := unresolved[wire]

        var signal uint16

        if len(instruction) == 3 {
            // Direct signal assignment
            signal = resolveSignal(instruction[0])

        } else if len(instruction) == 4 && instruction[0] == "NOT" {
            // NOT gate
            signal = ^resolveSignal(instruction[1])

        } else if len(instruction) == 5 {
            // Other gates: AND, OR, LSHIFT, RSHIFT
            operand1 := resolveSignal(instruction[0])
            operand2 := resolveSignal(instruction[2])

            switch instruction[1] {
            case "AND":
                signal = operand1 & operand2
            case "OR":
                signal = operand1 | operand2
            case "LSHIFT":
                shift, _ := strconv.Atoi(instruction[2])
                signal = operand1 << uint16(shift)
            case "RSHIFT":
                shift, _ := strconv.Atoi(instruction[2])
                signal = operand1 >> uint16(shift)
            }
        }

        // Cache the computed signal in the map
        signalMap[wire] = signal
        return signal
    }

    // Compute the signal for all the wires in the unresolved map
    for wire := range unresolved {
        resolveSignal(wire)
    }

    return signalMap
}