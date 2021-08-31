package main

import (
	"bufio"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/margostino/walpurgis/pkg/config"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type UserInfo struct {
	Username     string
	LastActivity time.Time
}

type User struct {
	ID        string
	Username  string
	CreatedAt time.Time
}

var client *twitter.Client

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getAllUsers() []User {
	var cursor int64
	//var allUsers = make([]twitter.User, 0)
	var allUsers = make([]User, 0)
	cursor = -1
	filename := "../data/users"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		cursor, _ = strconv.ParseInt(values[3], 10, 64)
		createdAt, _ := time.Parse("Wed Jan 09 20:56:37 +0000 2019", values[2])
		allUsers = append(allUsers, User{
			ID:        values[1],
			Username:  values[0],
			CreatedAt: createdAt,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for ok := true; ok; ok = cursor != 0 {
		friends, resp, _ := client.Friends.List(&twitter.FriendListParams{
			Cursor: cursor,
		})
		if resp.StatusCode != 429 {
			for _, user := range friends.Users {
				createdAt, _ := time.Parse("Wed Jan 09 20:56:37 +0000 2019", user.CreatedAt)
				allUsers = append(allUsers, User{
					ID:        user.IDStr,
					Username:  user.ScreenName,
					CreatedAt: createdAt,
				})
				//allUsers = append(allUsers, user)
				text := fmt.Sprintf("%s,%s,%s,%s\n", user.ScreenName, user.IDStr, user.CreatedAt, friends.NextCursorStr)
				_, err := file.WriteString(text)
				check(err)
			}
			cursor = friends.NextCursor
		} else {
			break
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
		var lastActivity time.Time
		id, _ := strconv.ParseInt(user.ID, 10, 64)
		tweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
			UserID: id,
			Count:  1,
		})

		if len(tweets) > 0 {
			lastActivity, _ = tweets[0].CreatedAtTime()
		}
		profiles = append(profiles, UserInfo{Username: user.Username, LastActivity: lastActivity})
	}

	sort.SliceStable(profiles, func(i, j int) bool {
		return profiles[i].LastActivity.Before(profiles[j].LastActivity)
	})

	for _, profile := range profiles {
		fmt.Printf("[%s] - %s]\n", profile.Username, profile.LastActivity)
	}

}
