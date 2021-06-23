package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	r := bufio.NewReader(os.Stdin)
	line1, _, _ := r.ReadLine()
	N, _ := strconv.Atoi(string(line1))

	for i := 0; i < N; i++ {

		line1, _, _ := r.ReadLine()
		atoi, _ := strconv.Atoi(string(line1))

		m := make(map[int]struct{}, atoi)
		for i := 0; i < atoi; i++ {
			word, _, _ := r.ReadLine()
			i2, _ := strconv.Atoi(string(word))
			m[i2] = struct{}{}
		}
		slice := []int{}
		for i, _ := range m {
			slice = append(slice, i)
		}
		sort.Ints(slice)
		for _, i2 := range slice {
			fmt.Println(i2)
		}
	}

}
