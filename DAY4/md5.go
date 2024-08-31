package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func Md5Hash(s string) int {
	// var result int
	start := 0

	for {
		// Creating a new md5 hash
		hash := md5.New()
		// creating the input data for the hash
		inputData := s + strconv.Itoa(start) 
		// Write the input to the hash
		hash.Write([]byte(inputData))
		// Compute the hash and get the resulting bytes
		hashBytes := hash.Sum(nil)

		// Convert the bytes into hexadecimal string
		hashString := hex.EncodeToString(hashBytes)
		fmt.Println("This is hash for start is =>", hashString)
		// Count the number of "0"

		if hashString[:6] == "000000"{
			return start
		}
	
		start ++

	}
}

func main() {
	s := "ckczppom"
	ans := Md5Hash(s)
	fmt.Println(ans)
}
