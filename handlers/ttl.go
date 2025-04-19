package handlers

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"os"
)

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func initialModel() model {
	return model{
		choices:  []string{"Sylvia", "Josymar", "Christian", "Raag", "Parker", "Thomas", "Kate", "Konan", "Mario", "Chris"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor]

			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		case "r":
			m.selected = make(map[int]struct{})
		}
	}

	return m, nil
}

func (m model) View() string {
	// The header
	s := "List of ES team members\n\n"

	for i, choice := range m.choices {

		cursor := " " // no cursor
		if m.cursor == i {
			cursor = "→" // cursor!
		}

		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nInstructions:\n\n" + "↑/k: up    ↓/j: down  space/enter: select\nr: reset q: quit"

	// Send the UI for rendering
	return s
}

func TTLHandler(cmd *cobra.Command, args []string) {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
