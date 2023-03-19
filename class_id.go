package bidib

import "strings"

// This is a bit field indicating the class membership of this node. A node may also belong to several classes at once.
// The classes serve as a quick reference for the host about which functionalities can be found on this specific node.
// To rapidly locate a feature-based subfunctionality it is enough to only query the nodes which have set the corresponding class bit.
// If a node has implemented commands of a particular class, the appropriate class bit must be set as well.
// Conversely, it must know the commands of the announced classes and answer them correctly.
// Even if in the current configuration no objects are available, it should register the class and yield 0 for the count.
type ClassID uint8

// Bit 7	1: Node contains sub-nodes (is an interface itself)
func (cid ClassID) HasSubNodes() bool {
	return (cid & (1 << 7)) != 0
}

// Bit 6	1: Node contains occupancy detection functions
func (cid ClassID) HasOccupancyDetectionFunctions() bool {
	return (cid & (1 << 6)) != 0
}

// Bit 5	reserved (1: Node contains operating functions, HMI)

// Bit 4	1: Node contains DCC signal generator for driving, switching
func (cid ClassID) HasDCCSignalGenerator() bool {
	return (cid & (1 << 4)) != 0
}

// Bit 3	1: Node contains DCC signal generator for programming
func (cid ClassID) HasDCCSignalGeneratorForProgramming() bool {
	return (cid & (1 << 3)) != 0
}

// Bit 2	1: Node contains accessory control functions
func (cid ClassID) HasAccessoryControlFunctions() bool {
	return (cid & (1 << 2)) != 0
}

// Bit 1	1: Node contains booster functions
func (cid ClassID) HasBoosterFunctions() bool {
	return (cid & (1 << 1)) != 0
}

// Bit 0	1: Node contains switching functions, e.g. light animation
func (cid ClassID) HasSwitchingFunctions() bool {
	return (cid & (1)) != 0
}

// String convers to a human readable string
func (cid ClassID) String() string {
	result := make([]string, 0, 8)
	if cid.HasSubNodes() {
		result = append(result, "subNodes")
	}
	if cid.HasOccupancyDetectionFunctions() {
		result = append(result, "occupancy")
	}
	if cid.HasDCCSignalGenerator() {
		result = append(result, "dcc")
	}
	if cid.HasDCCSignalGeneratorForProgramming() {
		result = append(result, "dcc-prog")
	}
	if cid.HasAccessoryControlFunctions() {
		result = append(result, "accessory")
	}
	if cid.HasBoosterFunctions() {
		result = append(result, "booster")
	}
	if cid.HasSwitchingFunctions() {
		result = append(result, "switch")
	}
	return strings.Join(result, ",")
}
