package bidib

import "fmt"

// Bidib error code
type ErrorCode uint8

const (
	//===============================================================================
	//
	// 5. Error Codes
	//
	//===============================================================================
	//
	// a) general error codes
	BIDIB_ERR_NONE              ErrorCode = 0x00 // void
	BIDIB_ERR_TXT               ErrorCode = 0x01 // general text error, 1..n byte characters
	BIDIB_ERR_CRC               ErrorCode = 0x02 // received crc was errornous, 1 byte with msg num
	BIDIB_ERR_SIZE              ErrorCode = 0x03 // missing parameters, 1 byte with msg num
	BIDIB_ERR_SEQUENCE          ErrorCode = 0x04 // sequence was wrong, 1 or 2 byte with last good and current num
	BIDIB_ERR_PARAMETER         ErrorCode = 0x05 // parameter out of range, 1 byte with msg num
	BIDIB_ERR_BUS               ErrorCode = 0x10 // Bus Fault, capacity exceeded, 1 byte fault code
	BIDIB_ERR_ADDRSTACK         ErrorCode = 0x11 // Address Stack, 4 bytes
	BIDIB_ERR_IDDOUBLE          ErrorCode = 0x12 // Double ID, 7 bytes
	BIDIB_ERR_SUBCRC            ErrorCode = 0x13 // Message in Subsystem had crc error, 1 byte node addr
	BIDIB_ERR_SUBTIME           ErrorCode = 0x14 // Message in Subsystem timed out, 1 byte node addr
	BIDIB_ERR_SUBPAKET          ErrorCode = 0x15 // Message in Subsystem Packet Size Error, 1..n byte node addr and data
	BIDIB_ERR_OVERRUN           ErrorCode = 0x16 // Message buffer in downstream overrun, messages lost.
	BIDIB_ERR_HW                ErrorCode = 0x20 // self test failed, 1 byte vendor error code
	BIDIB_ERR_RESET_REQUIRED    ErrorCode = 0x21 // reset needed (ie. due to reconfiguration)
	BIDIB_ERR_NO_SECACK_BY_HOST ErrorCode = 0x30 // Occupancy message was not mirrored by host as required
	//
	// b) error cause (2nd parameter)
	// for MSG_LC_NA
	BIDIB_ERR_LC_PORT_NONE     = 0x00 // no (more) error (internal use in nodes)
	BIDIB_ERR_LC_PORT_GENERAL  = 0x01 // unknown cause
	BIDIB_ERR_LC_PORT_UNKNOWN  = 0x02 // port not existing
	BIDIB_ERR_LC_PORT_INACTIVE = 0x03 // port not usable
	BIDIB_ERR_LC_PORT_EXEC     = 0x04 // exec not possible
	BIDIB_ERR_LC_PORT_BROKEN   = 0x7F // hardware failure
)

func (ec ErrorCode) String() string {
	switch ec {
	case BIDIB_ERR_NONE: //            ErrorCode = 0x00 // void
		return "None"
	case BIDIB_ERR_TXT: //      ErrorCode = 0x01 // general text error, 1..n byte characters
		return "General text"
	case BIDIB_ERR_CRC: //   ErrorCode = 0x02 // received crc was errornous, 1 byte with msg num
		return "CRC"
	case BIDIB_ERR_SIZE: //         ErrorCode = 0x03 // missing parameters, 1 byte with msg num
		return "Missing parameters"
	case BIDIB_ERR_SEQUENCE: //      ErrorCode = 0x04 // sequence was wrong, 1 or 2 byte with last good and current num
		return "Wrong sequence number"
	case BIDIB_ERR_PARAMETER: //       ErrorCode = 0x05 // parameter out of range, 1 byte with msg num
		return "Parameter out of range"
	case BIDIB_ERR_BUS: //    ErrorCode = 0x10 // Bus Fault, capacity exceeded, 1 byte fault code
		return "Bus fault"
	case BIDIB_ERR_ADDRSTACK: // ErrorCode = 0x11 // Address Stack, 4 bytes
		return "Address stack"
	case BIDIB_ERR_IDDOUBLE: //          ErrorCode = 0x12 // Double ID, 7 bytes
		return "Double ID"
	case BIDIB_ERR_SUBCRC: //       ErrorCode = 0x13 // Message in Subsystem had crc error, 1 byte node addr
		return "Subsystem CRC error"
	case BIDIB_ERR_SUBTIME: //    ErrorCode = 0x14 // Message in Subsystem timed out, 1 byte node addr
		return "Subsystem timeout"
	case BIDIB_ERR_SUBPAKET: // ErrorCode = 0x15 // Message in Subsystem Packet Size Error, 1..n byte node addr and data
		return "Subsystem package size error"
	case BIDIB_ERR_OVERRUN: //           ErrorCode = 0x16 // Message buffer in downstream overrun, messages lost.
		return "Overrun"
	case BIDIB_ERR_HW: //        ErrorCode = 0x20 // self test failed, 1 byte vendor error code
		return "Self test failed"
	case BIDIB_ERR_RESET_REQUIRED: //    ErrorCode = 0x21 // reset needed (ie. due to reconfiguration)
		return "Reset needed"
	case BIDIB_ERR_NO_SECACK_BY_HOST: // ErrorCode = 0x30 // Occupancy message was not mirrored by host as required
		return "Occupancy message was not mirrored by host"
	default:
		return fmt.Sprintf("0x%02x", uint8(ec))
	}
}
