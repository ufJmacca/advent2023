package main

type queue_item struct {
	cell            [2]int
	steps_remaining int
}

type queue struct {
	items []queue_item
}

var directions = [][]int{
	{1, 0},  // Down
	{-1, 0}, // Up
	{0, 1},  // Right
	{0, -1}, // Left
}

func (q *queue) Enqueue(item queue_item) {
	q.items = append(q.items, item)
}

func (q *queue) Dequeue() queue_item {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func Function(grid [][]string, start [2]int, steps int) int {
	rows := len(grid)
	cols := len(grid[0])

	answer := make(map[[2]int]int)
	seen := make(map[[2]int]int)

	p_queue := queue{}
	p_queue.Enqueue(queue_item{
		cell:            start,
		steps_remaining: steps,
	})

	for len(p_queue.items) > 0 {
		element := p_queue.Dequeue()

		if element.steps_remaining%2 == 0 {
			answer[[2]int(element.cell)] = 1
		}
		if element.steps_remaining == 0 {
			continue
		}

		for _, direction := range directions {
			next_x, next_y := element.cell[0]+direction[0], element.cell[1]+direction[1]
			_, already_seen := seen[[2]int{next_x, next_y}]

			if next_x >= 0 && next_x < rows && next_y >= 0 && next_y < cols && grid[next_x][next_y] != "#" && !already_seen {
				seen[[2]int{next_x, next_y}] = 1
				p_queue.Enqueue(queue_item{
					cell:            [2]int{next_x, next_y},
					steps_remaining: element.steps_remaining - 1,
				})
			}
		}

	}

	return len(answer)
}

func Puzzle1(input string, steps int) int {

	return 0
}
