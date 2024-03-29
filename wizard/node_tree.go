package wizard

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/binkynet/bidib/host"
)

func NewNodeTree(h host.Host) NodeTree {
	m := NodeTree{
		state:        nodeTreeStateTree,
		host:         h,
		list:         list.New(nil, list.NewDefaultDelegate(), 0, 0),
		menu:         NewNodeMenu(nil),
		info:         NewNodeInfo(nil),
		driversCab:   NewDriversCab(nil),
		cvProgrammer: NewCVProgrammer(nil),
		nodeChanges:  make(chan nodeChangedMsg, 64),
	}
	m.list.Title = "Nodes"
	m.list.SetShowStatusBar(false)
	m.keyMap.LevelChange = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "change level"),
	)
	m.keyMap.ShowMenu = key.NewBinding(
		key.WithKeys("m"),
		key.WithHelp("m", "show menu"),
	)
	m.keyMap.ShowFeatures = key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "show features"),
	)
	m.keyMap.Back = key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "go back"),
	)
	return m
}

type nodeTreeState uint8

const (
	nodeTreeStateTree nodeTreeState = iota
	nodeTreeStateFeatures
	nodeTreeStateMenu
	nodeTreeStateDriversCab
	nodeTreeStateCVProgrammer
)

// Application model
type NodeTree struct {
	state         nodeTreeState
	width, height int
	host          host.Host
	node          *host.Node
	list          list.Model
	nodeChanges   chan nodeChangedMsg
	featureTable  FeatureTable
	menu          NodeMenu
	info          NodeInfo
	driversCab    *DriversCab
	cvProgrammer  *CVProgrammer
	keyMap        struct {
		LevelChange  key.Binding
		ShowMenu     key.Binding
		ShowFeatures key.Binding
		Back         key.Binding
	}
}

// Reload all the items into the list, based on the current node.
func (m *NodeTree) reloadListItems() {
	var items []list.Item
	if m.node != nil {
		if m.node.Address.HasParent() {
			if parent, ok := m.host.GetNode(m.node.Address.Parent()); ok {
				items = append(items, nodeTreeItem{node: parent, role: "...", icon: iconArrowLeft})
			}
		}
		items = append(items, nodeTreeItem{node: m.node})
		m.node.ForEachChild(func(child *host.Node) {
			items = append(items, nodeTreeItem{node: child, icon: iconArrowRight})
		})
	}
	m.list.SetItems(items)
	m.info.SetNode(m.getSelectedNode())
	m.driversCab.SetNode(m.getSelectedNode())
	m.cvProgrammer.SetNode(m.getSelectedNode())
}

type selectCurrentNodeMsg *host.Node
type nodeChangedMsg host.NodeEvent

func (m NodeTree) Init() tea.Cmd {
	m.host.RegisterNodeChanged(func(n host.NodeEvent) {
		m.nodeChanges <- nodeChangedMsg(n)
	})
	return tea.Batch(
		func() tea.Msg {
			return selectCurrentNodeMsg(m.host.GetRootNode())
		},
		m.onNodeChanged(),
	)
}

func (m NodeTree) onNodeChanged() tea.Cmd {
	return func() tea.Msg {
		return <-m.nodeChanges
	}
}

func (m NodeTree) Update(msg tea.Msg) (NodeTree, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.LevelChange) && m.state == nodeTreeStateTree:
			if selectedNode := m.getSelectedNode(); selectedNode != nil {
				if m.node != selectedNode {
					m.node = selectedNode
					m.reloadListItems()
				}
				return m, nil
			}
		case key.Matches(msg, m.keyMap.ShowFeatures) && m.state == nodeTreeStateTree:
			return m.showFeatures()
		case key.Matches(msg, m.keyMap.ShowMenu) && m.state == nodeTreeStateTree:
			if selectedNode := m.getSelectedNode(); selectedNode != nil {
				m.menu = NewNodeMenu(selectedNode)
				m.state = nodeTreeStateMenu
				m.applyLayout()
				return m, m.menu.Init()
			}
		case key.Matches(msg, m.keyMap.Back) && m.state != nodeTreeStateTree:
			m.state = nodeTreeStateTree
			return m, nil
		default:
			switch m.state {
			case nodeTreeStateTree:
				cmds = append(cmds, m.updateList(msg))
			case nodeTreeStateFeatures:
				cmds = append(cmds, m.updateFeatureTable(msg))
			case nodeTreeStateMenu:
				cmds = append(cmds, m.updateMenu(msg))
			case nodeTreeStateDriversCab:
				cmds = append(cmds, m.updateDriversCab(msg))
			case nodeTreeStateCVProgrammer:
				cmds = append(cmds, m.updateCVProgrammer(msg))
			}
		}
		return m, tea.Batch(cmds...)
	case nodeMenuItemReset:
		m.getSelectedNode().Reset()
		m.state = nodeTreeStateTree
		return m, nil
	case nodeMenuItemShowFeatures:
		return m.showFeatures()
	case nodeMenuItemShowDriversCab:
		return m.showDriversCab()
	case nodeMenuItemShowCVProgrammer:
		return m.showCVProgrammer()
	case nodeMenuItemCsOff:
		m.getSelectedNode().Cs().Off()
		m.state = nodeTreeStateTree
		return m, nil
	case nodeMenuItemCsGo:
		m.getSelectedNode().Cs().Go()
		m.state = nodeTreeStateTree
		return m, nil
	case nodeMenuItemCsStop:
		m.getSelectedNode().Cs().Stop()
		m.state = nodeTreeStateTree
		return m, nil
	case selectCurrentNodeMsg:
		m.node = msg
		m.reloadListItems()
	case nodeChangedMsg:
		m.reloadListItems()
		m.info.reloadInfo()
		if msg.Payload != nil {
			cmds = append(cmds, m.updateCVProgrammer(msg))
		}
		cmds = append(cmds, m.onNodeChanged())
	default:
		cmds = append(cmds, m.updateList(msg))
		cmds = append(cmds, m.updateFeatureTable(msg))
		cmds = append(cmds, m.updateMenu(msg))
		cmds = append(cmds, m.updateInfo(msg))
		cmds = append(cmds, m.updateDriversCab(msg))
		cmds = append(cmds, m.updateCVProgrammer(msg))
	}

	return m, tea.Batch(cmds...)
}

func (m *NodeTree) updateList(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	m.info.SetNode(m.getSelectedNode())
	return cmd
}

func (m *NodeTree) updateFeatureTable(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.featureTable, cmd = m.featureTable.Update(msg)
	return cmd
}

func (m *NodeTree) updateMenu(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.menu, cmd = m.menu.Update(msg)
	return cmd
}

func (m *NodeTree) updateInfo(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.info, cmd = m.info.Update(msg)
	return cmd
}

func (m *NodeTree) updateDriversCab(msg tea.Msg) tea.Cmd {
	return m.driversCab.Update(msg)
}

func (m *NodeTree) updateCVProgrammer(msg tea.Msg) tea.Cmd {
	return m.cvProgrammer.Update(msg)
}

func (m NodeTree) showFeatures() (NodeTree, tea.Cmd) {
	if selectedNode := m.getSelectedNode(); selectedNode != nil {
		m.featureTable = NewFeatureTable(selectedNode)
		m.state = nodeTreeStateFeatures
		m.applyLayout()
		return m, m.featureTable.Init()
	}
	return m, nil
}

func (m NodeTree) showDriversCab() (NodeTree, tea.Cmd) {
	if selectedNode := m.getSelectedNode(); selectedNode != nil {
		m.driversCab.SetNode(selectedNode)
		m.state = nodeTreeStateDriversCab
		m.applyLayout()
		return m, m.driversCab.Init()
	}
	return m, nil
}

func (m NodeTree) showCVProgrammer() (NodeTree, tea.Cmd) {
	if selectedNode := m.getSelectedNode(); selectedNode != nil {
		m.cvProgrammer.SetNode(selectedNode)
		m.state = nodeTreeStateCVProgrammer
		m.applyLayout()
		return m, m.cvProgrammer.Init()
	}
	return m, nil
}

func (m NodeTree) View() string {
	switch m.state {
	case nodeTreeStateTree:
		s := lipgloss.NewStyle().Width(m.list.Width())
		return lipgloss.JoinHorizontal(lipgloss.Top,
			s.Render(m.list.View()),
			m.info.View(),
		)
	case nodeTreeStateFeatures:
		return m.featureTable.View()
	case nodeTreeStateMenu:
		return m.menu.View()
	case nodeTreeStateDriversCab:
		return m.driversCab.View()
	case nodeTreeStateCVProgrammer:
		return m.cvProgrammer.View()
	}
	return ""
}

func (m *NodeTree) SetSize(w, h int) {
	m.width, m.height = w, h
	m.applyLayout()
}

func (m *NodeTree) applyLayout() {
	listWidth := m.width / 2
	m.list.SetSize(listWidth, m.height)
	m.info.SetSize(m.width-listWidth, m.height)
	m.featureTable.SetSize(m.width, m.height)
	m.menu.SetSize(m.width, m.height)
	m.driversCab.SetSize(m.width, m.height)
}

// Returns the currently selected node (if any)
func (m *NodeTree) getSelectedNode() *host.Node {
	if item := m.list.SelectedItem(); item != nil {
		return item.(nodeTreeItem).node
	}
	return nil
}
