package wizard

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/binkynet/bidib/host"
)

func NewNodeInfo(n *host.Node) NodeInfo {
	m := NodeInfo{
		node: n,
		view: viewport.New(0, 0),
	}
	m.styles.Title = lipgloss.NewStyle().
		Background(lipgloss.Color("62")).
		Foreground(lipgloss.Color("230")).
		Padding(0, 1)
	m.styles.Info = lipgloss.NewStyle().
		Padding(1, 0)

	m.reloadInfo()
	return m
}

// Node info model
type NodeInfo struct {
	width, height int
	node          *host.Node
	view          viewport.Model
	styles        struct {
		Title lipgloss.Style
		Info  lipgloss.Style
	}
}

// Reload all information about the node
func (m *NodeInfo) reloadInfo() {
	b := strings.Builder{}
	if m.node != nil {
		if cs := m.node.Cs(); cs != nil {
			b.WriteString(fmt.Sprintf("DCC Generator State: %s\n", cs.GetState()))
		}
	}
	m.view.SetContent(b.String())
}

// SetNode changes the node for which info is shown.
func (m *NodeInfo) SetNode(n *host.Node) {
	m.node = n
	m.reloadInfo()
}

func (m NodeInfo) Init() tea.Cmd {
	return nil
}

func (m NodeInfo) Update(msg tea.Msg) (NodeInfo, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case nodeChangedMsg:
		m.reloadInfo()
		return m, nil
	default:
		cmds = append(cmds, m.updateList(msg))
	}

	return m, tea.Batch(cmds...)
}

func (m *NodeInfo) updateList(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.view, cmd = m.view.Update(msg)
	return cmd
}

func (m NodeInfo) View() string {
	title := "No node selected"
	if m.node != nil {
		title = "Node " + m.node.Address.String()
	}
	titleView := m.styles.Title.Render(title)
	maxHeight := m.height - lipgloss.Height(titleView)
	infoView := m.styles.Info.MaxHeight(maxHeight).Render(m.view.View())
	return titleView + infoView
}

func (m *NodeInfo) SetSize(w, h int) {
	m.width, m.height = w, h
	m.applyLayout()
}

func (m *NodeInfo) applyLayout() {
	m.view.Height = m.height
	m.view.Width = m.width
}
