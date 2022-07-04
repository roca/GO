package deadlock

import (
	"sync"
	"time"

	. "github.com/roca/GO/tree/staging/udemy/multithreadingingo/deadlocks_train/common"
)

var (
	controller = sync.Mutex{}
	cond       = sync.NewCond(&controller)
)

func allFree(intersectionsToLock []*Intersection) bool {
	for _, it := range intersectionsToLock {
		if it.LockedBy != -1 {
			return false
		}
	}
	return true
}

func lockIntersectionsInDistance(id, reserveStart, reserveEnd int, crossings []*Crossing) {
	var intersectionsToLock []*Intersection
	for _, crossing := range crossings {
		if reserveEnd >= crossing.Position && reserveStart <= crossing.Position && crossing.Intersections.LockedBy != id {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersections)
		}
	}

	controller.Lock()
	for !allFree(intersectionsToLock) {
		cond.Wait()
	}

	for _, it := range intersectionsToLock {
		it.LockedBy = id
		time.Sleep(30 * time.Millisecond)
	}
	controller.Unlock()
}

func MoveTrain(train *Train, distance int, crossings []*Crossing) {
	for train.Front < distance {
		train.Front += 1
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				lockIntersectionsInDistance(train.Id, crossing.Position, crossing.Position+train.TrainLength, crossings)
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				controller.Lock()
				crossing.Intersections.LockedBy = -1
				cond.Broadcast()
				controller.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}

}
