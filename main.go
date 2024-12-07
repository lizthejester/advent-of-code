package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	left := []int{}
	right := []int{}
	numbers := []int{}
	for scanner.Scan() {
		strconv.Atoi(scanner.Text())
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	for i := range numbers {
		num := numbers[i]
		if i%2 == 0 {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	slices.Sort(left)
	slices.Sort(right)

	total := 0
	for i := range left {
		dist := Distance(left[i], right[i])
		total = total + dist
	}
	fmt.Println(total)
}

func Distance(A int, B int) int {
	num := A - B
	if num < 0 {
		num = num * -1
	}
	return num
}
