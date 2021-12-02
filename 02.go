package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var horizontal int
	var depth int
	var result int
	type outputData struct {
		direction string
		value     int
	}
	var allOutputData []outputData

	b, err := ioutil.ReadFile("input.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	inputArray := strings.SplitN(str, "\r\n", -1)
	for _, j := range inputArray {
		tempPremenna := strings.SplitN(j, " ", -1)
		tempInt, err := strconv.ParseInt(tempPremenna[1], 10, 64)
		if err != nil {
			fmt.Println("ERROR CONVERTING TO INT: ", tempInt)
			return
		}

		var tempoutput outputData = outputData{
			tempPremenna[0],
			int(tempInt),
		}
		allOutputData = append(allOutputData, tempoutput)
	}

	for _, j := range allOutputData {
		switch j.direction {
		case "forward":
			horizontal += j.value
		case "up":
			depth -= j.value
		case "down":
			depth += j.value
		default:
			fmt.Println("ERROR PARSING VALUE")
		}
	}
	result = horizontal * depth
	fmt.Println("\n\nTHE RESULT IS: ", result)

}
