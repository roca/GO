package alexa

import (
	"math/rand"
	"time"
)

func HandleFallbackIntent(request Request) Response {
	return NewSimpleResponse("Fallback", FALLBACKMESSAGE, FALLBACKREPROMPT)
}

func HandleStopIntent(request Request) Response {
	return NewSimpleResponse("Stop", STOPMESSAGE, ERRORMESSAGE)
}

func HandleHelpIntent(request Request) Response {
	return NewSimpleResponse("Help", HELPMESSAGE, HELPREPROMPT)
}

func HandleCancelIntent(request Request) Response {
	return NewSimpleResponse("Cancel", STOPMESSAGE, ERRORMESSAGE)
}

func HandleNavigateHomeIntent(request Request) Response {
	return NewSimpleResponse("Navigate", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.", "")
}

func HandleGetNewFactIntent(request Request) Response {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return NewSimpleResponse("Get New Fact", DATA[r.Intn(len(DATA))], "")
}

func HandleAnotherFactIntent(request Request) Response {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return NewSimpleResponse("Another Fact", DATA[r.Intn(len(DATA))], "")
}

func HandleRepeatIntent(request Request) Response {
	return NewSimpleResponse("Repeat", REPEATMESSAGE, CONTINUEREPROMPT)
}

func HandleYesIntent(request Request) Response {
	return NewSimpleResponse("Yes", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.", "")
}

func HandleNoIntent(request Request) Response {
	return NewSimpleResponse("No", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.", "")
}

func HandleAboutIntent(request Request) Response {
	return NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.", "")
}