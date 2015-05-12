package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar authAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no vaule present")
	}
	// set value
	testUrl := "http://url-to-gravatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}

	url, err := authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should return no error when value present")
	}
	if url != testUrl {
		t.Error("AuthAvatar.GetAvatarURL should return correct URL")
	}
}
