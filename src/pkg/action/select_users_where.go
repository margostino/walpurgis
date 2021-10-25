package action

import (
	"fmt"
	"github.com/margostino/walpurgis/pkg/context"
	"strings"
)

func ExecuteSelectUsersWhere(args []string) {
	users := context.GetUsersData()

	for _, user := range users {
		match := false
		key := getKey(args[0], user)
		switch strings.ToLower(args[1]) {
		case "like":
			match = strings.Contains(key, args[2])
		case "=":
			match = key == args[2]
		case "not like":
			match = !strings.Contains(key, args[2])
		}

		if match {
			fmt.Printf("[%s] - %s]\n", user.Username, user.Description)
		}
	}
}

func getKey(param string, user *context.User) string {
	switch strings.ToLower(param) {
	case "description":
		return user.Description
	case "email":
		return user.Email
	case "status":
		return user.StatusText
	case "name":
		return user.Name
	default:
		return ""
	}
}
