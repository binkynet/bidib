package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// This message type is used for the answer to a userconfig.
// Followed by the data below, which are structured as follows:
// VENDOR_DATA ::= V_NAME  V_VALUE
// V_NAME ::= LENGTH  V_NAME_STR
// V_NAME_STR ::= V_NAME_CHAR | V_NAME_CHAR  V_NAME_STR
// V_VALUE ::= LENGTH  V_VALUE_STR
// V_VALUE_STR ::= ε | V_VALUE_CHAR  V_VALUE_STR
type Vendor struct {
	BaseMessage
	Name  string
	Value string
}

func (m Vendor) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	lName := byte(len(m.Name))
	lValue := byte(len(m.Value))
	data := make([]byte, 2+lName+lValue)
	data[0] = lName
	copy(data[1:], []byte(m.Name))
	data[1+lName] = lValue
	copy(data[2+lName:], []byte(m.Value))
	bidib.EncodeMessage(write, bidib.MSG_VENDOR, m.Address, seqNum, data)
}

func (m Vendor) String() string {
	return fmt.Sprintf("%T addr=%s name=%s value=%s", m, m.Address, m.Name, m.Value)
}

func decodeVendor(addr bidib.Address, data []byte) (Vendor, error) {
	var result Vendor
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

// Followed by a date: 0: No user config-mode, 1: Confirmation that the node has changed into userconfig mode.
type VendorAck struct {
	BaseMessage
	Changed bool
}

func (m VendorAck) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{0}
	if m.Changed {
		data[0] = 1
	}
	bidib.EncodeMessage(write, bidib.MSG_VENDOR_ACK, m.Address, seqNum, data)
}

func (m VendorAck) String() string {
	return fmt.Sprintf("%T addr=%s changed=%t", m, m.Address, m.Changed)
}

func decodeVendorAck(addr bidib.Address, data []byte) (VendorAck, error) {
	var result VendorAck
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.Changed = data[0] != 0
	return result, nil
}

// This message type is used as the answer to MSG_STRING_SET or MSG_STRING_GET.
// The message may also be sent spontaneously for namespace 1 once the node is enabled.
// Followed by data denoting the used namespace, the string_id, the string
// size and the string itself. For details see MSG_STRING_SET.
type String struct {
	BaseMessage
	Namespace uint8
	StringID  uint8
	Value     string
}

func (m String) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	lValue := byte(len(m.Value))
	data := make([]byte, 3+lValue)
	data[0] = m.Namespace
	data[1] = m.StringID
	data[2] = lValue
	copy(data[3:], []byte(m.Value))
	bidib.EncodeMessage(write, bidib.MSG_STRING, m.Address, seqNum, data)
}

func (m String) String() string {
	return fmt.Sprintf("%T addr=%s namespace=%d string=%d value=%s", m, m.Address, m.Namespace, m.StringID, m.Value)
}

func decodeString(addr bidib.Address, data []byte) (String, error) {
	var result String
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
