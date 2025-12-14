package main

import (
	"fmt"
	"os"
	"bufio"
	// "strconv"
	// "strings"
)

type Pair struct {
	first int
	second int
}

func main() {
	input, err := os.Open("input")	
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(input)

	ans := 0
	board := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		board = append(board, []rune(line))
	}
	
	diri := []int{0, 0, -1, -1, -1, +1, +1, +1} 
	dirj := []int{-1, +1, 0, -1, +1, 0, -1, +1} 
	
	for {
		todelete := []Pair{}
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
					todelete = append(todelete, Pair{i, j})
					// fmt.Println(i, j)
				}
			}
		}
		
		if len(todelete) == 0 {
			break
		}
		for _, p := range todelete {
			board[p.first][p.second] = '.'
		}
	}
	fmt.Println(ans)
}
