package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	var counter uint
	var sumfirst int
	var sumlast int
	var intArray []int

	str := string(b) // convert content to a 'string'
	inputArray := strings.SplitN(str, "\r\n", -1)


	for _, j := range inputArray {
		temp, err := strconv.ParseInt(j, 10, 64)
		if err != nil {
			fmt.Println("ERROR CONVERTING: ", err)
			return
		}
		intArray = append(intArray, int(temp))
	}
	intArray = append(intArray, 0, 0, 0, 0, 0, 0)

	for i:= 0; i< len(intArray); i+=1 {
		if 2000 >= i+3  {
			sumfirst = intArray[i] + intArray[i+1] + intArray[i+2]
			sumlast = intArray[i+1] + intArray[i+2] + intArray[i+3]

			fmt.Println("FIRST: ", sumfirst)
			fmt.Println("LAST: ", sumlast)

			if sumlast > sumfirst {
				counter += 1
				fmt.Println("TRUE: sum:", counter)

			}
		}

	}
	fmt.Println("THE FINAL ANSWER IS: ", counter)
}
