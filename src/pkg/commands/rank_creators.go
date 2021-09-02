package commands

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/margostino/walpurgis/pkg/client"
	"github.com/margostino/walpurgis/pkg/db"
	"sort"
	"strconv"
	"time"
)

// TODO: call in parallel, improve performance

func RankCreators() {
	var profiles = make([]db.UserInfo, 0)
	users := db.LoadUsersData()
	for _, user := range users {
		var lastActivity time.Time
		id, _ := strconv.ParseInt(user.ID, 10, 64)
		tweets, _, _ := client.Twitter.Timelines.UserTimeline(&twitter.UserTimelineParams{
			UserID: id,
			Count:  1,
		})

		if len(tweets) > 0 {
			lastActivity, _ = tweets[0].CreatedAtTime()
		}
		profiles = append(profiles, db.UserInfo{Username: user.Username, LastActivity: lastActivity})
	}

	sort.SliceStable(profiles, func(i, j int) bool {
		return profiles[i].LastActivity.Before(profiles[j].LastActivity)
	})

	for _, profile := range profiles {
		fmt.Printf("[%s] - %s]\n", profile.Username, profile.LastActivity)
	}
}
