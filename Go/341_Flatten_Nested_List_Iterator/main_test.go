package main

import "testing"

func TestNestedIterator(t *testing.T) {
	for input, output := range map[*[]*NestedInteger][]int{
		{
			{
				list: []*NestedInteger{
					{val: 1},
					{val: 1},
				},
			},
			{val: 2},
			{
				list: []*NestedInteger{
					{val: 1},
					{val: 1},
				},
			},
		}: {1, 1, 2, 1, 1},
		{
			{val: 1},
			{
				list: []*NestedInteger{
					{val: 4},
					{
						list: []*NestedInteger{
							{val: 6},
						},
					},
				},
			},
		}: {1, 4, 6},
		{
			{
				list: []*NestedInteger{},
			},
		}: {},
	} {
		ni := Constructor(*input)
		for _, item := range output {
			if !ni.HasNext() {
				t.Errorf("Expected ni.HasNext() == %t , But got %t", true, false)
				break
			}
			if value := ni.Next(); value != item {
				t.Errorf("Expected ni.Next() == %d, But got %d", item, value)
				break
			}
		}
		if ni.HasNext() {
			t.Errorf("Expected ni.HasNext() == %t , But got %t", false, true)
		}
	}
}
