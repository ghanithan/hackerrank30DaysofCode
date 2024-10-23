package hackerrank30daysofcode

import (
	"bufio"
	"fmt"
	"os"
)

func Day1() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	// fmt.Printf("started")

	writer := os.Stdout

	// Write some data

	buffer := bufio.NewScanner(os.Stdin)
	buffer.Scan()
	err := buffer.Err()
	if err != nil {
		fmt.Println("Error in reading input: ", err)
	}
	input := buffer.Text()
	fmt.Fprintln(writer, "Hello, World.")
	fmt.Fprintln(writer, input)
	// Close the writer
	writer.Close()
}
