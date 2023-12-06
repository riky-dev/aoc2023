package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

func main() {
	points := 1

	file := readFile()

	lines := strings.Split(file, "\n")

	timeLine := lines[0]
	distanceLine := lines[1]

	times := strings.Fields(timeLine)[1:]
	distances := strings.Fields(distanceLine)[1:]
	
	for i := 0; i < len(times); i++ {
		wins := 0
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])

		for j := 0; j < t; j++ {
			remaining := t - j
			travelled := j * remaining

			if travelled > d {
				wins++
			}
		}
		if wins > 0 {
			points *= wins	
		}
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