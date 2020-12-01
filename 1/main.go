package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func multiply(array []int) int {
	var result int
	for i := range array {
		if i == 0 {
			result = array[i]
			continue
		}
		result *= array[i]
	}
	return result
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		check(err)
		numbers = append(numbers, number)
	}

	//sort.Ints(numbers)

	partial := []int{}
	sumIt(numbers, 2020, partial)
}

func sumIt(numbers []int, target int, partial []int) {
	s := sum(partial)

	if s == target {
		if len(partial) == 2 {
			fmt.Printf("total: %d\n", multiply(partial))
		}
		if len(partial) == 3 {
			fmt.Printf("total: %d\n", multiply(partial))
		}
		fmt.Printf("sum(%d)=%d\n", partial, target)
	}

	if s >= target {
		return
	}

	for i, n := range numbers {
		remainingNumbers := numbers[i:]
		sumIt(remainingNumbers, target, append(partial, n))
	}

}
