package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("input.txt")
    str := string(b)
	lines := strings.Split(str, "\n")
	sum := 0

	for _, line := range lines {	
		strSplit := strings.Split(line, "")
		var fN *int
		var lN *int

		for _, v := range strSplit {
			if i, err := strconv.Atoi(v); err == nil {
				if fN == nil {
					fN = &i
					continue
				}
					lN = &i
			}
		}

		if lN == nil {
			lN = fN
		}

		fNlN := *fN * padding(*lN) + *lN
		sum += fNlN
	}

	log.Print(sum)
}

func padding(n int) int {
	var p int = 10

	for p < n {
		p *= 10
	}

	return p
}
