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

	Rooms := make(map[string]*Room)

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
			Rooms[name] = &Room{name: name}
		} else {
			path := strings.Split(line, "-")

			if Rooms[path[0]].name == path[0] {
				Rooms[path[0]].links = append(Rooms[path[0]].links, &Room{name: path[1]})
			}
		}

	}
	start.links = Rooms[start.name].links
	fmt.Println(Rooms[start.name].links[0])
	fmt.Println(start.links[0])
	fmt.Println(end)
	// fmt.Println(rooms["0"].links[0] == nil)
}
