package core

import "rogue/internal/spatial"

// Brain defines the logic for an actor's behavior.
type Brain interface {
	Act(actor Actor, pos spatial.Coord, world *World)
	// PerformAction should perform or delegate actions such as movement, attacking etc.
	// This is public so the ui layer can pass in player input, but can also be used
	// by the Brain internally to decouple decision making from execution.
	PerformAction(action string, actor Actor, pos spatial.Coord, world *World)
}

type Actor interface {
	GetName() string
	GetGlyph() rune
	GetHealth() int
	IsPlayerControlled() bool
	GetBrain() Brain
}

type ActorBlueprint struct {
	Name             string
	Glyph            rune
	Health           int
	PlayerControlled bool
}

var Blueprints = map[string]ActorBlueprint{
	"Player": {
		Name:             "You",
		Glyph:            '@',
		Health:           100,
		PlayerControlled: true,
	},
	"Groblin": {
		Name:             "Groblin",
		Glyph:            'g',
		Health:           20,
		PlayerControlled: false,
	},
}

type BaseActor struct {
	ActorBlueprint
	Brain Brain
}

func (a *BaseActor) GetName() string {
	return a.Name
}

func (a *BaseActor) GetGlyph() rune {
	return a.Glyph
}

func (a *BaseActor) GetHealth() int {
	return a.Health
}

func (a *BaseActor) IsPlayerControlled() bool {
	return a.PlayerControlled
}

func (a *BaseActor) GetBrain() Brain {
	return a.Brain
}

// Removed Act method from BaseActor as it wasn't doing anything and isn't part of the interface
// (well, the interface has GetBrain, the Brain has Act)

