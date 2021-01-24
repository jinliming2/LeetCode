package main

import "testing"

type inputData struct {
	action   string
	argument int
}

func TestFindMedianFromDataStream(t *testing.T) {
	for input, output := range map[*[]inputData][]float64{
		{
			{"addNum", 2},
			{"addNum", 3},
			{"addNum", 4},
			{"findMedian", 0},
		}: {3},
		{
			{"addNum", 2},
			{"addNum", 3},
			{"findMedian", 0},
		}: {2.5},
		{
			{"addNum", 1},
			{"addNum", 2},
			{"findMedian", 0},
			{"addNum", 3},
			{"findMedian", 0},
		}: {1.5, 2},
	} {
		obj := Constructor()

		index := 0
		for _, i := range *input {
			switch i.action {
			case "addNum":
				obj.AddNum(i.argument)
			case "findMedian":
				value := obj.FindMedian()
				if value != output[index] {
					t.Errorf("Expected LRUCache(%v)[%d] == %v, But got %v", *input, index, output[index], value)
					break
				}
				index++
			}
		}
	}
}
