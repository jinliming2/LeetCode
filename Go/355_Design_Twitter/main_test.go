package main

import "testing"

type inputData struct {
	action    string
	arguments []int
}

func TestTwitter(t *testing.T) {
allLoop:
	for input, output := range map[*[]inputData][][]int{
		{
			{"postTweet", []int{1, 5}},
			{"getNewsFeed", []int{1}},
			{"follow", []int{1, 2}},
			{"postTweet", []int{2, 6}},
			{"getNewsFeed", []int{1}},
			{"unfollow", []int{1, 2}},
			{"getNewsFeed", []int{1}},
		}: {{5}, {6, 5}, {5}},
		{
			{"postTweet", []int{1, 5}},
			{"unfollow", []int{1, 1}},
			{"getNewsFeed", []int{1}},
		}: {{5}},
	} {
		obj := Constructor()

		index := 0
		for _, i := range *input {
			switch i.action {
			case "postTweet":
				obj.PostTweet(i.arguments[0], i.arguments[1])
			case "getNewsFeed":
				value := obj.GetNewsFeed(i.arguments[0])
				if len(value) != len(output[index]) {
					t.Errorf("Expected len(Twitter(%v)[%d]) == %d , But got %d", input, index, len(output[index]), len(value))
					break
				}
				for x, item := range output[index] {
					if item != value[x] {
						t.Errorf("Expected Twitter(%v)[%d][%d] == %v, But got %v", *input, index, x, output[index][x], value)
						continue allLoop
					}
				}
				index++
			case "follow":
				obj.Follow(i.arguments[0], i.arguments[1])
			case "unfollow":
				obj.Unfollow(i.arguments[0], i.arguments[1])
			}
		}
	}
}
