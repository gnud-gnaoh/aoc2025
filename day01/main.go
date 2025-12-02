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
		if err != nil {
			panic(err)
		}
		// fmt.Println(num)
		add := +1
		if line[0] == 'L' {
			add = -1	
		}
		
		for num > 0 {
			dial += add
			if dial < 0 {
				dial = 99 
			}
			if dial >= 100 {
				dial = 0
			}
			if dial == 0 {
				ans += 1
			}
			num -= 1
		}
	}
	fmt.Println(ans)
}
