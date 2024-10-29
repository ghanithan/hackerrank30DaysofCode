package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	numberOfLines := 0
	fmt.Fscanf(os.Stdin, "%d", &numberOfLines)

	words := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < int(numberOfLines); i++ {
		scanner.Scan()
		input := scanner.Text()
		words = append(words, input)

	}
	for _, word := range words {
		even := make([]rune, 0)
		odd := make([]rune, 0)
		for i, v := range word {
			if i%2 != 0 {
				even = append(even, v)
			} else {
				odd = append(odd, v)
			}
		}
		fmt.Printf("%s %s\n", string(odd), string(even))

	}

}
