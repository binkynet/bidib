package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// Followed by 7 bytes of the previously read UNIQUE-ID. The node responds with a MSG_VENDOR_ACK.
type VendorEnable struct {
	BaseMessage
	UniqueID bidib.UniqueID
}

func (m VendorEnable) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_VENDOR_ENABLE, m.Address, seqNum, m.UniqueID[:])
}

func (m VendorEnable) String() string {
	return fmt.Sprintf("%T addr=%s uid=%s", m, m.Address, m.UniqueID)
}

func decodeVendorEnable(addr bidib.Address, data []byte) (VendorEnable, error) {
	var result VendorEnable
	if err := validateDataLength(data, 7); err != nil {
		return result, err
	}
	result.Address = addr
	copy(result.UniqueID[:], data)
	return result, nil
}

// No other data will follow; the node is disabled. The node responds with a MSG_VENDOR_ACK.
type VendorDisable struct {
	BaseMessage
}

func (m VendorDisable) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_VENDOR_DISABLE, m.Address, seqNum, nil)
}

func (m VendorDisable) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeVendorDisable(addr bidib.Address, data []byte) (VendorDisable, error) {
	var result VendorDisable
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// Followed by the data below, which are structured as follows:
// VENDOR_DATA ::= V_NAME  V_VALUE
// V_NAME ::= LENGTH  V_NAME_STR
// V_NAME_STR ::= V_NAME_CHAR | V_NAME_CHAR  V_NAME_STR
// V_VALUE ::= LENGTH   V_VALUE_STR
// V_VALUE_STR ::= ε | V_VALUE_CHAR  V_VALUE_STR
// The node responds with a MSG_VENDOR message, this includes V_NAME and V_VALUE as well.
type VendorSet struct {
	BaseMessage
	Name  string
	Value string
}

func (m VendorSet) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	lName := byte(len(m.Name))
	lValue := byte(len(m.Value))
	data := make([]byte, 2+lName+lValue)
	data[0] = lName
	copy(data[1:], []byte(m.Name))
	data[1+lName] = lValue
	copy(data[2+lName:], []byte(m.Value))
	bidib.EncodeMessage(write, bidib.MSG_VENDOR_SET, m.Address, seqNum, data)
}

func (m VendorSet) String() string {
	return fmt.Sprintf("%T addr=%s name=%s value=%s", m, m.Address, m.Name, m.Value)
}

func decodeVendorSet(addr bidib.Address, data []byte) (VendorSet, error) {
	var result VendorSet
	if err := validateMinDataLength(data, 4); err != nil {
		return result, err
	}
	result.Address = addr
	lName := data[0]
	data = data[1:]
	result.Name = string(data[:lName])
	data = data[lName+1:]
	result.Value = string(data)
	return result, nil
}

// Followed by the data below, which are structured as follows:
// V_NAME ::= LENGTH  V_NAME_STR
// V_NAME_STR ::= V_NAME_CHAR | V_NAME_CHAR  V_NAME_STR
// The node responds with a MSG_VENDOR message.
type VendorGet struct {
	BaseMessage
	Name string
}

func (m VendorGet) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	lName := byte(len(m.Name))
	data := make([]byte, 1+lName)
	data[0] = lName
	copy(data[1:], []byte(m.Name))
	bidib.EncodeMessage(write, bidib.MSG_VENDOR_GET, m.Address, seqNum, data)
}

func (m VendorGet) String() string {
	return fmt.Sprintf("%T addr=%s name=%s", m, m.Address, m.Name)
}

func decodeVendorGet(addr bidib.Address, data []byte) (VendorGet, error) {
	var result VendorGet
	if err := validateMinDataLength(data, 2); err != nil {
		return result, err
	}
	result.Address = addr
	lName := data[0]
	data = data[1:]
	result.Name = string(data[:lName])
	return result, nil
}

// This message type is used to send a string value to a node.
// This command is followed by data denoting the addressed namespace, the addressed
// identifier (variable to be set, channel), the string size and the string itself.
// This function is only available in a node if announced by the respective feature.
type StringSet struct {
	BaseMessage
	Namespace uint8
	StringID  uint8
	Value     string
}

func (m StringSet) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	lValue := byte(len(m.Value))
	data := make([]byte, 3+lValue)
	data[0] = m.Namespace
	data[1] = m.StringID
	data[2] = lValue
	copy(data[3:], []byte(m.Value))
	bidib.EncodeMessage(write, bidib.MSG_STRING_SET, m.Address, seqNum, data)
}

func (m StringSet) String() string {
	return fmt.Sprintf("%T addr=%s namespace=%d string=%d value=%s", m, m.Address, m.Namespace, m.StringID, m.Value)
}

func decodeStringSet(addr bidib.Address, data []byte) (StringSet, error) {
	var result StringSet
	if err := validateMinDataLength(data, 3); err != nil {
		return result, err
	}
	result.Address = addr
	result.Namespace = data[0]
	result.StringID = data[1]
	lValue := data[2]
	data = data[3:]
	result.Value = string(data[:lValue])
	return result, nil
}

// Query of a string variable inside a node. Two bytes are following, denoting namespace and string id.
// The node answer with a MSG_STRING. This function is only available in specific node,
// if FEATURE_STRING_SIZE exists and has a value > 0.
// If a string doesn't exist, the returned SIZE is 0.
type StringGet struct {
	BaseMessage
	Namespace uint8
	StringID  uint8
}

func (m StringGet) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{m.Namespace, m.StringID}
	bidib.EncodeMessage(write, bidib.MSG_STRING_GET, m.Address, seqNum, data)
}

func (m StringGet) String() string {
	return fmt.Sprintf("%T addr=%s namespace=%d string=%d", m, m.Address, m.Namespace, m.StringID)
}

func decodeStringGet(addr bidib.Address, data []byte) (StringGet, error) {
	var result StringGet
	if err := validateDataLength(data, 2); err != nil {
		return result, err
	}
	result.Address = addr
	result.Namespace = data[0]
	result.StringID = data[1]
	return result, nil
}
