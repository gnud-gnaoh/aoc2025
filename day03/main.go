package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
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
		
		// start from last 12 chars
		cur := line[len(line) - 12:]
		for i := len(line) - 13; i >= 0; i-- {
			nxt := string(line[i]) + cur
			best, e := strconv.Atoi(cur) 
			if e != nil {
				panic(e)
			}
			for j := 0; j < len(nxt); j++ {
				tmp := nxt[:j] + nxt[j + 1:]	
				val, e := strconv.Atoi(tmp)
				if e != nil { panic (e) }
				if val > best {
					best = val
					cur = tmp
				}
			}
		}
		v, e := strconv.Atoi(cur)
		if e != nil { panic(e) }
		ans += v 
	}
	fmt.Println(ans)
}
