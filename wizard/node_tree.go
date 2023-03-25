package wizard

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/binkynet/bidib/host"
)

const (
	iconArrowLeft  = "\u2190"
	iconArrowRight = "\u2192"
)

func NewNodeTree(h host.Host) NodeTree {
	m := NodeTree{
		host:        h,
		list:        list.New(nil, list.NewDefaultDelegate(), 0, 0),
		nodeChanges: make(chan nodeChangedMsg, 64),
	}
	m.list.Title = "Nodes"
	m.list.SetShowStatusBar(false)
	return m
}

// Application model
type NodeTree struct {
	host        host.Host
	node        *host.Node
	list        list.Model
	nodeChanges chan nodeChangedMsg
}

type nodeTreeItem struct {
	icon string
	role string
	node *host.Node
}

func (i nodeTreeItem) Title() string {
	prefix := i.icon
	if len(prefix) > 0 {
		prefix += " "
	}
	if i.node.Address.GetLength() == 0 {
		return prefix + "<interface>"
	}
	return prefix + i.node.Address.String()
}
func (i nodeTreeItem) Description() string {
	if len(i.role) != 0 {
		return i.role
	}
	return i.node.UniqueID.String()
}
func (i nodeTreeItem) FilterValue() string { return i.Title() }

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
}

type selectCurrentNodeMsg *host.Node
type nodeChangedMsg *host.Node

func (m NodeTree) Init() tea.Cmd {
	m.host.RegisterNodeChanged(func(n *host.Node) {
		m.nodeChanges <- n
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
		switch msg.String() {
		case "enter":
			if item := m.list.SelectedItem(); item != nil {
				m.node = item.(nodeTreeItem).node
				m.reloadListItems()
				return m, nil
			}
		}
	case selectCurrentNodeMsg:
		m.node = msg
		m.reloadListItems()
	case nodeChangedMsg:
		m.reloadListItems()
		cmds = append(cmds, m.onNodeChanged())
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m NodeTree) View() string {
	return m.list.View()
}

func (m *NodeTree) SetSize(w, h int) {
	m.list.SetSize(w, h)
}
