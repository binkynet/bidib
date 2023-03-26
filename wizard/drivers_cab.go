package wizard

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/host"
)

func NewDriversCab(n *host.Node) *DriversCab {
	m := &DriversCab{
		node: n,
	}
	m.driveOpts.DccAddress = 3
	m.driveOpts.DccFormat = bidib.BIDIB_CS_DRIVE_FORMAT_DCC128
	m.driveOpts.Speed = 0
	m.driveOpts.OutputSpeed = true
	m.driveOpts.OutputF1_F4 = true
	m.driveOpts.OutputF5_F8 = true
	m.driveOpts.Flags = make(bidib.DccFlags, 9)
	for i := 0; i < len(m.driveOpts.Flags); i++ {
		var title string
		if i == 0 {
			title = "FL"
		} else {
			title = fmt.Sprintf("F%d", i)
		}
		cb := NewCheckBox(title, false)
		idx := i // bring i into scope
		cb.OnChanged = func(b bool) {
			m.driveOpts.Flags.Set(idx, b)
			m.drive()
		}
		m.inputs = append(m.inputs, cb)
	}

	cbDir := NewCheckBox("Forward", true)
	cbDir.OnChanged = func(b bool) {
		m.driveOpts.DirectionForward = b
		m.drive()
	}
	m.inputs = append(m.inputs, cbDir)

	speedBox := NewSpeedBox("Speed", 0, 128)
	speedBox.OnChanged = func(v uint8) {
		m.driveOpts.Speed = v
		m.drive()
	}
	m.inputs = append(m.inputs, speedBox)
	m.updateFocus()

	return m
}

type inputModel interface {
	Focus() tea.Cmd
	Blur() tea.Cmd
	View() string
	Init() tea.Cmd
	Update(tea.Msg) tea.Cmd
}

// DriversCab model
type DriversCab struct {
	width, height int
	node          *host.Node
	inputs        []inputModel
	focusIndex    int
	driveOpts     host.DriveOptions
}

// drive sends the drive options
func (m *DriversCab) drive() {
	if m.node != nil {
		if cs := m.node.Cs(); cs != nil {
			cs.Drive(m.driveOpts)
		}
	}
}

// SetNode changes the node for which info is shown.
func (m *DriversCab) SetNode(n *host.Node) {
	m.node = n
}

func (m *DriversCab) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0, len(m.inputs))
	for _, im := range m.inputs {
		cmds = append(cmds, im.Init())
	}
	return tea.Batch(cmds...)
}

func (m *DriversCab) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			m.focusIndex--
			if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) - 1
			}
			cmds = append(cmds, m.updateFocus()...)
		case "right":
			m.focusIndex++
			if m.focusIndex >= len(m.inputs) {
				m.focusIndex = 0
			}
			cmds = append(cmds, m.updateFocus()...)
		}
	}

	for _, im := range m.inputs {
		cmds = append(cmds, im.Update(msg))
	}

	return tea.Batch(cmds...)
}

func (m *DriversCab) View() string {
	inputViews := make([]string, 0, len(m.inputs))
	for _, im := range m.inputs {
		inputViews = append(inputViews, im.View()+" ")
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, inputViews...)
}

func (m *DriversCab) SetSize(w, h int) {
	m.width, m.height = w, h
	m.applyLayout()
}

func (m *DriversCab) applyLayout() {
	/*m.view.Height = m.height
	m.view.Width = m.width*/
}

func (m *DriversCab) updateFocus() []tea.Cmd {
	var cmds []tea.Cmd
	for idx, im := range m.inputs {
		if idx == m.focusIndex {
			cmds = append(cmds, im.Focus())
		} else {
			cmds = append(cmds, im.Blur())
		}
	}
	return cmds
}
