package social

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/margostino/walpurgis/config"
)

func TwitterTimelines() *twitter.TimelineService {
	return config.Context.TwitterClient.Timelines
}

func TwitterUsers() *twitter.UserService {
	return config.Context.TwitterClient.Users
}

func TwitterFriends() *twitter.FriendService {
	return config.Context.TwitterClient.Friends
}
