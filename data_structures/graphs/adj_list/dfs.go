package adjlist

type GraphEdge struct {
	To     int
	Weight int
}

type WeightedAdjacencyList [][]GraphEdge

func AdjListDFS(graph WeightedAdjacencyList, source int, needle int) []int {
	var seen []bool = make([]bool, len(graph), len(graph))
	for i := range seen {
		seen[i] = false
	}
	var path []int = []int{}

	walk(graph, source, needle, &seen, &path)

	return path
}

func walk(graph WeightedAdjacencyList, curr int, needle int, seen *[]bool, path *[]int) bool {
	var temp_seen []bool = *seen
	var temp_path []int = *path

	if temp_seen[curr] {
		return false
	}

	temp_seen[curr] = true

	temp_path = append(temp_path, curr)
	if curr == needle {
		*path = temp_path
		return true
	}

	list := graph[curr]
	for i := 0; i < len(list); i++ {
		edge := list[i]
		*path = temp_path
		*seen = temp_seen
		if walk(graph, edge.To, needle, seen, path) {
			return true
		}
	}

	if len(temp_path) == 0 || len(temp_path) == 1 {
		temp_path = []int{}
	} else {
		temp_path = temp_path[:len(temp_path)]
	}
	*path = temp_path

	return false
}
