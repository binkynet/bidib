package wizard

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/host"
)

func NewCVProgrammer(n *host.Node) *CVProgrammer {
	m := &CVProgrammer{
		node: n,
		pomOpts: host.ProgramOnMainOptions{
			DccAddress: 3,
			Cv:         1,
		},
	}

	addrBox := NewNumberInput("Address", int(m.pomOpts.DccAddress), 10239)
	addrBox.MinValue = 1
	addrBox.OnChanged = func(v int) {
		m.pomOpts.DccAddress = uint32(v)
	}
	m.inputs = append(m.inputs, addrBox)

	cvBox := NewNumberInput("CV", int(m.pomOpts.Cv), 1024)
	cvBox.MinValue = 1
	cvBox.OnChanged = func(v int) {
		m.pomOpts.Cv = uint32(v)
	}
	m.inputs = append(m.inputs, cvBox)

	m.valueBox = NewNumberInput("Value", 0, 255)
	m.valueBox.OnChanged = func(v int) {
		m.pomOpts.Data = uint8(v)
	}
	m.inputs = append(m.inputs, m.valueBox)

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
	valueBox      *NumberInput
	focusIndex    int
	pomOpts       host.ProgramOnMainOptions
}

// readCV sends a read command
func (m *CVProgrammer) readCV() {
	if m.node != nil {
		if cs := m.node.Cs(); cs != nil {
			cs.ProgramOnMain(host.ProgramOnMainOptions{
				OpCode:     bidib.BIDIB_CS_POM_RD_BYTE,
				DccAddress: m.pomOpts.DccAddress,
				Cv:         m.pomOpts.Cv,
				Data:       0,
			})
		}
	}
}

// writeCVByte sends a write-byte command
func (m *CVProgrammer) writeCVByte() {
	if m.node != nil {
		if cs := m.node.Cs(); cs != nil {
			cs.ProgramOnMain(host.ProgramOnMainOptions{
				OpCode:     bidib.BIDIB_CS_POM_WR_BYTE,
				DccAddress: m.pomOpts.DccAddress,
				Cv:         m.pomOpts.Cv,
				Data:       m.pomOpts.Data,
			})
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
	case nodeChangedMsg:
		switch msg := msg.Payload.(type) {
		case host.ProgramOnMainOptions:
			if m.pomOpts.DccAddress == msg.DccAddress && m.pomOpts.Cv == msg.Cv {
				m.pomOpts.Data = msg.Data
				m.valueBox.setValue(int(msg.Data))
			}
			return nil
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
