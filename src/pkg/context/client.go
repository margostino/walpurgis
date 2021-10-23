package context

import (
	"github.com/dghubble/go-twitter/twitter"
)

func TwitterTimelines() *twitter.TimelineService {
	return appContext.Twitter.Timelines
}

func TwitterUsers() *twitter.UserService {
	return appContext.Twitter.Users
}

func TwitterFriends() *twitter.FriendService {
	return appContext.Twitter.Friends
}
