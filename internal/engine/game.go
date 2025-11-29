// Package engine
package engine

import (
	"rogue/internal/core"
	"rogue/internal/spatial"
)

type Game struct {
	World core.World
}

func dungeonStrTransformer(r rune) (bool, bool) {
	switch r {
	case '.':
		return false, true
	case '#':
		return true, true
	default:
		return false, false
	}
}

func NewGame() *Game {
	dungeon := `
		#############
		#....#.#....#
		#....#......#
		#....#.#....#
		##.###.##.###
		#......#....#
		##.##.##....#
		#...........#
		#############
	`
	exitPos := spatial.Coord{X: 10, Y: 6}
	tm, err := spatial.ParseStrMap(dungeon, dungeonStrTransformer)
	if err != nil {
		panic(err)
	}

	// Example of wiring up an actor (though not fully used yet in UI)
	// groblin := &core.BaseActor{
	// 	ActorBlueprint: core.Blueprints["Groblin"],
	// 	Brain:          &ai.SmoothBrain{},
	// }
	// _ = groblin // Prevent unused var error for now until we add actors list to World

	return &Game{
		World: core.World{
			Walls:     tm,
			ExitPos:   exitPos,
			TurnCount: 1,
		},
	}
}

func (g *Game) Update() (core.Actor, error) {
	return g.World.Step()
}

