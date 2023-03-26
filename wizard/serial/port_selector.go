package serial

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

// PortSelector to select a serial port
type PortSelector struct {
	ports []portPath

	list         list.Model
	selectedPath string
}

// New returns a new serial port selector with sensible defaults
func New() PortSelector {
	m := PortSelector{
		list: list.New(nil, list.NewDefaultDelegate(), 0, 0),
	}
	m.list.Title = "Select serial port"
	m.list.SetShowStatusBar(false)
	return m
}

type portPath string

func (i portPath) Title() string       { return string(i) }
func (i portPath) Description() string { return string(i) }
func (i portPath) FilterValue() string { return string(i) }

type portsLoadedMsg []portPath

func (m PortSelector) Init() tea.Cmd {
	return func() tea.Msg {
		// List serial ports
		ports := listSerialPorts()
		return portsLoadedMsg(ports)
	}
}

func (m PortSelector) Update(msg tea.Msg) (PortSelector, tea.Cmd) {
	switch msg := msg.(type) {
	case portsLoadedMsg:
		m.ports = msg
		items := make([]list.Item, 0, len(msg))
		for _, item := range msg {
			items = append(items, item)
		}
		m.list.SetItems(items)
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if item := m.list.SelectedItem(); item != nil {
				m.selectedPath = string(item.(portPath))
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m PortSelector) View() string {
	return docStyle.Render(m.list.View())
}

// Returns the selected path of the serial port or empty string
func (m PortSelector) SelectedPath() string {
	return m.selectedPath
}

// Set width/height of the list.
func (m *PortSelector) SetSize(w, h int) {
	m.list.SetSize(w, h)
}

func listSerialPorts() []portPath {
	entries, err := os.ReadDir("/dev")
	if err != nil {
		return nil
	}
	var result []portPath
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasPrefix(entry.Name(), "tty.") {
			result = append(result, portPath(filepath.Join("/dev", entry.Name())))
		}
	}
	return result
}
