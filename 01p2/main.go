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
	numStrings := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	lines := strings.Split(str, "\n")
	sum := 0

	for _, line := range lines {	
		strSplit := strings.Split(line, "")

		fNSI := len(line)
		rep := 0

		for i, v := range numStrings {
			ind := strings.Index(line, v)
			if (ind < fNSI && ind != -1) {
				fNSI = ind
				rep = i + 1
			}
		}

		if (fNSI != len(line)){
			strSplit[fNSI] = strconv.Itoa(rep)
		}

		lNSI := -1
		rep = 0

		for i, v := range numStrings {
			ind := strings.LastIndex(line, v)
			if (ind > lNSI) {
				lNSI = ind
				rep = i + 1
			}
		}

		if (lNSI != -1){
			strSplit[lNSI] = strconv.Itoa(rep)
		}

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
