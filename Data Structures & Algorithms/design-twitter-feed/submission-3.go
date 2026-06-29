type Tweet struct {
	id   int
	time int
}

// We need to store extra context in the heap so we know where to get the next tweet
type HeapItem struct {
	tweet      Tweet
	userId     int
	tweetIndex int // Where we currently are in this user's list
}

// Implement Max-Heap for HeapItem (based on time)
type MaxHeap []HeapItem
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].tweet.time > h[j].tweet.time } // > for Max-Heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(HeapItem)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Twitter struct {
	time    int
	tweets  map[int][]Tweet
	follows map[int]map[int]bool
}

func Constructor() Twitter {
	return Twitter{
		time:    0,
		tweets:  make(map[int][]Tweet),
		follows: make(map[int]map[int]bool),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.time++
	this.tweets[userId] = append(this.tweets[userId], Tweet{id: tweetId, time: this.time})
}

// The Priority Queue Implementation
func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &MaxHeap{}
	heap.Init(h)

	// 1. Compile a list of all relevant users (self + followees)
	usersToCheck := []int{userId}
	if followees, exists := this.follows[userId]; exists {
		for followeeId := range followees {
			usersToCheck = append(usersToCheck, followeeId)
		}
	}

	// 2. Push the MOST RECENT tweet from every user into the Max-Heap
	for _, uid := range usersToCheck {
		userTweets := this.tweets[uid]
		if len(userTweets) > 0 {
			lastIdx := len(userTweets) - 1
			heap.Push(h, HeapItem{
				tweet:      userTweets[lastIdx],
				userId:     uid,
				tweetIndex: lastIdx,
			})
		}
	}

	// 3. Pop the top 10 tweets
	var feed []int
	for len(feed) < 10 && h.Len() > 0 {
		// Get the absolute newest tweet
		item := heap.Pop(h).(HeapItem)
		feed = append(feed, item.tweet.id)

		// Push the NEXT most recent tweet from the same user into the heap
		if item.tweetIndex > 0 {
			nextIdx := item.tweetIndex - 1
			nextTweet := this.tweets[item.userId][nextIdx]
			heap.Push(h, HeapItem{
				tweet:      nextTweet,
				userId:     item.userId,
				tweetIndex: nextIdx,
			})
		}
	}

	return feed
}

func (this *Twitter) Follow(followerId int, followeeId int) {
    if followerId == followeeId {
        return // Guard clause: Cannot follow yourself
    }
    if this.follows[followerId] == nil {
        this.follows[followerId] = make(map[int]bool)
    }
    this.follows[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
    if followerId == followeeId {
        return // Guard clause: Cannot unfollow yourself
    }
    if this.follows[followerId] != nil {
        delete(this.follows[followerId], followeeId)
    }
}