package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func check(x int) bool {
	s := strconv.Itoa(x)
	h := len(s) / 2
	return s[:h] == s[h:]
}

func check2(x int) bool {
	s := strconv.Itoa(x)
	h := len(s) / 2
	for i := 1; i <= h; i += 1 {	
		if strings.Repeat(s[:i], len(s) / i) == s {
			return true
		}
	}
	return false
}

func main() {
	input, err := os.Open("input")	
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(input)

	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		ranges := strings.Split(line, ",")
		for _, r := range ranges {
			nums := strings.Split(r, "-")
			s, err := strconv.Atoi(nums[0])
			if err != nil {
				panic(err)
			}
			e, err := strconv.Atoi(nums[1])
			if err != nil {
				panic(err)
			}
			for i := s; i <= e; i++ {
				if check2(i) {
					ans += i
				}
			}
		} 
	}
	fmt.Println(ans)
}
