package wizard

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewCheckBox(title string, value bool) *CheckBox {
	m := &CheckBox{
		Title: title,
		Value: value,
	}
	m.Styles.TitleStyle = lipgloss.NewStyle()
	m.Styles.FocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	m.Styles.NormalStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	return m
}

// CheckBox model
type CheckBox struct {
	Title     string
	Value     bool
	OnChanged func(bool)
	Styles    struct {
		TitleStyle   lipgloss.Style
		FocusedStyle lipgloss.Style
		NormalStyle  lipgloss.Style
	}
	focus bool
}

func (m CheckBox) Init() tea.Cmd {
	return nil
}

func (m *CheckBox) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.focus {
			switch msg.String() {
			case " ":
				m.setValue(!m.Value)
				return nil
			}
		}
		return nil
	}

	return tea.Batch(cmds...)
}

func (m *CheckBox) View() string {
	title := m.Styles.TitleStyle.Render(m.Title)
	content := "[ ]"
	if m.Value {
		content = "[x]"
	}
	if m.focus {
		content = m.Styles.FocusedStyle.Render(content)
	} else {
		content = m.Styles.NormalStyle.Render(content)
	}
	return title + " " + content
}

func (m *CheckBox) Focus() tea.Cmd {
	m.focus = true
	return nil
}

func (m *CheckBox) Blur() tea.Cmd {
	m.focus = false
	return nil
}

func (m *CheckBox) setValue(v bool) {
	m.Value = v
	if m.OnChanged != nil {
		m.OnChanged(v)
	}
}
