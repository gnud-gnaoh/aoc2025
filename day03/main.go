package main

import (
	"fmt"
	"os"
	"bufio"
	// "strconv"
	// "strings"
)

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
		
		mx := 0
		for i := 0; i < len(line); i += 1 {
			for j := i + 1; j < len(line); j += 1 {
				mx = max(mx,(int(line[i]) - int('0')) * 10 + (int(line[j]) - int('0')))
			}
		}
		ans += mx
	}
	fmt.Println(ans)
}
