package wizard

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewNumberInput(title string, value, maxValue int) *NumberInput {
	m := &NumberInput{
		Title:    title,
		Value:    value,
		MaxValue: maxValue,
		MinValue: 0,
		input:    textinput.New(),
	}
	m.input.Validate = func(s string) error {
		if x, err := strconv.Atoi(s); err != nil {
			return err
		} else if x < 0 {
			return fmt.Errorf("negative")
		} else if x > int(m.MaxValue) {
			return fmt.Errorf("too high")
		}
		return nil
	}
	m.input.SetValue(strconv.Itoa(int(value)))
	m.Styles.TitleStyle = lipgloss.NewStyle()
	m.Styles.FocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	m.Styles.NormalStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	return m
}

// NumberInput model
type NumberInput struct {
	Title     string
	Value     int
	MaxValue  int
	MinValue  int
	OnChanged func(int)
	Styles    struct {
		TitleStyle   lipgloss.Style
		FocusedStyle lipgloss.Style
		NormalStyle  lipgloss.Style
	}
	focus bool
	input textinput.Model
}

func (m NumberInput) Init() tea.Cmd {
	return nil
}

func (m *NumberInput) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.focus {
			switch msg.String() {
			case "up":
				m.setValue(m.Value + 1)
				return nil
			case "down":
				m.setValue(m.Value - 1)
				return nil
			case "g":
				m.setValue(0)
				return nil
			case "G":
				m.setValue(int(m.MaxValue) - 1)
				return nil
			}
			var cmd tea.Cmd
			m.input, cmd = m.input.Update(msg)
			cmds = append(cmds, cmd)
			m.setStringValue(m.input.Value())
		}
	}

	return tea.Batch(cmds...)
}

func (m *NumberInput) View() string {
	title := m.Styles.TitleStyle.Render(m.Title)
	content := m.input.View()
	if m.focus {
		content = m.Styles.FocusedStyle.Render(content)
	} else {
		content = m.Styles.NormalStyle.Render(content)
	}
	return title + " " + content
}

func (m *NumberInput) Focus() tea.Cmd {
	m.focus = true
	return m.input.Focus()
}

func (m *NumberInput) Blur() tea.Cmd {
	m.focus = false
	m.input.Blur()
	return nil
}

func (m *NumberInput) setStringValue(v string) {
	if x, err := strconv.Atoi(v); err == nil {
		m.setValue(x)
	}
}

func (m *NumberInput) setValue(x int) {
	if x >= m.MinValue && x <= m.MaxValue {
		m.Value = x
		if m.OnChanged != nil {
			m.OnChanged(x)
		}
		strValue := strconv.Itoa(x)
		if m.input.Value() != strValue {
			m.input.SetValue(strValue)
		}
	}
}
