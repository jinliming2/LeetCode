package main

import "testing"

type inputData struct {
	action   string
	argument []int
}

func TestRandomizedSet(t *testing.T) {
	for input, output := range map[*[]inputData][]interface{}{
		{
			{"RandomizedSet", []int{}},
			{"insert", []int{1}},
			{"remove", []int{2}},
			{"insert", []int{2}},
			{"getRandom", []int{}},
			{"remove", []int{1}},
			{"insert", []int{2}},
			{"getRandom", []int{}},
		}: {nil, true, false, true, []int{1, 2}, true, false, []int{2}},
		{
			{"RandomizedSet", []int{}},
			{"insert", []int{0}},
			{"remove", []int{0}},
			{"insert", []int{-1}},
			{"remove", []int{0}},
			{"getRandom", []int{}},
			{"getRandom", []int{}},
			{"getRandom", []int{}},
			{"getRandom", []int{}},
			{"getRandom", []int{}},
			{"getRandom", []int{}},
			{"getRandom", []int{}},
			{"getRandom", []int{}},
			{"getRandom", []int{}},
			{"getRandom", []int{}},
		}: {nil, true, true, true, false, []int{-1}, []int{-1}, []int{-1}, []int{-1}, []int{-1}, []int{-1}, []int{-1}, []int{-1}, []int{-1}, []int{-1}},
	} {
		var obj RandomizedSet

	actionLoop:
		for index, i := range *input {
			var value interface{}
			switch i.action {
			case "RandomizedSet":
				obj = Constructor()
			case "insert":
				value = obj.Insert(i.argument[0])
			case "remove":
				value = obj.Remove(i.argument[0])
			case "getRandom":
				value = obj.GetRandom()
			}

			if output[index] == nil && value == nil {
				continue
			}
			if (output[index] == nil && value != nil) || (output[index] != nil && value == nil) {
				t.Errorf("Expected RandomizedSet(%v)[%d] == %v, But got %v", *input, index, output[index], value)
				break
			}
			if o, ok := output[index].([]int); ok {
				if v, ok := value.(int); ok {
					found := false
					for _, oItem := range o {
						if v == oItem {
							found = true
							break
						}
					}
					if found {
						continue actionLoop
					}
				}
				t.Errorf("Expected RandomizedSet(%v)[%d] in %v, But got %v", *input, index, o, value)
				break
			}
			if o, ok := output[index].(bool); ok {
				if v, ok := value.(bool); !ok || o != v {
					t.Errorf("Expected RandomizedSet(%v)[%d] == %t, But got %v", *input, index, o, value)
					break
				}
			}
		}
	}
}
