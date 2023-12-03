package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	// "unicode"
)

type numberData struct {
	Line  int
	Start int
	End   int
	Value []int
}

func main() {
	lines := strings.Split(readFile(), "\n")

	fmt.Println(getNumbersDetails(lines))

}

func getNumbersDetails(lines []string) (numData []numberData) {
	currentNum := numberData{
		Line:  -1,
		Start: -1,
		End:   -1,
		Value: make([]int, 3),
	}
	var numValue []int
	for i, l := range lines {
		for j, c := range l {
			digit, err := strconv.Atoi(string(c))
			if err == nil {
				if currentNum.Line < 0 && currentNum.Start < 0 {
					currentNum.Line = i
					currentNum.Start = j
				}
				numValue = append(numValue, digit)
				_, err := strconv.Atoi(string(l[j+1]))
				if err != nil {
					currentNum.End = i
				}
			} else {
				if currentNum.End >= 0 {
					numData = append(numData, currentNum)
				} else {
					numData = append(numData, currentNum)
					currentNum = numberData{
						Line:  -1,
						Start: -1,
						End:   -1,
						Value: make([]int, 3),
					}
				}

			}

		}
	}
	return
}

func readFile() string {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	return string(data)
}
