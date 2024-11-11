package adjlist

import (
	"reflect"
	"testing"
)

func TestAdjListDFS(t *testing.T) {
	test_list := [][]GraphEdge{
		[]GraphEdge{
			GraphEdge{
				To:     1,
				Weight: 5,
			},
			GraphEdge{
				To:     2,
				Weight: 5,
			},
		},
		[]GraphEdge{
			GraphEdge{
				To:     4,
				Weight: 5,
			},
		},
		[]GraphEdge{
			GraphEdge{
				To:     3,
				Weight: 5,
			},
		},
		[]GraphEdge{},
		[]GraphEdge{
			GraphEdge{
				To:     1,
				Weight: 5,
			},
			GraphEdge{
				To:     3,
				Weight: 5,
			},
			GraphEdge{
				To:     5,
				Weight: 5,
			},
		},
		[]GraphEdge{
			GraphEdge{
				To:     2,
				Weight: 5,
			},
			GraphEdge{
				To:     6,
				Weight: 5,
			},
		},
		[]GraphEdge{
			GraphEdge{
				To:     3,
				Weight: 5,
			},
		},
	}

	res := AdjListDFS(test_list, 0, 6)
	if !reflect.DeepEqual(res, []int{0, 1, 4, 5, 6}) {
		t.Errorf("\nexpected: %v\ngot: %v", []int{0, 1, 4, 5, 6}, res)
	}
	res = AdjListDFS(test_list, 6, 0)
	if !reflect.DeepEqual(res, []int{}) {
		t.Errorf("\nexpected: %v\ngot: %v", []int{}, res)
	}
}
