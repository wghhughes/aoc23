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
		sum += loop(line)
	}

	log.Print(sum)
}

func loop(line string) int {
	s1 := strings.Split(line, ": ")
	s2 := s1[1]
	s3 := strings.ReplaceAll(s2, ";", ",")
	s4 := strings.Split(s3, ", ")

	balls := make(map[string]int)
	balls["red"] = 0
	balls["green"] = 0
	balls["blue"] = 0

	for _, b := range [3]string{"red", "blue", "green"} {
		for _, s := range s4 {
			ss := strings.Split(s, " ")
			i, _ := strconv.Atoi(ss[0])

			if strings.Contains(ss[1], b) && i > balls[b] {
				balls[b] = i
			}
		}
	}
	
	return balls["red"] * balls["green"] * balls["blue"]
}