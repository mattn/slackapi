package api

import (
	"net/http"
	"net/url"
)

// AuthTest tests if our auth token is good
// Checks authentication & identity.
func AuthTest(token string) (*http.Response, error) {
	return http.PostForm("https://slack.com/api/auth.test",
		url.Values{"token": {token}})
}
