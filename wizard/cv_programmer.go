package wizard

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/host"
)

func NewCVProgrammer(n *host.Node) *CVProgrammer {
	m := &CVProgrammer{
		node:       n,
		dccAddress: 3,
		cv:         1,
	}

	addrBox := NewNumberInput("Address", int(m.dccAddress), 10239)
	addrBox.MinValue = 1
	addrBox.OnChanged = func(v int) {
		m.dccAddress = uint16(v)
	}
	m.inputs = append(m.inputs, addrBox)

	cvBox := NewNumberInput("CV", int(m.cv), 1024)
	cvBox.MinValue = 1
	cvBox.OnChanged = func(v int) {
		m.cv = uint16(v)
	}
	m.inputs = append(m.inputs, cvBox)

	valueBox := NewNumberInput("Value", 0, 255)
	valueBox.OnChanged = func(v int) {
		m.value = uint8(v)
	}
	m.inputs = append(m.inputs, valueBox)

	readButton := NewButton("Read", func() {
		m.readCV()
	})
	m.inputs = append(m.inputs, readButton)

	writeButton := NewButton("Write!", func() {
		m.writeCVByte()
	})
	m.inputs = append(m.inputs, writeButton)

	m.updateFocus()

	return m
}

// DriversCab model
type CVProgrammer struct {
	width, height int
	node          *host.Node
	inputs        []inputModel
	focusIndex    int
	dccAddress    uint16
	cv            uint16
	value         uint8
	busy          bool
}

// readCV sends a read command
func (m *CVProgrammer) readCV() {
	if m.node != nil {
		if cs := m.node.Cs(); cs != nil {
			cs.ProgramOnMain(bidib.BIDIB_CS_POM_RD_BYTE, uint32(m.dccAddress), uint32(m.cv), 0)
		}
	}
}

// writeCVByte sends a write-byte command
func (m *CVProgrammer) writeCVByte() {
	if m.node != nil {
		if cs := m.node.Cs(); cs != nil {
			cs.ProgramOnMain(bidib.BIDIB_CS_POM_WR_BYTE, uint32(m.dccAddress), uint32(m.cv), m.value)
		}
	}
}

// SetNode changes the node for which info is shown.
func (m *CVProgrammer) SetNode(n *host.Node) {
	m.node = n
}

func (m *CVProgrammer) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0, len(m.inputs))
	for _, im := range m.inputs {
		cmds = append(cmds, im.Init())
	}
	return tea.Batch(cmds...)
}

func (m *CVProgrammer) Update(msg tea.Msg) tea.Cmd {
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

func (m *CVProgrammer) View() string {
	inputViews := make([]string, 0, len(m.inputs))
	for _, im := range m.inputs {
		inputViews = append(inputViews, im.View()+" ")
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, inputViews...)
}

func (m *CVProgrammer) SetSize(w, h int) {
	m.width, m.height = w, h
	m.applyLayout()
}

func (m *CVProgrammer) applyLayout() {
	/*m.view.Height = m.height
	m.view.Width = m.width*/
}

func (m *CVProgrammer) updateFocus() []tea.Cmd {
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
