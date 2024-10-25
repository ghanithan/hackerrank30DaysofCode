package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkWeird(number int32) {
	if number%2 == 0 {
		// This is even number
		switch {
		case number >= 2 && number <= 5:
			fmt.Println("Not Weird")
		case number >= 6 && number <= 20:
			fmt.Println("Weird")
		case number > 20:
			fmt.Println("Not Weird")
		default:
			fmt.Println("default")
		}
	} else {
		fmt.Println("Weird")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)

	checkWeird(N)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
