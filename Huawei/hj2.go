package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line1, _, _ := r.ReadLine()
	line2, _, _ := r.ReadLine()
	lower := strings.ToLower(string(line1))
	lower2 := strings.ToLower(string(line2))
	cnt := 0
	for i := 0; i < len(lower); i++ {
		if string(lower[i]) == lower2 {
			cnt++
		}
	}
	fmt.Print(cnt)
}
