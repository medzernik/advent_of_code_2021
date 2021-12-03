package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var gammaRate string
	var epsilonRate string
	//var result int
	var valueToAvg float64

	var currentSymbol string

	b, err := ioutil.ReadFile("input.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	inputArray := strings.SplitN(str, "\r\n", -1)

	//fmt.Print(inputArray[0][0:1])
	for i := 0; i < 12; i++ {
		valueToAvg = 0
		for _, j := range inputArray {

			currentSymbol = j[i : i+1]

			switch currentSymbol {
			case "0":
				valueToAvg += 0
			case "1":
				valueToAvg += 1

			}

		}
		if valueToAvg/float64(len(inputArray)) >= 0.5 {
			gammaRate += "1"
			//gammaRate = append(gammaRate, 1)
		} else {
			gammaRate += "0"
			//gammaRate = append(gammaRate, 0)
		}

	}
	for i := 0; i < 12; i++ {
		valueToAvg = 0
		for _, j := range inputArray {

			currentSymbol = j[i : i+1]

			switch currentSymbol {
			case "0":
				valueToAvg += 0
			case "1":
				valueToAvg += 1

			}

		}
		if valueToAvg/float64(len(inputArray)) >= 0.5 {
			epsilonRate += "0"
			//epsilonRate = append(epsilonRate, 0)
		} else {
			epsilonRate += "1"
			//epsilonRate = append(epsilonRate, 1)
		}

	}
	gammaNumber, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonNumber, _ := strconv.ParseInt(epsilonRate, 2, 64)

	fmt.Println("FINAL VALUE: ", gammaNumber*epsilonNumber)
}
