package bidib

// Booster State flags
//
//go:generate stringer -type=BstState
type BstState uint8

// Is the booster power on?
func (s BstState) IsOn() bool {
	return s&BIDIB_BST_STATE_ON == BIDIB_BST_STATE_ON
}

// Is the booster power off due to output shortend
func (s BstState) IsShort() bool {
	return s == BIDIB_BST_STATE_OFF_SHORT
}

// Is the booster hot?
func (s BstState) IsHot() bool {
	return s == BIDIB_BST_STATE_OFF_HOT || s == BIDIB_BST_STATE_ON_HOT
}
