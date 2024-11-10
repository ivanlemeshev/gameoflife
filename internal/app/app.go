package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/ivanlemeshev/gameoflife/internal/game"
)

const (
	gameGridWidth  = 40
	gameGridHeight = 18
)

// App is the main application structure.
type App struct {
	program *tea.Program
}

// New creates a new application and initializes it.
func New() *App {
	return &App{
		program: tea.NewProgram(
			game.New(gameGridWidth, gameGridHeight),
			tea.WithAltScreen(),
			tea.WithMouseAllMotion()),
	}
}

// Run starts the application.
func (a *App) Run() error {
	_, err := a.program.Run()
	return err
}
