package mouse_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ivanlemeshev/gameoflife/internal/game/mouse"
	"github.com/stretchr/testify/assert"
)

func TestIsLeftButtonPressed(t *testing.T) {
	tt := []struct {
		name     string
		button   tea.MouseButton
		action   tea.MouseAction
		expected bool
	}{
		{
			name:     "left mouse button is pressed",
			button:   tea.MouseButtonLeft,
			action:   tea.MouseActionPress,
			expected: true,
		},
		{
			name:     "other button is pressed",
			button:   tea.MouseButtonRight,
			action:   tea.MouseActionPress,
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			msg := tea.MouseMsg{
				Button: tc.button,
				Action: tc.action,
			}
			assert.Equal(t, tc.expected, mouse.IsLeftButtonPressed(msg))
		})
	}
}

func TestIsClickWithinArea(t *testing.T) {
	minX := 5
	maxX := 10
	minY := 5
	maxY := 10

	tt := []struct {
		name        string
		mouseEventX int
		mouseEventY int
		expected    bool
	}{
		{
			name:        "click is within the area",
			mouseEventX: 7,
			mouseEventY: 7,
			expected:    true,
		},
		{
			name:        "click is left of the area",
			mouseEventX: 3,
			mouseEventY: 7,
			expected:    false,
		},
		{
			name:        "click is right of the area",
			mouseEventX: 12,
			mouseEventY: 7,
			expected:    false,
		},
		{
			name:        "click is above the area",
			mouseEventX: 7,
			mouseEventY: 3,
			expected:    false,
		},
		{
			name:        "click is below the area",
			mouseEventX: 7,
			mouseEventY: 12,
			expected:    false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			msg := tea.MouseMsg{
				X: tc.mouseEventX,
				Y: tc.mouseEventY,
			}
			assert.Equal(t, tc.expected, mouse.IsClickWithinArea(msg, minX, minY, maxX, maxY))
		})
	}
}
