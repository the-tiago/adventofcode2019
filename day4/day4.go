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

func main() {
	count := Day4(168630, 718098)
	fmt.Println(count)
}
