package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	//fmt.Println(s)
	strings.TrimRight(s, " ")
	if len(s) == 0 {
		fmt.Printf("%d\n", 0)
	}
	cnt := -1
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			cnt++
		} else {
			break
		}
	}
	fmt.Print(cnt)
}
