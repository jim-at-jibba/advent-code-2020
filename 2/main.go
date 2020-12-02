package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type deconstructedPassword struct {
	max      int
	min      int
	match    string
	password string
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	filterPasswords(lines)
}

func filterPasswords(list []string) {

	pp := []deconstructedPassword{}

	for _, str := range list {
		pp = append(pp, formatPasswordString(str))
	}

	checkPasswordValidity(pp)
}

func formatPasswordString(p string) deconstructedPassword {

	partsSlice := strings.Fields(p)
	maxMin := strings.Split(partsSlice[0], "-")
	min, err := strconv.Atoi(maxMin[0])
	check(err)
	max, err := strconv.Atoi(maxMin[1])
	check(err)
	match := strings.Split(partsSlice[1], ":")
	return deconstructedPassword{max: max, min: min, match: match[0], password: partsSlice[2]}
}

func checkPasswordValidity(passwords []deconstructedPassword) {
	// tries 221
	// 1. check password has match in it
	// 2. check

	correct := []deconstructedPassword{}

	for _, pw := range passwords {

		count := strings.Count(pw.password, pw.match)
		//correct = append(correct, pw)
		if count >= pw.min && count <= pw.max {
			fmt.Printf("%v %v %v min: %v max: %v\n", pw.password, pw.match, strings.Count(pw.password, pw.match), pw.min, pw.max)
			correct = append(correct, pw)
		}
	}

	fmt.Println(len(correct))
}
