package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	reader := strings.NewReader("this is hte request body content")
	r, err := http.NewRequest("GET", "https://google.com", reader)
	r.Header.Add("Authorization", "ApiKey 1234")
	if err != nil {
		t.Fatalf("error making request")
	}
	_, err = GetAPIKey(r.Header)
	if err != nil {
		t.Fatalf("error getting api key making request")
	}

}
