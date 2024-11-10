package mouse

import tea "github.com/charmbracelet/bubbletea"

// IsLeftButtonPressed checks if the left mouse button is pressed.
func IsLeftButtonPressed(msg tea.MouseMsg) bool {
	event := tea.MouseEvent(msg)

	return event.Button == tea.MouseButtonLeft && event.Action == tea.MouseActionPress
}

// IsClickWithinArea checks if the mouse click is within the specified area.
func IsClickWithinArea(msg tea.MouseMsg, xMin, yMin, xMax, yMax int) bool {
	return msg.X >= xMin && msg.X <= xMax && msg.Y >= yMin && msg.Y <= yMax
}
