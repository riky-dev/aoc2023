package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	points := 0

	file := readFile()

	lines := strings.Split(file, "\n")

	for _, l := range lines {
		if l == "" {
			break
		}

		nums := strings.Split(l, ": ")[1]

		splitNums := strings.Split(nums, " ")

		var wNum []int
		var hNum []int

		isCheckingWinning := true
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
		}

		var winningNumbers int
		for _, n1 := range wNum {
			for _, n2 := range hNum {
				if n1 == n2 {
					winningNumbers++
				}
			}
		}

		points += int(math.Pow(2, float64(winningNumbers-1)))

		// fmt.Println(wNum, hNum)
	}
	fmt.Println(points)
}

func readFile() string {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	return string(data)
}
