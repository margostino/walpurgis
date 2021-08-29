package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/margostino/walpurgis/pkg/config"
	"time"
)

type UserInfo struct {
	Username     string
	LastActivity time.Time
}

var client *twitter.Client

func getAllUsers() []twitter.User {
	var cursor int64
	var allUsers = make([]twitter.User, 0)
	cursor = -1
	for ok := true; ok; ok = cursor != 0 {
		friends, resp, _ := client.Friends.List(&twitter.FriendListParams{
			Cursor: cursor,
		})
		if resp.StatusCode != 429 {
			for _, user := range friends.Users {
				allUsers = append(allUsers, user)
			}
			cursor = friends.NextCursor
		}
	}
	return allUsers
}

func main() {
	var profiles = make([]UserInfo, 0)
	configuration := config.GetConfiguration()
	twitterConfig := oauth1.NewConfig(configuration.Twitter.ApiKey, configuration.Twitter.ApiSecret)
	token := oauth1.NewToken(configuration.Twitter.AccessKey, configuration.Twitter.AccessSecret)
	httpClient := twitterConfig.Client(oauth1.NoContext, token)
	client = twitter.NewClient(httpClient)

	for _, user := range getAllUsers() {
		var lastActivity *time.Time
		tweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
			UserID: user.ID,
			Count:  1,
		})

		if len(tweets) > 0 {
			*lastActivity, _ = tweets[0].CreatedAtTime()
		} else {
			lastActivity = nil
		}
		profiles = append(profiles, UserInfo{Username: user.ScreenName, LastActivity: *lastActivity})

	}
	fmt.Printf("end")

}
