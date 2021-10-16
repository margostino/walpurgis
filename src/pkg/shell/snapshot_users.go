package shell

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/margostino/walpurgis/pkg/context"
	"github.com/margostino/walpurgis/pkg/db"
	"github.com/margostino/walpurgis/pkg/helper"
	"strconv"
	"strings"
	"time"
)

func ExecuteSnapshotUsers() {
	var cursor int64
	var allUsers = make([]db.User, 0)

	cursor = -1
	file := db.OpenFile()
	defer file.Close()

	for ok := true; ok; ok = cursor != 0 {
		friends, resp, _ := context.TwitterFriends().List(&twitter.FriendListParams{
			Cursor: cursor,
			Count:  200,
		})
		if resp.StatusCode != 429 {
			for _, user := range friends.Users {
				createdAt, _ := time.Parse("Wed Jan 09 20:56:37 +0000 2019", user.CreatedAt)
				description := strings.ReplaceAll(user.Description, ",", " %44% ")
				statusCreatedAt, _ := time.Parse("Wed Jan 09 20:56:37 +0000 2019", user.Status.CreatedAt)

				allUsers = append(allUsers, db.User{
					ID:        user.IDStr,
					Username:  user.ScreenName,
					CreatedAt: createdAt,
				})
				text := fmt.Sprintf("%s,%s,%s,%s,%s,%d,%s,%d,%d,%s,%s,%s,%s,%s,%s,%d,%d,%d,%s,%s,%s\n",
					user.ScreenName,
					user.Name,
					user.IDStr,
					user.CreatedAt,
					user.Email,
					user.FavouritesCount,
					user.FollowRequestSent,
					user.FollowersCount,
					user.FriendsCount,
					strconv.FormatBool(user.GeoEnabled),
					user.Lang,
					user.Location,
					statusCreatedAt,
					strconv.FormatBool(user.Status.Retweeted),
					user.Status.RetweetedStatus != nil,
					user.Status.RetweetCount,
					user.Status.ReplyCount,
					user.Status.QuoteCount,
					user.Status.Text,
					description,
					friends.NextCursorStr)
				_, err := file.WriteString(text)
				helper.Check(err)
			}
			cursor = friends.NextCursor
		} else {
			break
		}
	}
}
