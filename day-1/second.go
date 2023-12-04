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

	values := map[string]string{
		"zero":  "z0o",
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for k, v := range values {
		t = strings.Replace(t, k, v, -1)
	}

	var sum int

	for _, v := range strings.Split(t, "\n") {
		var num []int
		for i := 0; i < len(v); i++ {
			n, err := strconv.Atoi(string(v[i]))
			if err == nil {
				num = append(num, n)
			}
		}
		if len(num) > 0 {
			sum += num[0]*10 + num[len(num)-1]
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
