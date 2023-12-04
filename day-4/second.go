package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	scratchcards := 0

	file := readFile()

	lines := strings.Split(file, "\n")

	for i, l := range lines {
		if l == "" {
			break
		}

		updateScratchCards(i, l, lines, &scratchcards)

	}
	fmt.Println(scratchcards)
}

func updateScratchCards(i int, l string, lines []string, s *int) {
	n := findWinningNumbers(l)
	*s++

	if n > 0 {
		for j := 0; j < n; j++ {
			updateScratchCards(i+1+j, lines[i+1+j], lines, &(*s))
		}
	}
}

func findWinningNumbers(l string) int {
	nums := strings.Split(l, ": ")[1]

	splitNums := strings.Split(nums, " ")

	var wNum []int
	var hNum []int

	isCheckingWinning := true
	var winningNumbers int

	for _, v := range splitNums {
		if v == "" {
			continue
		}

		nValue, err := strconv.Atoi(v)
		if err == nil {
			if isCheckingWinning {
				wNum = append(wNum, nValue)
			} else {
				hNum = append(hNum, nValue)
			}
		} else {
			isCheckingWinning = false
		}

		winningNumbers = 0
		for _, n1 := range hNum {
			for _, n2 := range wNum {
				if n1 == n2 {
					winningNumbers++
				}
			}
		}
	}
	return winningNumbers
}

func readFile() string {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	return string(data)
}
