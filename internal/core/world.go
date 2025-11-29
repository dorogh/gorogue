// Package core
package core

import (
	"errors"
	"slices"

	"rogue/internal/spatial"
)

var (
	ErrQueueEmpty         = errors.New("no actors in queue")
	ErrActorIsNil         = errors.New("current actor is nil")
	ErrIsPlayerControlled = errors.New("actor is player-controlled")
	ErrTooManyNpcs        = errors.New("handled >30 npcs, relinquishing control")
)

// World - simplified for phase 1: Render procgen'd world and let player move around
type World struct {
	Walls *spatial.TileMap[bool]
	// Goal to 'win' the game
	ExitPos         spatial.Coord
	TurnCount       int
	currentActorIdx int
	turnQueue       []Actor
	pos2Actor       map[spatial.Coord]Actor
	actor2Pos       map[Actor]spatial.Coord
}

func NewWorld(walls *spatial.TileMap[bool], exitPos spatial.Coord) *World {
	return &World{
		Walls:           walls,
		ExitPos:         exitPos,
		TurnCount:       0,
		pos2Actor:       make(map[spatial.Coord]Actor, 0),
		actor2Pos:       make(map[Actor]spatial.Coord, 0),
		turnQueue:       make([]Actor, 0),
		currentActorIdx: 0,
	}
}

func (w *World) PutActor(c spatial.Coord, actor Actor) bool {
	if exists := slices.Contains(w.turnQueue, actor); exists {
		return false
	}
	ok := w.putActorOnMap(c, actor)
	if ok {
		w.turnQueue = append(w.turnQueue, actor)
	}
	return ok
}

func (w *World) putActorOnMap(c spatial.Coord, actor Actor) bool {
	wall, err := w.Walls.At(c)
	if err != nil || *wall {
		return false
	}
	if _, exists := w.pos2Actor[c]; exists {
		return false
	}
	if _, exists := w.actor2Pos[actor]; exists {
		return false
	}
	w.pos2Actor[c] = actor
	w.actor2Pos[actor] = c
	return true
}

func (w *World) RemoveActor(actor Actor) bool {
	idx := slices.Index(w.turnQueue, actor)
	if idx == -1 {
		return false
	}
	ok := w.removeActorFromMap(actor)
	if ok {
		w.turnQueue[idx] = nil
	}
	return ok
}

func (w *World) removeActorFromMap(actor Actor) bool {
	c, exists := w.actor2Pos[actor]
	if !exists {
		return false
	}
	delete(w.actor2Pos, actor)
	delete(w.pos2Actor, c)
	return true
}

// TryMoveActor tries to move the actor to the given target. It does not enforce
// that the target is adjacent to its current position, only that the target
// is not occupied and not identical to its current position.
func (w *World) TryMoveActor(target spatial.Coord, actor Actor) bool {
	wall, err := w.Walls.At(target)
	if err != nil || *wall {
		return false
	}
	if _, occupied := w.pos2Actor[target]; occupied {
		return false
	}
	return w.removeActorFromMap(actor) && w.putActorOnMap(target, actor)
}

// Step picks the next actor and let's it take its turn. If the actor is a player character,
// the actor + ErrIsPlayerControlled is returned and the caller must perform the update.
func (w *World) Step() (Actor, error) {
	if len(w.turnQueue) == 0 {
		return nil, ErrQueueEmpty
	}
	w.currentActorIdx++
	if w.currentActorIdx < 0 || w.currentActorIdx >= len(w.turnQueue) {
		w.currentActorIdx = 0
	}
	actor := w.turnQueue[w.currentActorIdx]
	c, exists := w.actor2Pos[actor]
	if !exists {
		panic("actor in turnQueue, but no on map?")
	}
	if actor.IsPlayerControlled() {
		// Expected to happen, player turn should be handled by ui.
		return actor, ErrIsPlayerControlled
	}
	brain := actor.GetBrain()
	if brain != nil {
		brain.Act(actor, c, w)
	}
	return actor, nil
}

func (w *World) HandleNpcTurns() (Actor, error) {
	var a Actor
	var err error
	for range 30 {
		a, err = w.Step()
		if err != nil {
			return a, err
		}
	}
	return a, ErrTooManyNpcs
}

func (w *World) GetActorAt(c spatial.Coord) Actor {
	return w.pos2Actor[c]
}

func (w *World) GetPosOf(a Actor) (spatial.Coord, bool) {
	c, ok := w.actor2Pos[a]
	return c, ok
}
