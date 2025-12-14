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
	board := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		board = append(board, line)
	}
	
	diri := []int{0, 0, -1, -1, -1, +1, +1, +1} 
	dirj := []int{-1, +1, 0, -1, +1, 0, -1, +1} 
	for i, line := range board {
		for j, c := range line {
			if c != '@' {
				continue
			}	
			cnt := 0
			for ii := 0; ii < 8; ii++ {
					ni := i + diri[ii] 
					nj := j + dirj[ii] 
					if 0 <= ni && ni < len(board) && 0 <= nj && nj < len(line) && board[ni][nj] == '@'{
						cnt++ 
					}
			}	
			// fmt.Println(i, j, cnt)
			if cnt < 4 {
				ans++
				// fmt.Println(i, j)
			}
		}
	}
	fmt.Println(ans)
}
