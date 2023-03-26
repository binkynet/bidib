package wizard

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/host"
)

func NewFeatureTable(n *host.Node) FeatureTable {
	m := FeatureTable{
		node: n,
		table: table.New(
			table.WithColumns([]table.Column{
				{"Feature", 40},
				{"Value", 10},
			}),
			table.WithFocused(true),
		),
	}
	m.reloadTableRows()
	return m
}

// Application model
type FeatureTable struct {
	node  *host.Node
	table table.Model
}

// Reload all the rows in the table.
func (m *FeatureTable) reloadTableRows() {
	var items []table.Row
	for id := bidib.FeatureID(0); id < 255; id++ {
		if value, found := m.node.GetFeature(id); found {
			items = append(items, table.Row{id.String(), strconv.Itoa(int(value))})
		}
	}
	m.table.SetRows(items)
}

func (m FeatureTable) Init() tea.Cmd {
	return nil
}

func (m FeatureTable) Update(msg tea.Msg) (FeatureTable, tea.Cmd) {
	var cmds []tea.Cmd

	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m FeatureTable) View() string {
	return m.table.View()
}

// Set the widget size
func (m *FeatureTable) SetSize(w, h int) {
	m.table.SetWidth(w)
	m.table.SetHeight(h)
}
