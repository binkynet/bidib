package wizard

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rs/zerolog"

	"github.com/binkynet/bidib/host"
	serialtx "github.com/binkynet/bidib/transport/serial"
	"github.com/binkynet/bidib/wizard/serial"
)

func NewApp() App {
	return App{
		state:           appStateSelectPort,
		serialSelection: serial.New(),
		nodeTree:        NewNodeTree(nil),
		logView:         NewLogView(),
	}
}

type appState uint8

const (
	appStateSelectPort appState = iota
	appStateNodeTree
	appStateLogView
)

var (
	selectPortStyle = lipgloss.NewStyle().
			Align(lipgloss.Left, lipgloss.Top).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#442222"))
	modelStyle = lipgloss.NewStyle().
			Align(lipgloss.Left, lipgloss.Top).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#3c3c3c"))
	focusedModelStyle = lipgloss.NewStyle().
				Align(lipgloss.Left, lipgloss.Top).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("69"))
)

// Application model
type App struct {
	state           appState
	serialSelection serial.PortSelector
	nodeTree        NodeTree
	logView         LogView

	height, width int
	host          host.Host
}

func (m App) Init() tea.Cmd {
	return tea.Batch(
		m.serialSelection.Init(),
		m.logView.Init(),
	)
}

func (m App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.applyLayout()
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "tab":
			if m.state == appStateNodeTree {
				m.state = appStateLogView
				return m, nil
			} else if m.state == appStateLogView {
				m.state = appStateNodeTree
				return m, nil
			}
		case "l":
			if m.state == appStateNodeTree {
				m.state = appStateLogView
				return m, nil
			}
		case "t":
			if m.state == appStateLogView {
				m.state = appStateNodeTree
				return m, nil
			}
		default:
			switch m.state {
			case appStateSelectPort:
				cmds = append(cmds, m.updateSerialSelection(msg)...)
			case appStateNodeTree:
				cmds = append(cmds, m.updateNodeTree(msg))
			case appStateLogView:
				cmds = append(cmds, m.updateLogView(msg))
			}
		}
	default:
		if m.state == appStateSelectPort {
			cmds = append(cmds, m.updateSerialSelection(msg)...)
		} else {
			cmds = append(cmds, m.updateNodeTree(msg))
			cmds = append(cmds, m.updateLogView(msg))
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *App) updateSerialSelection(msg tea.Msg) []tea.Cmd {
	var cmds []tea.Cmd

	var cmd tea.Cmd
	m.serialSelection, cmd = m.serialSelection.Update(msg)
	cmds = append(cmds, cmd)

	if serialPath := m.serialSelection.SelectedPath(); serialPath != "" {
		// Got serial port
		log := zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.Out = m.logView.Writer()
			w.TimeFormat = time.TimeOnly
		})).With().Timestamp().Logger()
		cfg := host.Config{
			Serial: &serialtx.Config{
				PortName: serialPath,
			},
		}
		h, err := host.New(cfg, log)
		if err != nil {
			log.Fatal().Err(err).Msg("host.New failed")
		}
		m.host = h
		m.nodeTree = NewNodeTree(h)

		m.state = appStateNodeTree
		cmds = append(cmds, m.nodeTree.Init())
		m.applyLayout()
	}
	return cmds
}

func (m *App) updateNodeTree(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.nodeTree, cmd = m.nodeTree.Update(msg)
	return cmd
}

func (m *App) updateLogView(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.logView, cmd = m.logView.Update(msg)
	return cmd
}

// Set the layout of all models
func (m *App) applyLayout() {
	contentWidth := m.width - modelStyle.GetHorizontalFrameSize()
	vfs := modelStyle.GetVerticalFrameSize()
	contentHeight1 := (m.height / 2) - vfs
	contentHeight2 := (m.height - contentHeight1) - (vfs * 2)

	m.serialSelection.SetSize(contentWidth, contentHeight1)
	m.nodeTree.SetSize(contentWidth, contentHeight1)
	m.logView.SetSize(contentWidth, contentHeight2)
}

func (m App) View() string {
	contentWidth := m.width - modelStyle.GetHorizontalFrameSize()
	vfs := modelStyle.GetVerticalFrameSize()
	contentHeight1 := (m.height / 2) - vfs
	contentHeight2 := (m.height - contentHeight1) - (vfs * 2)

	sized := func(s lipgloss.Style, h int) lipgloss.Style {
		return s.Width(contentWidth).Height(h)
	}

	switch m.state {
	case appStateSelectPort:
		return sized(focusedModelStyle, contentHeight1).Render(m.serialSelection.View())
	case appStateNodeTree:
		return lipgloss.JoinVertical(lipgloss.Left,
			sized(focusedModelStyle, contentHeight1).Render(m.nodeTree.View()),
			sized(modelStyle, contentHeight2).Render(m.logView.View()),
		)
	case appStateLogView:
		return lipgloss.JoinVertical(lipgloss.Left,
			sized(modelStyle, contentHeight1).Render(m.nodeTree.View()),
			sized(focusedModelStyle, contentHeight2).Render(m.logView.View()),
		)
	default:
		return ""
	}
}
