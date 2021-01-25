package main

// Runtime: 44 ms, faster than 100.00% of Go online submissions for Design Twitter.
// Memory Usage: 12.1 MB, less than 22.73% of Go online submissions for Design Twitter.
// https://leetcode.com/submissions/detail/447670527/

type tweet struct {
	tid, timestamp int
	next           *tweet
}

type user struct {
	uid      int
	followed map[*user]bool
	tweets   *tweet
}

func newUser(uid int) *user {
	u := &user{
		uid:      uid,
		followed: make(map[*user]bool),
	}
	u.follow(u)
	return u
}

func (u *user) post(timestamp, id int) {
	u.tweets = &tweet{tid: id, timestamp: timestamp, next: u.tweets}
}

func (u *user) follow(user *user) {
	u.followed[user] = true
}

func (u *user) unfollow(user *user) {
	if user.uid != u.uid {
		delete(u.followed, user)
	}
}

// Twitter ...
type Twitter struct {
	timestamp int
	users     map[int]*user
}

// Constructor : Initialize your data structure here.
func Constructor() Twitter {
	return Twitter{
		timestamp: 0,
		users:     make(map[int]*user),
	}
}

// PostTweet : Compose a new tweet.
func (twitter *Twitter) PostTweet(userID int, tweetID int) {
	user, ok := twitter.users[userID]
	if !ok {
		user = newUser(userID)
		twitter.users[userID] = user
	}
	user.post(twitter.timestamp, tweetID)
	twitter.timestamp++
}

// GetNewsFeed : Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent.
func (twitter *Twitter) GetNewsFeed(userID int) []int {
	user, ok := twitter.users[userID]
	if !ok {
		user = newUser(userID)
		twitter.users[userID] = user
		return []int{}
	}
	followedCount := len(user.followed)
	if followedCount == 0 {
		return []int{}
	}
	fq := newHeapFQ(followedCount)
	for followed := range user.followed {
		if followed.tweets != nil {
			fq.push(followed.tweets)
		}
	}
	res := make([]int, 0, 10)
	for i := 0; i < 10 && len(fq.data) > 0; i++ {
		t := fq.pop()
		res = append(res, t.tid)
		if t.next != nil {
			fq.push(t.next)
		}
	}
	return res
}

// Follow : Follower follows a followee. If the operation is invalid, it should be a no-op.
func (twitter *Twitter) Follow(followerID int, followeeID int) {
	follower, ok := twitter.users[followerID]
	if !ok {
		follower = newUser(followerID)
		twitter.users[followerID] = follower
	}
	followee, ok := twitter.users[followeeID]
	if !ok {
		followee = newUser(followeeID)
		twitter.users[followeeID] = followee
	}
	follower.follow(followee)
}

// Unfollow : Follower unfollows a followee. If the operation is invalid, it should be a no-op.
func (twitter *Twitter) Unfollow(followerID int, followeeID int) {
	follower, ok := twitter.users[followerID]
	if !ok {
		follower = newUser(followerID)
		twitter.users[followerID] = follower
	}
	followee, ok := twitter.users[followeeID]
	if !ok {
		followee = newUser(followeeID)
		twitter.users[followeeID] = followee
	} else {
		follower.unfollow(followee)
	}
}

type heapFQ struct {
	data []*tweet
}

func newHeapFQ(initSize int) heapFQ {
	return heapFQ{
		data: make([]*tweet, 0, initSize),
	}
}

func (h *heapFQ) push(val *tweet) {
	h.data = append(h.data, val)
	index := len(h.data) - 1
	for parent := (index - 1) / 2; index > 0 && h.data[index].timestamp > h.data[parent].timestamp; index, parent = parent, (parent-1)/2 {
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
	}
}

func (h *heapFQ) pop() *tweet {
	last := len(h.data) - 1
	val := h.data[0]
	h.data[0], h.data[last] = h.data[last], h.data[0]
	h.data = h.data[:last]
	for index, left := 0, 1; left < last; left = 2*index + 1 {
		max, right := left, left+1
		if right < last && h.data[right].timestamp > h.data[left].timestamp {
			max = right
		}
		if h.data[max].timestamp <= h.data[index].timestamp {
			break
		}
		h.data[index], h.data[max] = h.data[max], h.data[index]
		index = max
	}
	return val
}
