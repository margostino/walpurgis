package action

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/margostino/walpurgis/common"
	"github.com/margostino/walpurgis/config"
	"github.com/margostino/walpurgis/social"
	"strconv"
	"strings"
)

func ExecuteSnapshotUsers() {
	var cursor int64

	cursor = -1
	file := config.TruncateFile()
	defer file.Close()

	for ok := true; ok; ok = cursor != 0 {
		friends, resp, _ := social.TwitterFriends().List(&twitter.FriendListParams{
			Cursor: cursor,
			Count:  200,
		})
		if resp.StatusCode != 429 {
			for _, user := range friends.Users {
				//createdAt, _ := time.Parse("Wed Jan 09 20:56:37 +0000 2019", user.CreatedAt)
				//statusCreatedAt, _ := time.Parse("Wed Jan 09 20:56:37 +0000 2019", user.Status.CreatedAt)
				description := strings.ReplaceAll(strings.ReplaceAll(user.Description, ",", " %44% "), "\n", "")
				statusText := strings.ReplaceAll(strings.ReplaceAll(user.Status.Text, ",", " %44% "), "\n", "")

				text := fmt.Sprintf("%s,%s,%s,%s,%s,%d,%s,%d,%d,%s,%s,%s,%s,%s,%s,%d,%d,%d,%s,%s,%s\n",
					user.IDStr,
					user.ScreenName,
					user.Name,
					user.CreatedAt,
					user.Email,
					user.FavouritesCount,
					strconv.FormatBool(user.FollowRequestSent),
					user.FollowersCount,
					user.FriendsCount,
					strconv.FormatBool(user.GeoEnabled),
					user.Lang,
					user.Location,
					user.Status.CreatedAt,
					strconv.FormatBool(user.Status.Retweeted),
					strconv.FormatBool(user.Status.RetweetedStatus != nil),
					user.Status.RetweetCount,
					user.Status.ReplyCount,
					user.Status.QuoteCount,
					statusText,
					description,
					friends.NextCursorStr)
				_, err := file.WriteString(text)
				common.Check(err)
			}
			cursor = friends.NextCursor
		} else {
			break
		}
	}
}
