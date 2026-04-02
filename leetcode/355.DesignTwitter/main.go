package main

import (
	"fmt"
)

type User struct {
	id        int
	following map[int]bool
	tweets    []*Tweet
}

func NewUser(id int) *User {
	return &User{
		id:        id,
		following: make(map[int]bool),
		tweets:    []*Tweet{},
	}
}

type Tweet struct {
	id     int
	userId int
}

func NewTweet(id int, userId int) *Tweet {
	return &Tweet{
		id:     id,
		userId: userId,
	}
}

type Twitter struct {
	users  map[int]*User
	tweets []*Tweet
}

func Constructor() Twitter {
	return Twitter{
		users:  make(map[int]*User),
		tweets: []*Tweet{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	t := NewTweet(tweetId, userId)
	this.tweets = append(this.tweets, t)

	if _, ok := this.users[userId]; !ok {
		this.users[userId] = NewUser(userId)
	}

	this.users[userId].tweets = append(this.users[userId].tweets, t)

}

func (this *Twitter) GetNewsFeed(userId int) []int {
	if _, ok := this.users[userId]; !ok {
		return []int{}
	}

	user := this.users[userId]
	following := user.following

	res := []int{}
	for i := len(this.tweets) - 1; i >= 0 && len(res) < 10; i-- {
		tweet := this.tweets[i]
		if tweet.userId == userId || following[tweet.userId] {
			res = append(res, tweet.id)
		}
	}

	return res

}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.users[followerId]; !ok {
		this.users[followerId] = NewUser(followerId)
	}
	if _, ok := this.users[followeeId]; !ok {
		this.users[followeeId] = NewUser(followeeId)
	}
	this.users[followerId].following[followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.users[followerId]; ok {
		delete(this.users[followerId].following, followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

func main() {
	x := 0
	nums := []int{}
	fmt.Println(x, nums)
}
