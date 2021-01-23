package main

import "testing"

type inputData struct {
	action   string
	argument []int
}

func TestLRUCache(t *testing.T) {
	for input, output := range map[*[]inputData][][]int{
		{
			{"LRUCache", []int{2}},
			{"put", []int{1, 1}},
			{"put", []int{2, 2}},
			{"get", []int{1}},
			{"put", []int{3, 3}},
			{"get", []int{2}},
			{"put", []int{4, 4}},
			{"get", []int{1}},
			{"get", []int{3}},
			{"get", []int{4}},
		}: {nil, nil, nil, {1}, nil, {-1}, nil, {-1}, {3}, {4}},
		{
			{"LRUCache", []int{2}},
			{"put", []int{2, 1}},
			{"put", []int{2, 2}},
			{"get", []int{2}},
			{"put", []int{1, 1}},
			{"put", []int{4, 1}},
			{"get", []int{2}},
		}: {nil, nil, nil, {2}, nil, nil, {-1}},
		{
			{"LRUCache", []int{10}},
			{"put", []int{10, 13}},
			{"put", []int{3, 17}},
			{"put", []int{6, 11}},
			{"put", []int{10, 5}},
			{"put", []int{9, 10}},
			{"get", []int{13}},
			{"put", []int{2, 19}},
			{"get", []int{2}},
			{"get", []int{3}},
			{"put", []int{5, 25}},
			{"get", []int{8}},
			{"put", []int{9, 22}},
			{"put", []int{5, 5}},
			{"put", []int{1, 30}},
			{"get", []int{11}},
			{"put", []int{9, 12}},
			{"get", []int{7}},
			{"get", []int{5}},
			{"get", []int{8}},
			{"get", []int{9}},
			{"put", []int{4, 30}},
			{"put", []int{9, 3}},
			{"get", []int{9}},
			{"get", []int{10}},
			{"get", []int{10}},
			{"put", []int{6, 14}},
			{"put", []int{3, 1}},
			{"get", []int{3}},
			{"put", []int{10, 11}},
			{"get", []int{8}},
			{"put", []int{2, 14}},
			{"get", []int{1}},
			{"get", []int{5}},
			{"get", []int{4}},
			{"put", []int{11, 4}},
			{"put", []int{12, 24}},
			{"put", []int{5, 18}},
			{"get", []int{13}},
			{"put", []int{7, 23}},
			{"get", []int{8}},
			{"get", []int{12}},
			{"put", []int{3, 27}},
			{"put", []int{2, 12}},
			{"get", []int{5}},
			{"put", []int{2, 9}},
			{"put", []int{13, 4}},
			{"put", []int{8, 18}},
			{"put", []int{1, 7}},
			{"get", []int{6}},
			{"put", []int{9, 29}},
			{"put", []int{8, 21}},
			{"get", []int{5}},
			{"put", []int{6, 30}},
			{"put", []int{1, 12}},
			{"get", []int{10}},
			{"put", []int{4, 15}},
			{"put", []int{7, 22}},
			{"put", []int{11, 26}},
			{"put", []int{8, 17}},
			{"put", []int{9, 29}},
			{"get", []int{5}},
			{"put", []int{3, 4}},
			{"put", []int{11, 30}},
			{"get", []int{12}},
			{"put", []int{4, 29}},
			{"get", []int{3}},
			{"get", []int{9}},
			{"get", []int{6}},
			{"put", []int{3, 4}},
			{"get", []int{1}},
			{"get", []int{10}},
			{"put", []int{3, 29}},
			{"put", []int{10, 28}},
			{"put", []int{1, 20}},
			{"put", []int{11, 13}},
			{"get", []int{3}},
			{"put", []int{3, 12}},
			{"put", []int{3, 8}},
			{"put", []int{10, 9}},
			{"put", []int{3, 26}},
			{"get", []int{8}},
			{"get", []int{7}},
			{"get", []int{5}},
			{"put", []int{13, 17}},
			{"put", []int{2, 27}},
			{"put", []int{11, 15}},
			{"get", []int{12}},
			{"put", []int{9, 19}},
			{"put", []int{2, 15}},
			{"put", []int{3, 16}},
			{"get", []int{1}},
			{"put", []int{12, 17}},
			{"put", []int{9, 1}},
			{"put", []int{6, 19}},
			{"get", []int{4}},
			{"get", []int{5}},
			{"get", []int{5}},
			{"put", []int{8, 1}},
			{"put", []int{11, 7}},
			{"put", []int{5, 2}},
			{"put", []int{9, 28}},
			{"get", []int{1}},
			{"put", []int{2, 2}},
			{"put", []int{7, 4}},
			{"put", []int{4, 22}},
			{"put", []int{7, 24}},
			{"put", []int{9, 26}},
			{"put", []int{13, 28}},
			{"put", []int{11, 26}},
		}: {
			nil, nil, nil, nil, nil, nil, {-1}, nil, {19}, {17},
			nil, {-1}, nil, nil, nil, {-1}, nil, {-1}, {5}, {-1},
			{12}, nil, nil, {3}, {5}, {5}, nil, nil, {1}, nil,
			{-1}, nil, {30}, {5}, {30}, nil, nil, nil, {-1}, nil,
			{-1}, {24}, nil, nil, {18}, nil, nil, nil, nil, {-1},
			nil, nil, {18}, nil, nil, {-1}, nil, nil, nil, nil,
			nil, {18}, nil, nil, {-1}, nil, {4}, {29}, {30}, nil,
			{12}, {-1}, nil, nil, nil, nil, {29}, nil, nil, nil,
			nil, {17}, {22}, {18}, nil, nil, nil, {-1}, nil, nil,
			nil, {20}, nil, nil, nil, {-1}, {18}, {18}, nil, nil,
			nil, nil, {20}, nil, nil, nil, nil, nil, nil, nil,
		},
	} {
		var obj LRUCache

		for index, i := range *input {
			var value []int
			switch i.action {
			case "put":
				obj.Put(i.argument[0], i.argument[1])
			case "get":
				value = []int{obj.Get(i.argument[0])}
			case "LRUCache":
				obj = Constructor(i.argument[0])
			}

			if output[index] == nil && value == nil {
				continue
			}
			if (output[index] == nil && value != nil) ||
				(output[index] != nil && value == nil) ||
				(output[index][0] != value[0]) {
				t.Errorf("Expected LRUCache(%v)[%d] == %v, But got %v", *input, index, output[index], value)
				break
			}
		}
	}
}
