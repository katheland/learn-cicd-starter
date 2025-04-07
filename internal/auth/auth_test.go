package auth

import (
	"net/http"
	"testing"
)

func TestAPIKey(t *testing.T) {
	goodTest := http.Header{}
	goodTest.Add("Authorization", "ApiKey test")
	notGoodTest := http.Header{}
	notGoodTest.Add("Authorization", "wrong")
	badTest := http.Header{}

	goodText, err := GetAPIKey(goodTest)
	if err != nil {
		t.Errorf("Error in GetAPIKey: %v", err)
	}
	if goodText != "test" {
		t.Errorf("GetAPIKey wrong response: %v", goodText)
	}

	_, err = GetAPIKey(notGoodTest)
	if err == nil {
		t.Errorf("There was supposed to be an error in GetAPIKey")
	}

	_, err = GetAPIKey(badTest)
	if err == nil {
		t.Errorf("There was supposed to be an error in GetAPIKey")
	}
}
