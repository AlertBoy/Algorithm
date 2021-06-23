package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	s := line[2:]
	var ret float64
	ret = 0
	fmt.Println(s)
	m := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "A": 10, "B": 11, "C": 12, "D": 13, "E": 14, "F": 15}
	cnt := 0
	for i := len(line) - 1; i > 1; i-- {
		ret += float64(m[string(line[i])]) * math.Pow(16, float64(cnt))
		cnt++
	}
	fmt.Println(ret)

}
