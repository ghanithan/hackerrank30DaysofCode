package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var i1 uint64
	var d1 float64
	var s1 string
	// Read and save an integer, double, and String to your variables.
	scanner.Scan()
	i1, _ = strconv.ParseUint(scanner.Text(), 10, 64)

	fmt.Println(i + i1)

	scanner.Scan()
	d1, _ = strconv.ParseFloat(scanner.Text(), 64)
	fmt.Printf("%.1f\n", d+d1)

	scanner.Scan()
	s1 = scanner.Text()

	fmt.Printf("%s%s", s, s1)

	// Print the sum of both integer variables on a new line.

	// Print the sum of the double variables on a new line.

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.

}
