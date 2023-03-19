package bidib

import "encoding/binary"

// Each node has a distinct identifer, this number is called Unique-ID. The Unique-ID contains 7 bytes:
type UniqueID [7]byte

// This is a bit field indicating the class membership of this node. A node may also belong to several classes at once.
// The classes serve as a quick reference for the host about which functionalities can be found on this specific node.
// To rapidly locate a feature-based subfunctionality it is enough to only query the nodes which have set the corresponding class bit.
// If a node has implemented commands of a particular class, the appropriate class bit must be set as well.
// Conversely, it must know the commands of the announced classes and answer them correctly.
// Even if in the current configuration no objects are available, it should register the class and yield 0 for the count.
func (uid UniqueID) ClassID() uint8 {
	return uid[0]
}

// ClassID Extension; this byte is reserved and must be coded with 0.
func (uid UniqueID) ClassIDExtension() uint8 {
	return uid[1]
}

// Vendor-ID: The same coding as DCC is used here, see NMRA Manufacturer ID Numbers.
func (uid UniqueID) VendorID() uint8 {
	return uid[2]
}

// Product ID, comprising of 32 Bit.
// These 4 bytes (= 32 bits) are split into a product identifier (lower 'p' bits) and a serial
// number (upper 's' bits). This allows for an easier identification of nodes by analysis tools and
// host programs. The coding in BiDiB is always little-endian (low-byte first), the product identifier
// starts at bit 0, first byte. It is up to the vendor, how many bits he uses for product and for serial number.
// However, a default of 16 bit / 16 bit is recommended. The feature FEATURE_RELEVANT_PID_BITS defines,
// how many bits are used for product ('p').
// If the feature FEATURE_RELEVANT_PID_BITS doesn't exist, the default of 16 bits / 16 bits is used.
func (uid UniqueID) ProductID() uint32 {
	return binary.LittleEndian.Uint32(uid[3:])
}
