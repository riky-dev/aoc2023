package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

func main() {
	points := 0

	file := readFile()

	lines := strings.Split(file, "\n")

	timeLine := lines[0]
	distanceLine := lines[1]

	times := strings.Fields(timeLine)[1:]
	distances := strings.Fields(distanceLine)[1:]
	
	time, _ := strconv.Atoi(strings.Join(times, ""))
	distance, _ := strconv.Atoi(strings.Join(distances, ""))

	for j := 0; j < time; j++ {
		remaining := time - j
		travelled := j * remaining

		if travelled > distance {
			points++
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
