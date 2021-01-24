package main

import "testing"

type inputData struct {
	action   string
	argument []int
}

func TestLFUCache(t *testing.T) {
	for input, output := range map[*[]inputData][][]int{
		{
			{"LFUCache", []int{2}},
			{"put", []int{1, 1}},
			{"put", []int{2, 2}},
			{"get", []int{1}},
			{"put", []int{3, 3}},
			{"get", []int{2}},
			{"get", []int{3}},
			{"put", []int{4, 4}},
			{"get", []int{1}},
			{"get", []int{3}},
			{"get", []int{4}},
		}: {nil, nil, nil, {1}, nil, {-1}, {3}, nil, {-1}, {3}, {4}},
		{
			{"LFUCache", []int{0}},
			{"put", []int{0, 0}},
			{"get", []int{0}},
		}: {nil, nil, {-1}},
	} {
		var obj LFUCache

		for index, i := range *input {
			var value []int
			switch i.action {
			case "put":
				obj.Put(i.argument[0], i.argument[1])
			case "get":
				value = []int{obj.Get(i.argument[0])}
			case "LFUCache":
				obj = Constructor(i.argument[0])
			}

			if output[index] == nil && value == nil {
				continue
			}
			if (output[index] == nil && value != nil) ||
				(output[index] != nil && value == nil) ||
				(output[index][0] != value[0]) {
				t.Errorf("Expected LFUCache(%v)[%d] == %v, But got %v", *input, index, output[index], value)
				break
			}
		}
	}
}
