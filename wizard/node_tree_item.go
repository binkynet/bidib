package wizard

import (
	"github.com/binkynet/bidib/host"
)

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
