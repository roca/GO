package deadlock

import (
	"time"

	. "github.com/roca/GO/tree/staging/udemy/multithreadingingo/deadlocks_train/common"
)

func MoveTrain(train *Train, distance int, crossings []*Crossing) {
	for train.Front < distance {
		train.Front += 1
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				crossing.Intersections.Mutex.Lock()
				crossing.Intersections.LockedBy = train.Id
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				crossing.Intersections.LockedBy = -1
				crossing.Intersections.Mutex.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}

}
