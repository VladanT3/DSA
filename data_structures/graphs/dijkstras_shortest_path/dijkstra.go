package dijkstra

import "math"

type GraphEdge struct {
	To     int
	Weight uint8
}

type WeightedAdjecencyList [][]GraphEdge

func DijkstrasShortestPath(source, sink int, arr WeightedAdjecencyList) []int {
	var seen []bool = make([]bool, len(arr), len(arr))
	for i := range seen {
		seen[i] = false
	}
	var prev []int = make([]int, len(arr), len(arr))
	for i := range prev {
		prev[i] = -1
	}
	var dists []uint8 = make([]uint8, len(arr), len(arr))
	for i := range dists {
		dists[i] = uint8(math.Pow(2, 8) - 1)
	}
	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := arr[curr]
		for i := range adjs {
			edge := adjs[i]
			if seen[edge.To] {
				continue
			}

			dist := dists[curr] + edge.Weight
			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	out := []int{}
	path := []int{}
	curr := sink
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

func hasUnvisited(seen []bool, dists []uint8) bool {
	has := false

	for i := range len(seen) {
		if seen[i] == false && dists[i] < 255 {
			has = true
		}
	}

	return has
}

func getLowestUnvisited(seen []bool, dists []uint8) int {
	idx := -1
	var lowest_distance uint8 = 255

	for i := range dists {
		if dists[i] < lowest_distance && seen[i] == false {
			lowest_distance = dists[i]
			idx = i
		}
	}

	return idx
}
