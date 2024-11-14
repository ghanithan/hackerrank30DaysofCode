package main

import (
	"bufio"
	"fmt"
	"os"

	// "os/signal"
	//     "syscall"

	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	numberOfInput := 0
	fmt.Scanf("%d", &numberOfInput)

	phoneBook := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < numberOfInput; i++ {
		scanner.Scan()
		nameAndNumber := scanner.Text()
		splitInput := strings.Split(nameAndNumber, " ")
		phoneBook[splitInput[0]] = splitInput[1]
	}

	// --- comlex solution using channel
	// sig := make(chan bool)

	// for {

	// 	go func(sig chan bool) {
	// 		if scanner.Scan() {
	// 			name := scanner.Text()
	// 			if number, ok := phoneBook[name]; ok {
	// 				fmt.Printf("%s=%s\n", name, number)
	// 			} else {
	// 				fmt.Println("Not found")
	// 			}
	// 			sig <- true
	// 		} else {
	// 			sig <- false
	// 		}
	// 	}(sig)
	// 	if osSignal := <-sig; !osSignal {
	// 		break
	// 	}
	// }

	// --- simple solution

	for scanner.Scan() {

		name := scanner.Text()
		if number, ok := phoneBook[name]; ok {
			fmt.Printf("%s=%s\n", name, number)
		} else {
			fmt.Println("Not found")
		}

	}

}
