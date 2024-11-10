package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ivanlemeshev/gameoflife/internal/game/grid"
	"github.com/ivanlemeshev/gameoflife/internal/game/mouse"
)

const (
	speed         = 500 * time.Millisecond
	setGreenColor = "\033[32m"
	resetColor    = "\033[0m"
)

type tickMsg time.Time

// Game represents the bubbletea model for the game.
type Game struct {
	started bool
	width   int
	height  int
	grid    *grid.Grid
	spinner spinner.Model
	keys    keyMap
}

// New creates a new game with the specified width and height for the grid.
func New(width, height int) *Game {
	return &Game{
		width:   width,
		height:  height,
		grid:    grid.New(width, height),
		spinner: newSpinner(),
		keys:    gameKeys,
	}
}

// Init initializes the game.
func (g *Game) Init() tea.Cmd {
	return nil
}

// Update updates the game state depending on the message received.
func (g *Game) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		return g.handlePressedKey(msg)
	case tea.MouseMsg:
		return g.handleMouseEvent(msg)
	case spinner.TickMsg:
		return g.handleSpinnerTick(msg)
	case tickMsg:
		return g.handleTick()
	}
	return g, nil
}

// View returns the game view.
func (g *Game) View() string {
	var sb strings.Builder

	sb.WriteString("============================ Conway's Game of Life ============================\n")
	sb.WriteString("Use the mouse cursor and the left button to set the cell state.                \n")
	sb.WriteString("Press 'r' to reset the game, '␣' to start/pause the game, 'q' to quit the game.\n")
	sb.WriteString("===============================================================================\n")

	// Render the grid.
	currentState := g.grid.State()
	for y := 0; y < len(currentState); y++ {
		for x := 0; x < len(currentState[y]); x++ {
			currentCell := currentState[y][x]
			if currentCell == nil {
				sb.WriteString("□ ")
				continue
			}

			sb.WriteString(setGreenColor)
			sb.WriteString("■ ")
			sb.WriteString(resetColor)
		}

		sb.WriteString("\n")
	}

	// Render the generation number.
	generation := fmt.Sprintf("%s Generation: %d\n", g.spinner.View(), g.grid.Generation())
	sb.WriteString(generation)

	return sb.String()
}

func (g *Game) handleSpinnerTick(msg spinner.TickMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	g.spinner, cmd = g.spinner.Update(msg)
	return g, cmd
}

func (g *Game) handlePressedKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, g.keys.Quit):
		// Quit the game.
		return g, tea.Quit
	case key.Matches(msg, g.keys.Reset):
		// Reset the game.
		g.started = false
		g.resetGrid()
		g.resetSpinner()

		return g, nil
	case key.Matches(msg, g.keys.ToggleStartPause):
		// Start or pause the game.
		g.started = !g.started
		if g.started {
			return g, tea.Batch(g.tick(), g.spinner.Tick)
		}

		g.resetSpinner()

		return g, nil
	default:
		// Ignore other keys.
		if g.started {
			return g, g.tick()
		}

		return g, nil
	}
}

func (g *Game) handleMouseEvent(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	// We can set the cells only if the game is not started yet or during the
	// pause.
	if g.started {
		return g, nil
	}

	// We can set the cell state only by pressing the left mouse button.
	// All other buttons and actions will be ignored.
	if !mouse.IsLeftButtonPressed(msg) {
		return g, nil
	}

	// We have a spaces between cells in rows, so we need to multiply the width
	// by 2.
	gridXMin := 0
	gridXMax := g.width * 2

	// The first three lines is the title and help, we need to shift the grid by one line down.
	gridYMin := 4
	gridYMax := g.height + gridYMin

	// We can handle only mouse clicks withing the cell grid in terminal.
	if !mouse.IsClickWithinArea(msg, gridXMin, gridYMin, gridXMax, gridYMax) {
		return g, nil
	}
	// Each cell is represented by width 2 (one for the cell and one for
	// the space). We need to handle only the cell click.
	if msg.X%2 == 0 {
		g.grid.ToggleCell(msg.X/2, msg.Y-gridYMin)
	}

	return g, nil
}

func (g *Game) handleTick() (tea.Model, tea.Cmd) {
	if !g.started {
		return g, nil
	}

	g.grid.NextGeneration()

	return g, g.tick()
}

func (g *Game) tick() tea.Cmd {
	return tea.Tick(speed, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (g *Game) resetSpinner() {
	g.spinner = newSpinner()
}

func (g *Game) resetGrid() {
	g.grid = grid.New(g.width, g.height)
}

func newSpinner() spinner.Model {
	return spinner.New(spinner.WithSpinner(spinner.Globe))
}
