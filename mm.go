package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
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
			if ting[i-1] == "##start"  || ting[i-1] == "##end"{
				continue
			}
			Rooms[name] = &Room{name: name}
		} else {
			path := strings.Split(line, "-")
			if Rooms[path[1]].name == end.name {
				end = Rooms[path[1]]
			} else if Rooms[path[0]].name == end.name {
				end = Rooms[path[0]]
			}
			if Rooms[path[0]] != nil && Rooms[path[1]] != nil {
				Rooms[path[0]].links = append(Rooms[path[0]].links, Rooms[path[1]])
				Rooms[path[1]].links = append(Rooms[path[1]].links, Rooms[path[0]])
			}

		}
	}

	allPaths, err := DFSAll(start, end)
	if err != nil {
		log.Fatal(err)
	}

	sortByLength(allPaths)

	ants := make([]int, antnum)
	for i := range ants {
		ants[i] = 1 + i
	}

	// assign := sing(allPaths, ants)
	// fmt.Println(Rooms["2"].links[1].name)
	// newPaths := refine(allPaths, assign)

	// for i := 0; i < len(newPaths); i++ {
	// 	for j := 0; j < len(assign); j++ {
	// 		fmt.Printf("L%d-%d", ants[i], assign[i][j])
	// 	}
	// }
	allPathsnew := CombinePaths(allPaths)
	fmt.Println()
	fmt.Println(len(allPathsnew))
	//fmt.Println(allPathsnew[5][1][3])
	// fmt.Println(assign[3])
}

func CombinePaths(AllPaths [][]*Room) [][][]*Room {
	// THE UGLIEST FUNCTION I HAVE EVER WRITTEN

	Result := make([][][]*Room, 0)
	CombPaths := make([][]*Room, 0)
	var counter int
	var Breaker bool

	for _, P1 := range AllPaths {
		// compare 1 Path with all other Paths
		CombPaths = append(CombPaths, P1)
		for _, P2 := range AllPaths {
			// compare the Paths node by node.
			for i, P := range CombPaths {
				if !Breaker {
					for _, v := range P[1 : len(P)-1] {
						if inArray(P2[1:len(P2)-1], v) {
							Breaker = true
							break
						}
					}
				}
				if i == len(CombPaths)-1 && !Breaker {
					CombPaths = append(CombPaths, P2)
				}
			}
			Breaker = false
		}
		if counter <= len(CombPaths) {
			Result = append(Result, CombPaths)
			counter = len(CombPaths)
		}

		CombPaths = nil
	}
	return Result
}

func inArray(s []*Room, vp *Room) (result bool) {
	for _, v := range s {
		if reflect.DeepEqual(v, vp) {
			result = true
			return
		}
	}
	return
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

// func refine(all [][]*Room, maps map[int][]int) [][]*Room {
// 	var paths [][]*Room
// 	for i := 0; i < len(maps); i++ {
// 		paths = append(paths, all[i])
// 	}
// 	return paths
// }

// func sing(allPaths [][]*Room, ants []int) map[int][]int {
// 	assign := make(map[int][]int)
// 	done := make(map[int]bool)
// 	for i := 0; i <= len(allPaths); i++ {
// 		for j := 0; j < len(ants); j++ {
// 			if i < len(allPaths)-1 {
// 				if len(assign[i])+len(allPaths[i]) <= len(allPaths[i+1]) && !done[ants[j]] {
// 					assign[i] = append(assign[i], ants[j])
// 					done[ants[j]] = true
// 				}
// 			}
// 			if i == len(allPaths)-1 {
// 				assign[i] = append(assign[i], ants[j])
// 			}

// 		}
// 	}
// 	return assign
// }

// assign[7] = append(assign[7], 4)
// fmt.Println(len(assign[7]))
// done := make(map[int]bool)
// for i := 1; i <= len(allPaths); i++ {
// 	for j := 0; j < len(ants); j++ {
// 		if len(assign[i]) <= len(assign[i+1]) &&!done[ants[j]] {
// 			assign[i] = append(assign[i], ants[j])
// 			done[ants[j]]=true
// 		}
// 	}
// }
