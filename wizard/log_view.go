package wizard

import (
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func NewLogView() LogView {
	m := LogView{
		view: viewport.New(0, 0),
		lb:   &logBuffer{},
	}
	m.lb.changed = make(chan logChangedMsg, 32)
	m.view.SetContent("initial content")
	m.keyMap.GotoStart = key.NewBinding(
		key.WithKeys("home", "g"),
		key.WithHelp("g/home", "go to start"),
	)
	m.keyMap.GotoEnd = key.NewBinding(
		key.WithKeys("end", "G"),
		key.WithHelp("G/end", "go to end"),
	)
	return m
}

// Viewer for logs
type LogView struct {
	view   viewport.Model
	lb     *logBuffer
	keyMap struct {
		GotoStart key.Binding
		GotoEnd   key.Binding
	}
}

type logChangedMsg struct{}

type logBuffer struct {
	bytes.Buffer
	changed chan logChangedMsg
}

func (lb *logBuffer) Write(p []byte) (int, error) {
	n, err := lb.Buffer.Write(p)
	go func() {
		select {
		case lb.changed <- logChangedMsg{}:
			// Ok
		case <-time.After(time.Millisecond * 500):
			panic(p)
			// Ignore
		}
	}()
	return n, err
}

func (m LogView) Writer() io.Writer {
	return m.lb
}

func (m LogView) Init() tea.Cmd {
	return m.onChanged()
}

func (m LogView) onChanged() tea.Cmd {
	return func() tea.Msg {
		return <-m.lb.changed
	}
}

func (m LogView) Update(msg tea.Msg) (LogView, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keyMap.GotoStart) {
			m.view.SetYOffset(0)
			return m, nil
		} else if key.Matches(msg, m.keyMap.GotoEnd) {
			m.view.SetYOffset(m.view.TotalLineCount() - m.view.Height)
			return m, nil
		}
	case logChangedMsg:
		atBottom := m.view.AtBottom()
		m.view.SetContent(strings.TrimSpace(m.lb.String()))
		if atBottom {
			m.view.SetYOffset(m.view.TotalLineCount() - m.view.Height)
		}
		return m, m.onChanged()
	}

	var cmd tea.Cmd
	m.view, cmd = m.view.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m LogView) View() string {
	return m.view.View()
}

func (m *LogView) SetSize(w, h int) {
	m.view.Width = w
	m.view.Height = h
}
