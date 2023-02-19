package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	ting := strings.Split(string(file), "\n")

	antnum, err := strconv.Atoi(ting[0])
	if err != nil {
		log.Fatal("Invalid num of ants")
	}
	fmt.Println(antnum)

	type Room struct {
		name  string
		links []*Room
	}

	room := make(map[string]*Room)

	var start, end *Room

	for i, line := range ting {
		if strings.HasPrefix(line, "##start") {
			name := strings.Fields(ting[i+1])
			start = &Room{name: name[0]}
			continue
		}

		if strings.HasPrefix(line, "##end") {
			name := strings.Fields(ting[i+1])
			end = &Room{name: name[0]}
			continue
		}
		if strings.HasPrefix(line, "#") || i == 0 {
			continue
		}

		
	}
}
