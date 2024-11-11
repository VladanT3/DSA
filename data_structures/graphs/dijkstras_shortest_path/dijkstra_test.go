package dijkstra

import (
	"reflect"
	"testing"
)

func TestDijkstrasShortestPath(t *testing.T) {
	test_list := [][]GraphEdge{
		[]GraphEdge{
			GraphEdge{
				To:     1,
				Weight: 3,
			},
			GraphEdge{
				To:     2,
				Weight: 1,
			},
		},
		[]GraphEdge{
			GraphEdge{
				To:     0,
				Weight: 3,
			},
			GraphEdge{
				To:     2,
				Weight: 4,
			},
			GraphEdge{
				To:     4,
				Weight: 1,
			},
		},
		[]GraphEdge{
			GraphEdge{
				To:     1,
				Weight: 4,
			},
			GraphEdge{
				To:     3,
				Weight: 7,
			},
			GraphEdge{
				To:     0,
				Weight: 1,
			},
		},
		[]GraphEdge{
			GraphEdge{
				To:     2,
				Weight: 7,
			},
			GraphEdge{
				To:     4,
				Weight: 5,
			},
			GraphEdge{
				To:     6,
				Weight: 1,
			},
		},
		[]GraphEdge{
			GraphEdge{
				To:     1,
				Weight: 1,
			},
			GraphEdge{
				To:     3,
				Weight: 5,
			},
			GraphEdge{
				To:     5,
				Weight: 2,
			},
		},
		[]GraphEdge{
			GraphEdge{
				To:     6,
				Weight: 1,
			},
			GraphEdge{
				To:     4,
				Weight: 2,
			},
			GraphEdge{
				To:     2,
				Weight: 18,
			},
		},
		[]GraphEdge{
			GraphEdge{
				To:     1,
				Weight: 1,
			},
			GraphEdge{
				To:     5,
				Weight: 1,
			},
		},
	}

	res := DijkstrasShortestPath(0, 6, test_list)
	if !reflect.DeepEqual(res, []int{0, 1, 4, 5, 6}) {
		t.Errorf("\nexpected: %v\ngot: %v", []int{0, 1, 4, 5, 6}, res)
	}
}
