package alexa

func HandleFrontpageDealIntent(request Request) Response {
	return NewSimpleResponse("Frontpage Deals", "Frontpage deal data here")
}

func HandleFallbackIntent(request Request) Response {
	return NewSimpleResponse("Popular Deals", "Popular deal data here")
}

func HandleStopIntent(request Request) Response {
	return NewSimpleResponse("Help", "Help regarding the available commands here")
}

func HandleHelpIntent(request Request) Response {
	return NewSimpleResponse("Help", "Help regarding the available commands here")
}

func HandleCancelIntent(request Request) Response {
	return NewSimpleResponse("Help", "Help regarding the available commands here")
}

func HandleNavigateHomeIntent(request Request) Response {
	return NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleGetNewFactIntent(request Request) Response {
	return NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleAnotherFactIntent(request Request) Response {
	return NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleRepeatIntent(request Request) Response {
	return NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleYesIntent(request Request) Response {
	return NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleNoIntent(request Request) Response {
	return NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func HandleAboutIntent(request Request) Response {
	return NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}
