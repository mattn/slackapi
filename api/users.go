package api

import "net/url"

// users.getPresence
// users.info

// GetUserList returns the slack domain user list
// users.list
func GetUserList(token string) (*UserListResponse, error) {
	userList := &UserListResponse{}
	err := getJSON(
		"https://slack.com/api/users.list",
		url.Values{"token": {token}},
		userList)

	return userList, err
}

// users.setActive
// users.setPresence
