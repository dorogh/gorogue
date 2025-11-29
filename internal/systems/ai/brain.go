// Package ai provides actor ai routines
package ai

import (
	"rogue/internal/core"
	"rogue/internal/spatial"
)

type SmoothBrain struct{}

func (b *SmoothBrain) Act(actor core.Actor, c spatial.Coord, world *core.World) {
}

func (b *SmoothBrain) PerformAction(action string, actor core.Actor, c spatial.Coord, w *core.World) {
}

// func randomWalk(actor core.Actor, world *core.World) {
// }
