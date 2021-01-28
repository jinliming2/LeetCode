package main

import "testing"

type inputData struct {
	action   string
	argument []int
}

func TestMyStack(t *testing.T) {
	for input, output := range map[*[]inputData][]interface{}{
		{
			{"MyStack", nil},
			{"push", []int{1}},
			{"push", []int{2}},
			{"top", nil},
			{"pop", nil},
			{"empty", nil},
		}: {nil, nil, nil, 2, 2, false},
	} {
		var obj MyStack

		for index, i := range *input {
			var value interface{}
			switch i.action {
			case "MyStack":
				obj = Constructor()
			case "push":
				obj.Push(i.argument[0])
			case "top":
				value = obj.Top()
			case "pop":
				value = obj.Pop()
			case "empty":
				value = obj.Empty()
			}

			if output[index] == nil && value == nil {
				continue
			}
			if (output[index] == nil && value != nil) || (output[index] != nil && value == nil) {
				t.Errorf("Expected MyStack(%v)[%d] == %v, But got %v", *input, index, output[index], value)
				break
			}
			if o, ok := output[index].(int); ok {
				if v, ok := value.(int); !ok || o != v {
					t.Errorf("Expected MyStack(%v)[%d] == %d, But got %v", *input, index, o, value)
					break
				}
			}
			if o, ok := output[index].(bool); ok {
				if v, ok := value.(bool); !ok || o != v {
					t.Errorf("Expected MyStack(%v)[%d] == %t, But got %v", *input, index, o, value)
					break
				}
			}
		}
	}
}
