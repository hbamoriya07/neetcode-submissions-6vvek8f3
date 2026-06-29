type Tweet struct {
	id   int
	time int // Global timestamp to maintain chronological order
}

type Twitter struct {
	time    int
	tweets  map[int][]Tweet      // userId -> list of their tweets
	follows map[int]map[int]bool // followerId -> set of followeeIds
}

func Constructor() Twitter {
	return Twitter{
		time:    0,
		tweets:  make(map[int][]Tweet),
		follows: make(map[int]map[int]bool),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.time++ // Increment global time for every new tweet
	this.tweets[userId] = append(this.tweets[userId], Tweet{id: tweetId, time: this.time})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	var feed []Tweet

	// 1. Get the user's own recent tweets (at most 10)
	userTweets := this.tweets[userId]
	for i := len(userTweets) - 1; i >= 0 && i >= len(userTweets)-10; i-- {
		feed = append(feed, userTweets[i])
	}

	// 2. Get the recent tweets of everyone the user follows
	if followees, exists := this.follows[userId]; exists {
		for followeeId := range followees {
            // Edge case check: user might technically "follow" themselves in some test cases
			if followeeId != userId { 
				fTweets := this.tweets[followeeId]
				for i := len(fTweets) - 1; i >= 0 && i >= len(fTweets)-10; i-- {
					feed = append(feed, fTweets[i])
				}
			}
		}
	}

	// 3. Sort the aggregated pool of tweets by time (Most recent first)
	sort.Slice(feed, func(i, j int) bool {
		return feed[i].time > feed[j].time
	})

	// 4. Extract and return the top 10 tweet IDs
	var res []int
	for i := 0; i < len(feed) && i < 10; i++ {
		res = append(res, feed[i].id)
	}

	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.follows[followerId] == nil {
		this.follows[followerId] = make(map[int]bool)
	}
	this.follows[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.follows[followerId] != nil {
		delete(this.follows[followerId], followeeId)
	}
}