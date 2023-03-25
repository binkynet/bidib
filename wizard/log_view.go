package wizard

import (
	"bytes"
	"io"
	"time"

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
	return m
}

// Viewer for logs
type LogView struct {
	view viewport.Model
	lb   *logBuffer
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

	switch msg.(type) {
	case logChangedMsg:
		m.view.SetContent(m.lb.String())
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
