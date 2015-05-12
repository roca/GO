package main

import "errors"

// ErrNoAvatar is the error that is returned when the
// Avatar instance is unable to provide an avatar URL.
var ErrNoAvatarURL = errors.New("chat: Uanable to get an avatar URL.")

// Avatar represents types capable of representing
// user profile pictures

type Avater interface {
	// GetAvatarURL get the avatar URL for the specified client,
	// or returns an error if something goes wrong.
	// ErrNoAvatar is returned if the object is unable to get
	// a URL for the specified client.
	GetAvatarURL(c *client) (string, error)
}
