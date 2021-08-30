package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var Delimiter = "SuperCoolDelimiter"

func main() {
	fmt.Println("Enter your name")
	fmt.Print("-> ")
	CommandScan := bufio.NewScanner(os.Stdin)
	CommandScan.Scan()
	input := []byte(CommandScan.Text())

	stub, err := ioutil.ReadFile("Stub.exe")
	if err != nil || len(stub) == 0 {
		panic(err)
	}

	data := append(stub, append([]byte(Delimiter), input...)...)

	err = ioutil.WriteFile("BuiltStub.exe", data, 0755)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finished!")
}
