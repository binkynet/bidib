package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// Followed by 1 byte with the feature number and 1 byte with the value.
// Logical features are enabled at 1 and disabled at 0.
type Feature struct {
	BaseMessage
	Feature bidib.FeatureID
	Value   uint8
}

func (m Feature) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{byte(m.Feature), m.Value}
	bidib.EncodeMessage(write, bidib.MSG_FEATURE, m.Address, seqNum, data)
}

func (m Feature) String() string {
	return fmt.Sprintf("%T addr=%s feature=0x%02x value=0x%02x", m, m.Address, m.Feature, m.Value)
}

func decodeFeature(addr bidib.Address, data []byte) (Feature, error) {
	var result Feature
	if err := validateDataLength(data, 2); err != nil {
		return result, err
	}
	result.Address = addr
	result.Feature = bidib.FeatureID(data[0])
	result.Value = data[1]
	return result, nil
}

// This message will be sent if a feature was requested that is not available on this node.
// Followed by 1 byte with the (not implemented) feature number.
// This message is also sent (with feature number 255) in response to MSG_FEATURE_GETNEXT,
// if all features have already been transmitted.
type FeatureNa struct {
	BaseMessage
	Feature bidib.FeatureID
}

func (m FeatureNa) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{byte(m.Feature)}
	bidib.EncodeMessage(write, bidib.MSG_FEATURE_NA, m.Address, seqNum, data)
}

func (m FeatureNa) String() string {
	return fmt.Sprintf("%T addr=%s feature=0x%02x", m, m.Address, m.Feature)
}

func decodeFeatureNa(addr bidib.Address, data []byte) (FeatureNa, error) {
	var result FeatureNa
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.Feature = bidib.FeatureID(data[0])
	return result, nil
}

// This message is sent prior to the transmission if the host has made a request with MSG_FEATURE_GETALL.
// Followed by 1 byte with the number of the existing feature messages and optionally 1 byte
// for announcing the transmission mode.
// The mode byte is set to the value 1 when the host requested streaming and the node supports it.
// The node begins sending the MSG_FEATURE on its own and is responsible for flow control,
// adjusting to the available transport capacity itself. The node must remain fully operable
// and be able to receive and respond to other messages during the transmission.
// The count allows the host to determine when all feature messages have arrived.
// Otherwise, the feature values are polled individually using a sequence of MSG_FEATURE_GETNEXT.
// The count allows the host to make the suitable number of requests.
type FeatureCount struct {
	BaseMessage
	Count     uint8
	Streaming bool
}

func (m FeatureCount) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{m.Count, 0}
	if m.Streaming {
		data[1] = 1
	}
	bidib.EncodeMessage(write, bidib.MSG_FEATURE_COUNT, m.Address, seqNum, data)
}

func (m FeatureCount) String() string {
	return fmt.Sprintf("%T addr=%s count=%d streaming=%t", m, m.Address, m.Count, m.Streaming)
}

func decodeFeatureCount(addr bidib.Address, data []byte) (FeatureCount, error) {
	var result FeatureCount
	if err := validateMinDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.Count = data[0]
	if len(data) > 1 {
		result.Streaming = data[1] != 0
	}
	return result, nil
}
