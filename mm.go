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

			Rooms[name] = end
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
				if Rooms[path[1]].name == end.name {
					end = Rooms[path[1]]
				}
			}
		}

	}

	cur := start
	var Possible bool
	for cur != nil {

		fmt.Print(cur.name, "->")
		if cur == end {
			Possible = true
			fmt.Print("end")
			break
		}

		cur = cur.links[0]

	}
	fmt.Println()
	fmt.Println(Possible)

	allPaths := DFSAll(start, end)
fmt.Println("All possible paths:")
for _, path := range allPaths {
    for i, room := range path {
        fmt.Print(room.name)
        if i < len(path)-1 {
            fmt.Print("->")
        }
    }
    fmt.Println()
}
	
}
func DFSAll(start, end *Room) [][]*Room {
    visited := make(map[*Room]bool)
    return dfsHelperAll(start, end, visited)
}

func dfsHelperAll(curr, end *Room, visited map[*Room]bool) [][]*Room {
    if curr == end {
        return [][]*Room{{curr}}
    }

    visited[curr] = true
    var allPaths [][]*Room

    for _, neighbor := range curr.links {
        if !visited[neighbor] {
            neighborPaths := dfsHelperAll(neighbor, end, visited)
            for _, path := range neighborPaths {
                allPaths = append(allPaths, append([]*Room{curr}, path...))
            }
        }
    }

    visited[curr] = false
    return allPaths
}