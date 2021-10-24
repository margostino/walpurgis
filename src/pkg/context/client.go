package context

import (
	"github.com/dghubble/go-twitter/twitter"
)

func TwitterTimelines() *twitter.TimelineService {
	return appContext.TwitterClient.Timelines
}

func TwitterUsers() *twitter.UserService {
	return appContext.TwitterClient.Users
}

func TwitterFriends() *twitter.FriendService {
	return appContext.TwitterClient.Friends
}
