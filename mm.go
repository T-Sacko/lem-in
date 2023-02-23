package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	name  string
	links []*Room
}

func main() {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	ting := strings.Split(string(file), "\n")

	antnum, err := strconv.Atoi(ting[0])
	if err != nil {
		log.Fatal("Invalid num of ants")
	}
	fmt.Println("number of ants:", antnum)

	Rooms := make(map[string]*Room)

	var start, end *Room
	var name string

	for i, line := range ting {

		if strings.HasPrefix(line, "##start") {
			name = strings.Fields(ting[i+1])[0]
			start = &Room{name: name}
			Rooms[name] = start

			continue
		}

		if strings.HasPrefix(line, "##end") {
			name = strings.Fields(ting[i+1])[0]
			end = &Room{name: name}
			continue
		}
		if strings.HasPrefix(line, "#") || i == 0 || len(line) == 0 || line[0] == 'L' {
			continue
		}

		if len(strings.Fields(line)) == 3 {

			name = strings.Fields(line)[0]
			if name != start.name {
				Rooms[name] = &Room{name: name}
			}
		} else {
			path := strings.Split(line, "-")
			if Rooms[path[0]].name == path[0] {
				Rooms[path[0]].links = append(Rooms[path[0]].links, Rooms[path[1]])
			}
		}

	}

	fmt.Println("ignore", end.name)
	cur := start
	for cur.links != nil {
		if cur == start {
			fmt.Println("start")
		}
		fmt.Print(cur.name, "->")
		if cur == end {
			fmt.Print("end")
			break
		}
		if cur.links[0].links == nil {
			fmt.Print(cur.links[0].name)
		}
		cur = cur.links[0]

	}
}
