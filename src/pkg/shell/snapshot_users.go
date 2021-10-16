package shell

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/margostino/walpurgis/pkg/context"
	"github.com/margostino/walpurgis/pkg/db"
	"github.com/margostino/walpurgis/pkg/helper"
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
				allUsers = append(allUsers, db.User{
					ID:        user.IDStr,
					Username:  user.ScreenName,
					CreatedAt: createdAt,
				})
				text := fmt.Sprintf("%s,%s,%s,%s\n", user.ScreenName, user.IDStr, user.CreatedAt, friends.NextCursorStr)
				_, err := file.WriteString(text)
				helper.Check(err)
			}
			cursor = friends.NextCursor
		} else {
			break
		}
	}
}
