package wizard

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewButton(title string, onPressed func()) *Button {
	m := &Button{
		Title:     title,
		OnPressed: onPressed,
	}
	m.Styles.FocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	m.Styles.NormalStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	return m
}

// Button model
type Button struct {
	Title     string
	OnPressed func()
	Styles    struct {
		FocusedStyle lipgloss.Style
		NormalStyle  lipgloss.Style
	}
	focus bool
}

func (m Button) Init() tea.Cmd {
	return nil
}

func (m *Button) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.focus {
			switch msg.String() {
			case "enter":
				if m.OnPressed != nil {
					m.OnPressed()
				}
				return nil
			}
		}
		return nil
	}

	return tea.Batch(cmds...)
}

func (m *Button) View() string {
	content := fmt.Sprintf(" [%s] ", m.Title)
	if m.focus {
		content = m.Styles.FocusedStyle.Render(content)
	} else {
		content = m.Styles.NormalStyle.Render(content)
	}
	return content
}

func (m *Button) Focus() tea.Cmd {
	m.focus = true
	return nil
}

func (m *Button) Blur() tea.Cmd {
	m.focus = false
	return nil
}
