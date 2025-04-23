package asteroids

import "github.com/solarlune/resolv"

func (g *GameScene) checkCollision(obj, against *resolv.Circle) bool {
	if against == nil {
		return obj.IntersectionTest(resolv.IntersectionTestSettings{
			TestAgainst: obj.SelectTouchingCells(1).FilterShapes(),
			OnIntersect: func(set resolv.IntersectionSet) bool {
				return true
			},
		})
	}

	return obj.IntersectionTest(resolv.IntersectionTestSettings{
		TestAgainst: against.SelectTouchingCells(1).FilterShapes(),
		OnIntersect: func(set resolv.IntersectionSet) bool {
			return true
		},
	})
}
