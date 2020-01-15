package main

import (
	"fmt"
	"strconv"
)

func Day4(start, end int) (count int) {
	for ; start <= end; start++ {
		str := strconv.Itoa(start)
		adjacent := false
		increasing := true
		prevS := ""
		prevI := -1
		for _, s := range str {
			curS := string(s)
			curI, _ := strconv.Atoi(curS)
			if prevS == curS {
				adjacent = true
			}
			if curI < prevI {
				increasing = false
				break
			}
			prevS = curS
			prevI = curI
		}
		if adjacent && increasing {
			count++
		}
	}
	return
}

func Day4Part2(start, end int) (count int) {
	for ; start <= end; start++ {
		str := strconv.Itoa(start)
		adjacent := false
		groupCount := 0
		increasing := true
		prevS := ""
		prevI := -1
		for _, s := range str {
			curS := string(s)
			curI, _ := strconv.Atoi(curS)
			if prevS == curS {
				groupCount++
			} else {
				if groupCount == 1 {
					adjacent = true
				}
				groupCount = 0
			}
			if curI < prevI {
				increasing = false
				break
			}
			prevS = curS
			prevI = curI
		}
		if groupCount == 1 {
			adjacent = true
		}
		if adjacent && increasing {
			count++
		}
	}
	return
}

func main() {
	count := Day4(168630, 718098)
	fmt.Println(count)
	count = Day4Part2(168630, 718098)
	fmt.Println(count)
}
