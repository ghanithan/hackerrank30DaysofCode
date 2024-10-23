package hackerrank30daysofcode

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
)

func TestDay1(t *testing.T) {

	t.Run("test day1", func(t *testing.T) {
		// Create a mock stdin
		r, w, _ := os.Pipe()
		defer r.Close()
		defer w.Close()
		os.Stdin = r

		R, W, _ := os.Pipe()
		defer R.Close()
		defer W.Close()
		os.Stdout = W

		got := ""
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			var buf bytes.Buffer
			io.Copy(&buf, R)
			t.Logf("%s", buf.String())
			got = buf.String()
			wg.Done()
		}()

		input := "Welcome to 30 days of coding!\n"
		// Write test input to the mock stdin
		fmt.Fprintln(w, input)
		Day1()
		want := append([]byte("Hello, World.\n"), input...)

		wg.Wait()
		if string(want) != got {
			t.Errorf("Day()= %q , want %q", got, want)
		} else {
			t.Log("Success")
		}
	})

}
