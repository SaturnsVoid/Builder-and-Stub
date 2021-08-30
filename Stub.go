package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var Delimiter = "SuperCoolDelimiter"

func main() {
	bytes, err := ioutil.ReadFile(os.Args[0])
	if err != nil || len(bytes) == 0 {
		panic(err)
	}
	position := strings.LastIndex(string(bytes), Delimiter)
	if position == -1 {
		os.Exit(1)
	}
	position += len(Delimiter)

	data := bytes[position:]

	fmt.Println("How are you,", string(data)+"?")

	for {
		time.Sleep(50 * time.Millisecond)
	}
}
