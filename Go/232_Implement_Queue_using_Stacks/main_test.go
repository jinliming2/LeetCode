package main

import "testing"

type inputData struct {
	action   string
	argument []int
}

func TestMyQueue(t *testing.T) {
	for input, output := range map[*[]inputData][]interface{}{
		{
			{"MyQueue", nil},
			{"push", []int{1}},
			{"push", []int{2}},
			{"peek", nil},
			{"pop", nil},
			{"empty", nil},
		}: {nil, nil, nil, 1, 1, false},
	} {
		var obj MyQueue

		for index, i := range *input {
			var value interface{}
			switch i.action {
			case "MyQueue":
				obj = Constructor()
			case "push":
				obj.Push(i.argument[0])
			case "peek":
				value = obj.Peek()
			case "pop":
				value = obj.Pop()
			case "empty":
				value = obj.Empty()
			}

			if output[index] == nil && value == nil {
				continue
			}
			if (output[index] == nil && value != nil) || (output[index] != nil && value == nil) {
				t.Errorf("Expected MyQueue(%v)[%d] == %v, But got %v", *input, index, output[index], value)
				break
			}
			if o, ok := output[index].(int); ok {
				if v, ok := value.(int); !ok || o != v {
					t.Errorf("Expected MyQueue(%v)[%d] == %d, But got %v", *input, index, o, value)
					break
				}
			}
			if o, ok := output[index].(bool); ok {
				if v, ok := value.(bool); !ok || o != v {
					t.Errorf("Expected MyQueue(%v)[%d] == %t, But got %v", *input, index, o, value)
					break
				}
			}
		}
	}
}
