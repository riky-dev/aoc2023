package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"strconv"
)

func main() {
	// points := 0

	file := readFile()

	lines := strings.Split(file, "\n")

	for _, l := range(lines) {
		if l == ""{
			break
		}
		nums := strings.Split(l, ": ")[1]

		fmt.Println(nums)

		var wNum []int
		var hNum []int

		for _, n := range(nums) {
			if n == ' ' {
				continue
			}
			nValue, err := strconv.Atoi(string(n))
			isCheckingWinning := true
			if err == nil {
				if isCheckingWinning {
					wNum = append(wNum, nValue)
				} else {
					hNum = append(hNum, nValue)
				}
			} else {
				fmt.Println(err)
				isCheckingWinning = false
			}
		}
		fmt.Println(wNum, hNum)
	}
}

func readFile() string {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	return string(data)
}
