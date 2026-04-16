
// -------------------- Tweet --------------------

type Tweet struct {
	id   int
	time int
	next *Tweet
}

// -------------------- Twitter --------------------

type Twitter struct {
	time      int
	followMap map[int]map[int]bool
	tweetMap  map[int]*Tweet
}

func Constructor() Twitter {
	return Twitter{
		time:      0,
		followMap: make(map[int]map[int]bool),
		tweetMap:  make(map[int]*Tweet),
	}
}

// -------------------- Post Tweet --------------------

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.time++

	tweet := &Tweet{
		id:   tweetId,
		time: this.time,
		next: this.tweetMap[userId],
	}

	this.tweetMap[userId] = tweet
}

// -------------------- Heap --------------------

type MaxHeap []*Tweet

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool {
	return h[i].time > h[j].time // max heap
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(*Tweet))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	tweet := old[n-1]
	*h = old[:n-1]
	return tweet
}

// -------------------- Get News Feed --------------------

func (this *Twitter) GetNewsFeed(userId int) []int {
	res := []int{}
	h := &MaxHeap{}
	heap.Init(h)

	// ensure user follows themselves
	if this.followMap[userId] == nil {
		this.followMap[userId] = make(map[int]bool)
	}
	this.followMap[userId][userId] = true

	// push latest tweet of each followee
	for followee := range this.followMap[userId] {
		if tweet := this.tweetMap[followee]; tweet != nil {
			heap.Push(h, tweet)
		}
	}

	// get top 10 recent tweets
	for h.Len() > 0 && len(res) < 10 {
		tweet := heap.Pop(h).(*Tweet)
		res = append(res, tweet.id)

		if tweet.next != nil {
			heap.Push(h, tweet.next)
		}
	}

	return res
}

// -------------------- Follow --------------------

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.followMap[followerId] == nil {
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true
}

// -------------------- Unfollow --------------------

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.followMap[followerId] != nil {
		delete(this.followMap[followerId], followeeId)
	}
}