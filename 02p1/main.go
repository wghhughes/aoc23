package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var balls = make(map[string]int)

func main() {
	balls["red"] = 12
	balls["green"] = 13
	balls["blue"] = 14
	b, _ := os.ReadFile("input.txt")
    str := string(b)
	lines := strings.Split(str, "\n")
	sum := 0

	for i, line := range lines {
		sum += loop(line, i+1)
	}

	log.Print(sum)
}

func loop(line string, lineNo int) int {
	s1 := strings.Split(line, ": ")
	s2 := s1[1]
	s3 := strings.ReplaceAll(s2, ";", ",")
	s4 := strings.Split(s3, ", ")

	for _, s := range s4 {
		ss := strings.Split(s, " ")
		i, _ := strconv.Atoi(ss[0])
		for k, v := range balls {
			if strings.Contains(ss[1], k) && i > v {
				return 0
			}
		}
	}

	return lineNo
}