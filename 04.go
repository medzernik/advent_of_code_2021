package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Value struct {
	Value   string
	Correct bool
}

type Board [5][5]Value

var Boards [602]Board

var generatedNumber []string

var counter int
var SplitNumsArray []string
var winningBoard int
var winningNumbers []string
var stopSearching bool
var winningNumbersToCount []int
var otherNumbersToCount []int
var result int
var lastNumberCalled string

func main() {
	b, err := ioutil.ReadFile("input.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	tempArray := strings.SplitN(str, "\r\n", -1)

	tempGeneratedString := strings.Split(tempArray[0], ",")
	for _, j := range tempGeneratedString {
		generatedNumber = append(generatedNumber, j)
	}

	for i := 2; i < len(tempArray); i++ {
		tempSplitNumsArray := strings.SplitN(tempArray[i], " ", -1)
		for _, j := range tempSplitNumsArray {
			SplitNumsArray = append(SplitNumsArray, j)
		}

	}

	var x, y int

	for _, j := range SplitNumsArray {
		if j == "" {
			continue
		}

		if x > 4 {
			y += 1
			x = 0

		}
		if y > 4 {
			y = 0
			x = 0
			counter += 1
		}

		Boards[counter][y][x].Value = j
		x += 1
	}
	generatedNumberToPick := generatedNumber[:]

	for _, q := range generatedNumberToPick {
		for z := range Boards {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if Boards[z][i][j].Value == q {
						Boards[z][i][j].Correct = true
						lastNumberCalled = q
					}
				}
			}
		}

		for z := range Boards {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if stopSearching == false && (Boards[z][0][j].Correct == true && Boards[z][1][j].Correct == true && Boards[z][2][j].Correct == true && Boards[z][3][j].Correct == true && Boards[z][4][j].Correct == true) || Boards[z][i][0].Correct == true && Boards[z][i][1].Correct == true && Boards[z][i][2].Correct == true && Boards[z][i][3].Correct == true && Boards[z][i][4].Correct == true {
						winningBoard = z
						winningNumbers = append(winningNumbers, Boards[z][0][j].Value, Boards[z][1][j].Value, Boards[z][2][j].Value, Boards[z][3][j].Value, Boards[z][4][j].Value)
						stopSearching = true
						goto FINISHED
					}
				}
			}
		}
	}

FINISHED:

	fmt.Println(Boards[0:2])

	fmt.Println("Winning board: ", winningBoard)
	fmt.Println("Winning numbers: ", winningNumbers)

	for _, j := range winningNumbers {
		tempVar, err := strconv.ParseInt(j, 10, 64)
		if err != nil {
			fmt.Println("ERROR CONVERTING NUMBER: ", err)
			return
		}
		winningNumbersToCount = append(winningNumbersToCount, int(tempVar))
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {

			if Boards[winningBoard][i][j].Correct != true {
				tempVar, err := strconv.ParseInt(Boards[winningBoard][i][j].Value, 10, 64)
				if err != nil {
					fmt.Println("ERROR CONVERTING NUMBER: ", err)
					return
				}
				otherNumbersToCount = append(otherNumbersToCount, int(tempVar))
			}

		}
	}

	for _, i := range otherNumbersToCount {
		result += i
	}

	lastNumberCalledInt, err := strconv.ParseInt(lastNumberCalled, 10, 64)
	if err != nil {
		fmt.Println("ERROR CONVERTING NUMBER: ", err)
		return
	}
	fmt.Println("LAST NUMBER DRAWN: ", lastNumberCalled)

	result *= int(lastNumberCalledInt)
	fmt.Println("RESULT IS: ", result)
}
