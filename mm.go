package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
			if Rooms[path[0]] != nil && Rooms[path[1]] != nil {
				Rooms[path[0]].links = append(Rooms[path[0]].links, Rooms[path[1]])
				Rooms[path[1]].links = append(Rooms[path[1]].links, Rooms[path[0]])
			}
			if Rooms[path[1]].name == end.name {
				end = Rooms[path[1]]
			} else if Rooms[path[0]].name == end.name {
				end = Rooms[path[0]]
			}

		}
	}

	allPaths, err := DFSAll(start, end)
	if err != nil {
		log.Fatal(err)
	}

	sortByLength(allPaths)
	fmt.Println(allPaths[0][1].name)

	// for i, path := range allPaths{
	// 	if l
	// }

	ants := make([]int, antnum)
	for i := range ants {
		ants[i] = 1 + i
	}
	assign := make(map[int][]int)
	assign[7] = append(assign[7], 4)
	fmt.Println(len(assign[7]))
	done := make(map[int]bool)
	for i := 1; i <= len(allPaths); i++ {
		for j := 0; j < len(ants); j++ {
			if len(assign[i]) <= len(assign[i+1]) &&!done[ants[j]] {
				assign[i] = append(assign[i], ants[j])
				done[ants[j]]=true
			}
		}
	}
	fmt.Println(assign[5])
}

func sortByLength(paths [][]*Room) {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
}

func DFSAll(start, end *Room) ([][]*Room, error) {
	visited := make(map[*Room]bool)
	paths := dfsHelperAll(start, end, visited)
	if len(paths) == 0 {
		return nil, fmt.Errorf("no path found between %s and %s", start.name, end.name)
	}
	return paths, nil
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
