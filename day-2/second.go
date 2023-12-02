package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	t := readFile()

	var sum int

	for _, v := range strings.Split(t, "\n") {
		if v == "" {
			break
		}

		gameNumToCheck := v[5:8]
		gameNumStr := strings.Replace(strings.Replace(gameNumToCheck, ":", "", -1), " ", "", -1)
		// gameNum, _ := strconv.Atoi(gameNumStr)

		v := strings.Replace(strings.Replace(v, ", ", ";", -1), "; ", ";", -1)[4+len(gameNumStr)+3:]

		sets := strings.Split(v, ";")

		var maxCubes [3]int

		for _, v := range sets {
			v := strings.Split(v, " ")
			num, _ := strconv.Atoi(string(v[0]))
			color := v[1]
			switch color {
			case "red":
				if maxCubes[0] < num {
					maxCubes[0] = num
				}
			case "green":
				if maxCubes[1] < num {
					maxCubes[1] = num
				}
			case "blue":
				if maxCubes[2] < num {
					maxCubes[2] = num
				}
			}
		}

		sum += maxCubes[0] * maxCubes[1] * maxCubes[2]
	}
	fmt.Println(sum)
}

func readFile() string {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	return string(data)
}
