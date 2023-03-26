package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// This command is used to begin the query of all feature settings. Followed by an optional byte
// for requesting a streaming transmission.
// The node resets its internal counter for MSG_FEATURE_GETNEXT queries and responds with a MSG_FEATURE_COUNT,
// which specifies the number of existing features. If this number is 0, the node has no features.
// If the optional parameter is set to value 1, this signals to the node that it should begin sending
// the feature messages without waiting for MSG_FEATURE_GETNEXT queries.
// Supporting this functionality is optional, but recommended for nodes from declared protocol version 0.8.
type FeatureGetAll struct {
	BaseMessage
	Streaming bool
}

func (m FeatureGetAll) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	var data []byte
	if m.Streaming {
		data = []byte{1}
	}
	bidib.EncodeMessage(write, bidib.MSG_FEATURE_GETALL, m.Address, seqNum, data)
}

func (m FeatureGetAll) String() string {
	return fmt.Sprintf("%T addr=%s streaming=%t", m, m.Address, m.Streaming)
}

func decodeFeatureGetAll(addr bidib.Address, data []byte) (FeatureGetAll, error) {
	var result FeatureGetAll
	if err := validateMinDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	if len(data) > 0 {
		result.Streaming = data[0] != 0
	}
	return result, nil
}

// With this message, a feature value is queried. No byte will follow.
// The answer is either a MSG_FEATURE (the node itself selects and sends the respective next FEATURE)
// or a MSG_FEATURE_NA message (with feature_num = 255), if all features have been already submitted.
type FeatureGetNext struct {
	BaseMessage
}

func (m FeatureGetNext) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_FEATURE_GETNEXT, m.Address, seqNum, nil)
}

func (m FeatureGetNext) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeFeatureGetNext(addr bidib.Address, data []byte) (FeatureGetNext, error) {
	var result FeatureGetNext
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// Query for a single feature. Followed by a byte with the feature number, which was queried.
// The node responds with MSG_FEATURE.
type FeatureGet struct {
	BaseMessage
	Feature bidib.FeatureID
}

func (m FeatureGet) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{byte(m.Feature)}
	bidib.EncodeMessage(write, bidib.MSG_FEATURE_GET, m.Address, seqNum, data)
}

func (m FeatureGet) String() string {
	return fmt.Sprintf("%T addr=%s feature=%s", m, m.Address, m.Feature)
}

func decodeFeatureGet(addr bidib.Address, data []byte) (FeatureGet, error) {
	var result FeatureGet
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.Feature = bidib.FeatureID(data[0])
	return result, nil
}

// Setting of a single feature. Followed by 2 bytes: feature number, value.
// The node responds with a MSG_FEATURE as confirmation.
// If a value has been sent that is not adjustable, the actual value which was set is returned.
type FeatureSet struct {
	BaseMessage
	Feature bidib.FeatureID
	Value   uint8
}

func (m FeatureSet) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{byte(m.Feature), m.Value}
	bidib.EncodeMessage(write, bidib.MSG_FEATURE_SET, m.Address, seqNum, data)
}

func (m FeatureSet) String() string {
	return fmt.Sprintf("%T addr=%s feature=%s value=0x%02x", m, m.Address, m.Feature, m.Value)
}

func decodeFeatureSet(addr bidib.Address, data []byte) (FeatureSet, error) {
	var result FeatureSet
	if err := validateDataLength(data, 2); err != nil {
		return result, err
	}
	result.Address = addr
	result.Feature = bidib.FeatureID(data[0])
	result.Value = data[1]
	return result, nil
}
