package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	graph := homeworkGraph("tc1.txt")
	grev := graph.reverse()

	expected := map[int][]int{9: []int{6}, 6: []int{3, 8}, 3: []int{9}, 1: []int{7}, 4: []int{1}, 8: []int{2}, 2: []int{5}, 5: []int{8}, 7: []int{4, 9}}

	actual := grev.List
	sorted := make(map[int][]int)

	for key, val := range actual {
		sort.Ints(val)
		sorted[key] = val
	}

	assert.Equal(t, expected, sorted)
}

func TestSampleCase1(t *testing.T) {
	graph := homeworkGraph("tc1.txt")

	expected := []int{3, 3, 3, 0, 0}
	actual := graph.kosaraju()

	assert.Equal(t, expected, actual)
}

func TestSampleCase2(t *testing.T) {
	graph := homeworkGraph("tc2.txt")

	expected := []int{3, 3, 2, 0, 0}
	actual := graph.kosaraju()

	assert.Equal(t, expected, actual)
}

func TestSampleCase3(t *testing.T) {
	graph := homeworkGraph("tc3.txt")

	expected := []int{3, 3, 1, 1, 0}
	actual := graph.kosaraju()

	assert.Equal(t, expected, actual)
}

func TestSampleCase4(t *testing.T) {
	graph := homeworkGraph("tc4.txt")

	expected := []int{7, 1, 0, 0, 0}
	actual := graph.kosaraju()

	assert.Equal(t, expected, actual)
}

func TestAssignment(t *testing.T) {
	graph := homeworkGraph("scc.txt")

	expected := []int{434821, 968, 459, 313, 211}
	actual := graph.kosaraju()

	assert.Equal(t, expected, actual)
}

func homeworkGraph(filename string) *Graph {
	input := integersFromFile(filename)

	return &Graph{List: input}
}

func integersFromFile(s string) map[int][]int {
	inputBytes, err := ioutil.ReadFile(s)

	if err != nil {
		panic(err)
	}

	inputString := string(inputBytes[:])
	rawRows := strings.Split(inputString, "\n")
	data := make(map[int][]int)

	for _, i := range rawRows {
		values := strings.Split(i, " ")

		cols := []int{}

		for _, v := range values {
			j, err := strconv.Atoi(strings.TrimSpace(v))

			if err == nil {
				cols = append(cols, j)
			}
		}

		if len(cols) > 1 {
			vertex := cols[0]
			arc := cols[1]

			data[vertex] = append(data[vertex], arc)
		}
	}

	return data
}
