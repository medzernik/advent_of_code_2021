package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	var counter uint64

	str := string(b) // convert content to a 'string'
	inputArray := strings.SplitN(str, "\n", -1)

	fmt.Println(inputArray[1])

	for i := range inputArray{
		if i+1 < len(inputArray) {
			if inputArray[i+1] > inputArray[i] {
				counter += 1
			}

		} else if inputArray[i] > inputArray[i-1] {
			counter += 1
		}
	}
	fmt.Println("The final result is: ", counter)

}
