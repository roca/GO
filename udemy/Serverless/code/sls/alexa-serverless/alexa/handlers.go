package alexa

func HandleFallbackIntent(request Request) Response {
	return NewSimpleResponse("Fallback", "Popular deal data here")
}

func HandleStopIntent(request Request) Response {
	return NewSimpleResponse("Stop", "Help regarding the available commands here")
}

func HandleHelpIntent(request Request) Response {
	return NewSimpleResponse("Help", "Help regarding the available commands here")
}

func HandleCancelIntent(request Request) Response {
	return NewSimpleResponse("Cancel", "Help regarding the available commands here")
}

func HandleNavigateHomeIntent(request Request) Response {
	return NewSimpleResponse("Navigate", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleGetNewFactIntent(request Request) Response {
	return NewSimpleResponse("Get New Fact", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleAnotherFactIntent(request Request) Response {
	return NewSimpleResponse("Another Fact", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleRepeatIntent(request Request) Response {
	return NewSimpleResponse("Repeat", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleYesIntent(request Request) Response {
	return NewSimpleResponse("Yes", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleNoIntent(request Request) Response {
	return NewSimpleResponse("No", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleAboutIntent(request Request) Response {
	return NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}
