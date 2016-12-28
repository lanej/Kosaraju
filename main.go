package main

import (
	"fmt"
	"sort"
)

// Graph represents a collection of nodes
type Graph struct {
	List map[int][]int
}

func (g *Graph) vertices() []int {
	var keys []int
	for key := range g.List {
		keys = append(keys, key)
	}
	return keys
}

func (g *Graph) finishingTimes() []int {
	stack := NewStack()
	explored := make(map[int]bool)
	finishingTime := make(map[int]int)
	timing := 0

	vertices := g.vertices()

	for _, vertex := range vertices {
		stack.Push(vertex)
	}

	for !stack.Empty() {
		vertex := stack.Top().(int)
		_, hasExplored := explored[vertex]

		if hasExplored {
			timing++

			// Set timing if not already set
			_, alreadyFinished := finishingTime[vertex]
			if !alreadyFinished {
				finishingTime[vertex] = timing
			}

			// Lose the top element
			stack.Pop()
		} else {
			unexploredArcs := []int{}
			explored[vertex] = true

			// Find unexplored arcs from this vertex
			for _, v := range g.List[vertex] {
				_, alreadyExplored := explored[v]
				if !alreadyExplored {
					unexploredArcs = append(unexploredArcs, v)
				}
			}

			// If no unexplored arcs, this path ends
			if len(unexploredArcs) == 0 {
				timing++
				finishingTime[vertex] = timing
				stack.Pop()
			} else {
				// Add unexplored arcs to the stack on top of the current top
				for _, v := range unexploredArcs {
					stack.Push(v)
				}
			}
		}
	}

	sortedTimes := keysSortedByValue(finishingTime)

	return sortedTimes
}

func keysSortedByValue(m map[int]int) []int {
	var a []int
	reversed := map[int][]int{}
	var sortedTimes []int

	for k, v := range m {
		reversed[v] = append(reversed[v], k)
	}
	for k := range reversed {
		a = append(a, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	for _, k := range a {
		for _, s := range reversed[k] {
			sortedTimes = append(sortedTimes, s)
		}
	}

	return sortedTimes
}

func (g *Graph) kosaraju() []int {
	grev := g.reverse()
	l := grev.finishingTimes()
	leaders := make(map[int]int)
	explored := make(map[int]bool)

	for _, leader := range l {
		leaders[leader] = 0
		_, leaderExplored := explored[leader]

		if !leaderExplored {
			stack := NewStack()
			stack.Push(leader)

			for !stack.Empty() {
				curr := stack.Pop().(int)

				_, vertexExplored := explored[curr]

				if vertexExplored {
					continue
				}

				leaders[leader]++
				explored[curr] = true

				for _, v := range g.List[curr] {
					_, arcExplored := explored[v]

					if !arcExplored {
						stack.Push(v)
					}
				}
			}
		}
	}

	sizes := []int{}

	for _, v := range leaders {
		sizes = append(sizes, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	return sizes[0:5]
}

func (g *Graph) reverse() *Graph {
	grev := make(map[int][]int)
	for key, val := range g.List {
		for _, v := range val {
			grev[v] = append(grev[v], key)
		}
	}
	return &Graph{List: grev}
}

func (g *Graph) clone() *Graph {
	newList := make(map[int][]int)
	for key, val := range g.List {
		newList[key] = val
	}
	return &Graph{List: newList}
}

func (g *Graph) print(name string) {
	fmt.Printf("Graph: %s\n", name)
	for key := range g.List {
		fmt.Printf("%v - %v\n", key, g.List[key])
	}
}

func nChooseK(n, k uint64) uint64 {
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func factorial(x uint64) uint64 {
	var result uint64
	if x < 1 {
		result = 1
	} else {
		result = x * factorial(x-1)
	}
	return result
}
