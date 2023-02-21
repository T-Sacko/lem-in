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
	fmt.Println(antnum)

	rooms := make(map[string]*Room)

	var start, end *Room
	var name string

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
		if strings.HasPrefix(line, "#") || i == 0 || len(line) == 0 || line[0] == 'L' {
			continue
		}

		if len(strings.Fields(line)) == 3 {
			name = strings.Fields(line)[0]
			rooms[name] = &Room{name: name}
		} else {
			path := strings.Split(line, "-")

			if rooms[path[0]].name == path[0] {
				rooms[path[0]].links = append(rooms[path[0]].links, &Room{name: path[1]})
			}
		}

	}
	for start!=nil{
		
	}

	// fmt.Println(rooms[name].links)
	fmt.Println(start.links)
	fmt.Println(end)
	fmt.Println(rooms["0"].links[0] == nil)
}
