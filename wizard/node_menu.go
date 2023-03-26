package wizard

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/binkynet/bidib/host"
)

func NewNodeMenu(n *host.Node) NodeMenu {
	d := list.NewDefaultDelegate()
	d.ShowDescription = false
	d.SetHeight(1)
	d.SetSpacing(0)
	m := NodeMenu{
		node: n,
		list: list.New(nil, d, 0, 0),
	}
	m.list.Title = "Menu for "
	m.list.SetShowStatusBar(false)
	m.keyMap.Select = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Select menu item"),
	)
	m.reloadListItems()
	return m
}

// Node menu model
type NodeMenu struct {
	width, height int
	node          *host.Node
	list          list.Model
	keyMap        struct {
		Select key.Binding
	}
}

type nodeMenuItem string

func (i nodeMenuItem) Title() string       { return string(i) }
func (i nodeMenuItem) Description() string { return "" }
func (i nodeMenuItem) FilterValue() string { return i.Title() }

type (
	nodeMenuItemReset        struct{ nodeMenuItem }
	nodeMenuItemShowFeatures struct{ nodeMenuItem }
	nodeMenuItemCsOff        struct{ nodeMenuItem }
	nodeMenuItemCsGo         struct{ nodeMenuItem }
	nodeMenuItemCsStop       struct{ nodeMenuItem }

	nodeMenuItemCsLightsOn3  struct{ nodeMenuItem }
	nodeMenuItemCsLightsOff3 struct{ nodeMenuItem }
)

// Reload all the items into the list, based on the current node.
func (m *NodeMenu) reloadListItems() {
	var items []list.Item
	if m.node != nil {
		items = append(items, nodeMenuItemShowFeatures{"Show features"})
		if m.node.Cs() != nil {
			items = append(items,
				nodeMenuItemCsOff{"DCC Generator Off"},
				nodeMenuItemCsGo{"DCC Generator Go"},
				nodeMenuItemCsStop{"DCC Generator Stop"},
				nodeMenuItemCsLightsOn3{"Light on @for address 3"},
				nodeMenuItemCsLightsOff3{"Light off @for address 3"},
			)
		}
		if !m.node.Address.HasParent() {
			items = append(items, nodeMenuItemReset{"Reset"})
		}
	}
	m.list.SetItems(items)
}

func (m NodeMenu) Init() tea.Cmd {
	return nil
}

func (m NodeMenu) Update(msg tea.Msg) (NodeMenu, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Select):
			if item := m.list.SelectedItem(); item != nil {
				return m, func() tea.Msg { return item }
			}
		default:
			cmds = append(cmds, m.updateList(msg))
		}
		return m, tea.Batch(cmds...)
	default:
		cmds = append(cmds, m.updateList(msg))
	}

	return m, tea.Batch(cmds...)
}

func (m *NodeMenu) updateList(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return cmd
}

func (m NodeMenu) View() string {
	return m.list.View()
}

func (m *NodeMenu) SetSize(w, h int) {
	m.width, m.height = w, h
	m.applyLayout()
}

func (m *NodeMenu) applyLayout() {
	m.list.SetSize(m.width, m.height)
}
