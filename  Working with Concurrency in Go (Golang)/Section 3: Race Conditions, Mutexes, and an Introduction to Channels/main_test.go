package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// var mutex sync.Mutex
// // msg= "Hello, world"

// // wg.Add(2)

// // go updateMessage("Goodbye", &mutex)
// // go updateMessage("Goodbye 1", &mutex)

// // wg.Wait()

// // if msg != "Goodbye" {
// //   t.Error("Incorrect value")
// // }

// func Test_updateMessage(t *testing.T) {
// }

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "$34320.00") {
		t.Error("wrong balance returned")
	}
}
