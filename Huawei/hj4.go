package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	len := len(line)
	cnt := len / 8
	let := len % 8
	for i := 0; i <= cnt; i++ {
		fmt.Print(string(line[8*i : 8*(i+1)]))
	}
	if let > 0 {
		for i := 8 * cnt; i < 8*cnt+let; i++ {
			fmt.Print(string(line[i]))
		}
		for i := 0; i < 8-let; i++ {
			fmt.Print("0")
		}
	}

}
