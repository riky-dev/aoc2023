package main

import (
	"fmt"
	"strconv"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	t := readFile()

	var sum int

	for _, v := range(strings.Split(t, "\n")){
		var num []int
		for i := 0; i < len(v); i++ {
			n, err := strconv.Atoi(string(v[i]))
			if err == nil {
				num = append(num, n)
			}
		}
		if len(num) > 0 {
			sum += num[0] * 10 + num[len(num)-1]
		}
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