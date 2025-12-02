package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	// fmt.Println("tuturuu")
	input, err := os.Open("input")	
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(input)
	
	dial := 50
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		num, err := strconv.Atoi(line[1:])
		num %= 100
		if err != nil {
			panic(err)
		}
		if line[0] == 'L' {
			dial -= num
			if dial < 0 {
				dial += 100
			}
		} else {
			dial += num
			if dial >= 100 {
				dial -= 100
			}
		}
		if dial == 0 {
			ans += 1
		}
	}
	fmt.Println(ans)
}
