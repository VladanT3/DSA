package adjmatrix

type WeightedAdjacencyMatrix [][]int

func AdjMatrixBFS(graph WeightedAdjacencyMatrix, source int, needle int) []int {
	var seen []bool = make([]bool, len(graph), len(graph))
	for i := range seen {
		seen[i] = false
	}

	var prev []int = make([]int, len(graph), len(graph))
	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true
	queue := []int{source}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := range adjs {
			if adjs[i] == 0 {
				continue
			}
			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			queue = append(queue, i)
		}
	}

	curr := needle
	var out []int
	var path []int

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) > 0 {
		out = append(out, source)
		for i := len(out) - 1; i >= 0; i-- {
			path = append(path, out[i])
		}
	}

	return path
}
