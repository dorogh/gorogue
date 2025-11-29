// Package ui
package ui

import (
	"rogue/internal/core"
	"rogue/internal/engine"
	"rogue/internal/spatial"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Game               *engine.Game
	IsWon              bool
	currentPlayerActor core.Actor
}

func (m Model) Init() tea.Cmd {
	return nil
}

func InitialModel() Model {
	return Model{
		Game: engine.NewGame(),
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
			// case "h":
			// 	m.Game.World.TryMovePlayer(m.Game.World.PlayerPos.Left())
			// case "l":
			// 	m.Game.World.TryMovePlayer(m.Game.World.PlayerPos.Right())
			// case "j":
			// 	m.Game.World.TryMovePlayer(m.Game.World.PlayerPos.Down())
			// case "k":
			// 	m.Game.World.TryMovePlayer(m.Game.World.PlayerPos.Up())
			// }
		}
	}
	// if m.Game.World.PlayerPos == m.Game.World.ExitPos {
	// 	m.IsWon = true
	// }

	return m, nil
}

func (m Model) View() string {
	transformer := func(isWall bool, c spatial.Coord) string {
		switch {
		// case c == m.Game.World.PlayerPos:
		// 	return "🧑"
		case c == m.Game.World.ExitPos:
			return "🚪"
		case isWall:
			return "🧱"
		default: // Floor
			return "⬛"
		}
	}
	// TODO: Add generic Transform method and transform to colored output
	str, err := m.Game.World.Walls.Stringify(transformer)
	if err != nil {
		panic(err)
	}
	if m.IsWon {
		str += "\nYou win!"
	}
	return str
}

