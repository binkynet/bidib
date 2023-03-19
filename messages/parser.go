package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// Parse a message
func Parse(mType bidib.MessageType, addr bidib.Address, seqNum bidib.SequenceNumber, data []byte) (bidib.Message, error) {
	switch mType {
	// System common downstream
	case bidib.MSG_SYS_GET_MAGIC:
		return decodeSysGetMagic(addr, data)
	case bidib.MSG_SYS_GET_P_VERSION:
		return decodeSysGetPVersion(addr, data)
	case bidib.MSG_SYS_ENABLE:
		return decodeSysEnable(addr, data)
	case bidib.MSG_SYS_DISABLE:
		return decodeSysDisable(addr, data)
	case bidib.MSG_SYS_GET_UNIQUE_ID:
		return decodeSysGetUniqueID(addr, data)
	case bidib.MSG_SYS_GET_SW_VERSION:
		return decodeSysGetSwVersion(addr, data)
	case bidib.MSG_SYS_PING:
		return decodeSysPing(addr, data)
	case bidib.MSG_LOCAL_PING:
		return decodeLocalPing(addr, data)
	case bidib.MSG_SYS_IDENTIFY:
		return decodeSysIdentify(addr, data)
	case bidib.MSG_SYS_GET_ERROR:
		return decodeSysGetError(addr, data)
	case bidib.MSG_LOCAL_SYNC:
		return decodeLocalSync(addr, data)

	// System bus management downstream
	case bidib.MSG_SYS_RESET:
		return decodeSysReset(addr, data)
	case bidib.MSG_NODETAB_GETALL:
		return decodeNodeTabGetAll(addr, data)
	case bidib.MSG_NODETAB_GETNEXT:
		return decodeNodeTabGetNext(addr, data)
	case bidib.MSG_GET_PKT_CAPACITY:
		return decodeGetPktCapacity(addr, data)
	case bidib.MSG_NODE_CHANGED_ACK:
		return decodeNodeChangedAck(addr, data)
	case bidib.MSG_LOGON_ACK:
		return decodeLocalLogonAck(addr, data)
	case bidib.MSG_LOGON_REJECTED:
		return decodeLocalLogonRejected(addr, data)

	// System layout management downstream
	case bidib.MSG_SYS_CLOCK:
		return decodeSysClock(addr, data)

	// System upstream
	case bidib.MSG_SYS_MAGIC:
		return decodeSysMagic(addr, data)
	case bidib.MSG_SYS_PONG:
		return decodeSysPong(addr, data)
	case bidib.MSG_LOCAL_PONG:
		return decodeLocalPong(addr, data)
	case bidib.MSG_SYS_P_VERSION:
		return decodeSysPVersion(addr, data)
	case bidib.MSG_SYS_UNIQUE_ID:
		return decodeSysUniqueID(addr, data)
	case bidib.MSG_SYS_SW_VERSION:
		return decodeSysSwVersion(addr, data)
	case bidib.MSG_SYS_IDENTIFY_STATE:
		return decodeSysIdentityState(addr, data)
	case bidib.MSG_SYS_ERROR:
		return decodeSysError(addr, data)

	default:
		return nil, fmt.Errorf("failed to parse message of type %s", mType)
	}
}
