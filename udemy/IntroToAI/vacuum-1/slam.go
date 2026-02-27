package main

func CleanRoomSlam(room *Room, robot *Robot) {
	// Set start time and movecount

	// Initialize the robot's intyernal map

	// Initialize visiuted cells (tracking)

	// Initialize a frontier

	// Mark starting position as visited and update map for the first time

	// Clean the current position

	// Add neighbors to the frontier

	// Display the initial state

	// for - if the frontier is not empty and the room is not fully cleaned
	for {

		// Get closest frontier point

		// If not valid target, break

		// remove target from frontier

		// find path to target using A*

		// If not path found, go to next frontier point (continue)

		// Move along the path
		for {

			// Update map (internal)

		}

		// every 10 minutes, do a more thorough frontier check

		// check if we have sufficient coverage -- break
	}

	// final cleanup phase

}
