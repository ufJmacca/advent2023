package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type Step struct {
	heatloss           int
	row                int
	col                int
	next_row           int
	next_col           int
	steps_in_direction int
	index              int
}

type PreviousSteps struct {
	row                int
	col                int
	next_row           int
	next_col           int
	steps_in_direction int
}

type Set map[PreviousSteps]struct{}

func (s Set) Add(item PreviousSteps) {
	s[item] = struct{}{}
}

func (s Set) Contains(item PreviousSteps) bool {
	_, exists := s[item]
	return exists
}

type PriorityQueue []*Step

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heatloss < pq[j].heatloss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Step)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(step *Step, value string, priority int) {
	step.heatloss = priority
	heap.Fix(pq, step.index)
}

func Puzzle1(input string) int {
	lines := strings.Split(input, "\n")

	var grid [][]int

	for _, line := range lines {
		if len(line) > 0 {
			elements := strings.Split(line, "")
			var int_slice []int
			for _, element := range elements {
				int_val, _ := strconv.Atoi(element)
				int_slice = append(int_slice, int_val)
			}
			grid = append(grid, int_slice)
		}
	}

	grid_rows := len(grid)
	grid_cols := len(grid[0])

	pq := make(PriorityQueue, 1)
	previous_steps := make(Set)

	pq[0] = &Step{
		heatloss:           0,
		row:                0,
		col:                0,
		next_row:           0,
		next_col:           0,
		steps_in_direction: 0,
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		step := heap.Pop(&pq).(*Step)
		fmt.Printf("hl: %d, cell: (%d, %d), next direction: (%d, %d), step: %d\n", step.heatloss, step.row, step.col, step.next_row, step.next_col, step.steps_in_direction)

		if step.row == grid_rows-1 && step.col == grid_cols-1 {
			return step.heatloss
		}

		current_step := PreviousSteps{
			row:                step.row,
			col:                step.col,
			next_row:           step.next_row,
			next_col:           step.next_col,
			steps_in_direction: step.steps_in_direction,
		}
		if previous_steps.Contains(current_step) {
			fmt.Println("skip already seen value")
			continue
		}
		previous_steps.Add(current_step)

		if step.steps_in_direction < 3 && !(step.next_row == 0 && step.next_col == 0) {
			possible_row := step.row + step.next_row
			possible_col := step.col + step.next_col
			if possible_row >= 0 && possible_row < grid_rows && possible_col >= 0 && possible_col < grid_cols {
				fmt.Println("Continue in direction of Traval")
				fmt.Printf("PUSHED - hl: %d, cell: (%d, %d), next direction: (%d, %d), step: %d\n", step.heatloss+grid[possible_row][possible_col], possible_row, possible_col, step.next_row, step.next_col, step.steps_in_direction+1)
				heap.Push(&pq, &Step{
					heatloss:           step.heatloss + grid[possible_row][possible_col],
					row:                possible_row,
					col:                possible_col,
					next_row:           step.next_row,
					next_col:           step.next_col,
					steps_in_direction: step.steps_in_direction + 1,
				})
			}
		}

		possible_next_directions := [][]int{
			{0, 1},
			{0, -1},
			{1, 0},
			{-1, 0},
		}

		for _, possible_next := range possible_next_directions {
			if !(possible_next[0] == step.next_row && possible_next[1] == step.next_col) && !(possible_next[0] == step.next_row*-1 && possible_next[1] == step.next_col*-1) {
				possible_row := step.row + possible_next[0]
				possible_col := step.col + possible_next[1]
				if possible_row >= 0 && possible_row < grid_rows && possible_col >= 0 && possible_col < grid_cols {
					fmt.Printf("PUSHED - hl: %d, cell: (%d, %d), next direction: (%d, %d), step: %d\n", step.heatloss+grid[possible_row][possible_col], possible_row, possible_col, possible_next[0], possible_next[1], 1)
					heap.Push(&pq, &Step{
						heatloss:           step.heatloss + grid[possible_row][possible_col],
						row:                possible_row,
						col:                possible_col,
						next_row:           possible_next[0],
						next_col:           possible_next[1],
						steps_in_direction: 1,
					})
				}
			}
		}
	}
	return 0
}
