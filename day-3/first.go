package main

import (
        "fmt"
        "io/ioutil"
        "os"
        "strconv"
        "strings"
	"unicode"
)

func main() {
        t := readFile()

        var sum int

	lines := strings.Split(t, "\n")

        for i, v := range lines {
                if v == "" {
                        break
                }

		elements := strings.Split(v, ".")

		for j, e := range(elements){
			var isPartNumber bool

			n, err := strconv.Atoi(e)
			if err == nil {
				// Up above
				if i > 0 {
					fmt.Println("PREVIOUS LINE: ", lines[i-1])
					fmt.Println("CURRENT  LINE: ", lines[i])
					//fmt.Println("CHAR ABOVE: ", rune(lines[i-1][j]))
					if string(lines[i-1][j]) != "" && !unicode.IsDigit(rune(lines[i-1][j])) {
						isPartNumber = true
					}
				// Up left
				// Up right
				}
//				if i < len(v)-1 {
//					if string(lines[i+1][j]) != "" && !unicode.IsDigit(rune(lines[i+1][j])) {
//						isPartNumber = true
//					}
//				}
				// Down below
				// Down left
				// Down right
			}

			if isPartNumber {
				// fmt.Println(v)
				sum += n
			}
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
