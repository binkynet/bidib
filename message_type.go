package bidib

import (
	"fmt"
)

// Type of message
type MessageType uint8

// String returns a human readable representation of a MessageType.
func (mt MessageType) String() string {
	return fmt.Sprintf("0x%02x", uint8(mt))
}

const (
	//===============================================================================
	//
	// 1. Defines for Downstream Messages
	//
	//===============================================================================
	//*// = broadcast messages, a interface must forward this to subnodes
	//      a node must not answer these messages, if not connected
	MSG_DSTRM MessageType = 0x00

	//-- system messages                                    // Parameters
	MSG_DSYS               MessageType = (MSG_DSTRM + 0x00)
	MSG_SYS_GET_MAGIC                  = (MSG_DSYS + 0x01) // - // these must stay here
	MSG_SYS_GET_P_VERSION              = (MSG_DSYS + 0x02) // - // these must stay here
	MSG_SYS_ENABLE                     = (MSG_DSYS + 0x03) //*// -
	MSG_SYS_DISABLE                    = (MSG_DSYS + 0x04) //*// -
	MSG_SYS_GET_UNIQUE_ID              = (MSG_DSYS + 0x05) // -
	MSG_SYS_GET_SW_VERSION             = (MSG_DSYS + 0x06) // -
	MSG_SYS_PING                       = (MSG_DSYS + 0x07) // 1:dat
	MSG_SYS_IDENTIFY                   = (MSG_DSYS + 0x08) // 1:id_state
	MSG_SYS_RESET                      = (MSG_DSYS + 0x09) //*// -
	MSG_GET_PKT_CAPACITY               = (MSG_DSYS + 0x0a) // -
	MSG_NODETAB_GETALL                 = (MSG_DSYS + 0x0b) // -
	MSG_NODETAB_GETNEXT                = (MSG_DSYS + 0x0c) // -
	MSG_NODE_CHANGED_ACK               = (MSG_DSYS + 0x0d) // 1:nodetab_version
	MSG_SYS_GET_ERROR                  = (MSG_DSYS + 0x0e) // -
	MSG_FW_UPDATE_OP                   = (MSG_DSYS + 0x0f) // 1:opcode, 2..n parameters

	//-- feature and user config messages
	MSG_DFC             = (MSG_DSTRM + 0x10)
	MSG_FEATURE_GETALL  = (MSG_DFC + 0x00) // -
	MSG_FEATURE_GETNEXT = (MSG_DFC + 0x01) // -
	MSG_FEATURE_GET     = (MSG_DFC + 0x02) // 1:feature_num
	MSG_FEATURE_SET     = (MSG_DFC + 0x03) // 1:feature_num, 2:feature_val
	MSG_VENDOR_ENABLE   = (MSG_DFC + 0x04) // 1-7: unique-id of node
	MSG_VENDOR_DISABLE  = (MSG_DFC + 0x05) // -
	MSG_VENDOR_SET      = (MSG_DFC + 0x06) // V_NAME,V_VALUE
	MSG_VENDOR_GET      = (MSG_DFC + 0x07) // V_NAME
	MSG_SYS_CLOCK       = (MSG_DFC + 0x08) //*// 1:TCODE0, 2:TCODE1, 3:TCODE2, 4:TCODE3
	MSG_STRING_GET      = (MSG_DFC + 0x09) // 1:Nspace, 2:ID
	MSG_STRING_SET      = (MSG_DFC + 0x0a) // 1:Nspace, 2:ID, 3:Strsize, 4...n: string

	//-- occupancy messages
	MSG_DBM                = (MSG_DSTRM + 0x20)
	MSG_BM_GET_RANGE       = (MSG_DBM + 0x00) // 1:start, 2:end
	MSG_BM_MIRROR_MULTIPLE = (MSG_DBM + 0x01) // 1:start, 2:size, 3..n:data
	MSG_BM_MIRROR_OCC      = (MSG_DBM + 0x02) // 1:mnum
	MSG_BM_MIRROR_FREE     = (MSG_DBM + 0x03) // 1:mnum
	MSG_BM_ADDR_GET_RANGE  = (MSG_DBM + 0x04) // 1:start, 2:end
	MSG_BM_GET_CONFIDENCE  = (MSG_DBM + 0x05) // -
	MSG_BM_MIRROR_POSITION = (MSG_DBM + 0x06) // 1:addr_l, 2:addr_h, 3:type, 4:location_id_l, 5:location_id_h

	//-- booster messages
	MSG_DBST        = (MSG_DSTRM + 0x30)
	MSG_BOOST_OFF   = (MSG_DBST + 0x00) //*// 1:unicast
	MSG_BOOST_ON    = (MSG_DBST + 0x01) //*// 1:unicast
	MSG_BOOST_QUERY = (MSG_DBST + 0x02) // -

	//-- accessory control messages
	MSG_DACC               = (MSG_DSTRM + 0x38)
	MSG_ACCESSORY_SET      = (MSG_DACC + 0x00) // 1:anum, 2:aspect
	MSG_ACCESSORY_GET      = (MSG_DACC + 0x01) // 1:anum
	MSG_ACCESSORY_PARA_SET = (MSG_DACC + 0x02) // 1:anum, 2:para_num, 3..n: data
	MSG_ACCESSORY_PARA_GET = (MSG_DACC + 0x03) // 1:anum, 2:para_num

	//-- switch/light/servo control messages
	MSG_DLC                = (MSG_DSTRM + 0x3F)
	MSG_LC_PORT_QUERY_ALL  = (MSG_DLC + 0x00) // 1:selL, 2:selH, [3:startL, 4:startH, 5:endL, 6:endH]
	MSG_LC_OUTPUT          = (MSG_DLC + 0x01) // 1,2:port, 3:state
	MSG_LC_CONFIG_SET      = (MSG_DLC + 0x02) // (deprecated) 1:type, 2:num, 3:off_val, 4:on_val, 5:dimm_off, 6:dimm_on
	MSG_LC_CONFIG_GET      = (MSG_DLC + 0x03) // (deprecated) 1:type, 2:num
	MSG_LC_KEY_QUERY       = (MSG_DLC + 0x04) // (deprecated) 1:num
	MSG_LC_OUTPUT_QUERY    = (MSG_DLC + 0x05) // (deprecated) 1,2:port
	MSG_LC_PORT_QUERY      = (MSG_DLC + 0x05) // 1,2:port
	MSG_LC_CONFIGX_GET_ALL = (MSG_DLC + 0x06) // [1:startL, 2:startH, 3:endL, 4:endH]
	MSG_LC_CONFIGX_SET     = (MSG_DLC + 0x07) // 1,2:port, [3:p_enum, 4:p_val]  (up to 16)
	MSG_LC_CONFIGX_GET     = (MSG_DLC + 0x08) // 1,2:port

	//-- macro messages
	MSG_DMAC              = (MSG_DSTRM + 0x48)
	MSG_LC_MACRO_HANDLE   = (MSG_DMAC + 0x00) // 1:macro, 2:opcode
	MSG_LC_MACRO_SET      = (MSG_DMAC + 0x01) // 1:macro, 2:item, 3:delay, 4:lstate, 5:lvalue, 6:0 / 4:port[0], 5:port[1], 6:portstat
	MSG_LC_MACRO_GET      = (MSG_DMAC + 0x02) // 1:macro, 2:item
	MSG_LC_MACRO_PARA_SET = (MSG_DMAC + 0x03) // 1:macro, 2:para_idx, 3,4,5,6:value
	MSG_LC_MACRO_PARA_GET = (MSG_DMAC + 0x04) // 1:macro, 2:para_idx

	//-- dcc gen messages
	MSG_DGEN         = (MSG_DSTRM + 0x60)
	MSG_CS_ALLOCATE  = (MSG_DGEN + 0x00)
	MSG_CS_SET_STATE = (MSG_DGEN + 0x02) // 1:state
	MSG_CS_DRIVE     = (MSG_DGEN + 0x04) // 1:addrl, 2:addrh, 3:format, 4:active, 5:speed, 6:1-4, 7:5-12, 8:13-20, 9:21-28
	MSG_CS_ACCESSORY = (MSG_DGEN + 0x05) // 1:addrl, 2:addrh, 3:data(aspect), 4:time_l, 5:time_h
	MSG_CS_BIN_STATE = (MSG_DGEN + 0x06) // 1:addrl, 2:addrh, 3:bin_statl, 4:bin_stath
	MSG_CS_POM       = (MSG_DGEN + 0x07) // 1..4:addr/did, 5:MID, 6:opcode, 7:cv_l, 8:cv_h, 9:cv_x, 10..13: data
	MSG_CS_RCPLUS    = (MSG_DGEN + 0x08) // 1:opcode, [2..n:parameter]

	// #define MSG_CS_QUERY (MSG_DGEN + 0x09) // 1:what (1=loco list, 2... tbd.)
	// experimental
	MSG_CS_QUERY = MSG_DGEN + 0x0A

	//-- service mode
	MSG_CS_PROG = (MSG_DGEN + 0x0F) // 1:opcode, 2:cv_l, 3:cv_h, 4: data

	//-- local message
	MSG_DLOCAL          = (MSG_DSTRM + 0x70)  // only locally used
	MSG_LOGON_ACK       = (MSG_DLOCAL + 0x00) // 1:node_addr, 2..8:unique_id
	MSG_LOCAL_PING      = (MSG_DLOCAL + 0x01)
	MSG_LOGON_REJECTED  = (MSG_DLOCAL + 0x02) // 1..7:unique_id
	MSG_LOCAL_ACCESSORY = (MSG_DLOCAL + 0x03) //*// 1:statusflag, 2,3: DCC-accessory addr
	MSG_LOCAL_SYNC      = (MSG_DLOCAL + 0x04) //*// 1:time_l, 2:time_h

	//===============================================================================
	//
	// 2. Defines for Upstream Messages
	//
	//===============================================================================

	MSG_USTRM MessageType = 0x80

	//-- system messages
	MSG_USYS               = (MSG_USTRM + 0x00)
	MSG_SYS_MAGIC          = (MSG_USYS + 0x01) // 1:0xFE 2:0xAF
	MSG_SYS_PONG           = (MSG_USYS + 0x02) // 1:mirrored dat
	MSG_SYS_P_VERSION      = (MSG_USYS + 0x03) // 1:proto-ver_l, 2:proto-ver_h
	MSG_SYS_UNIQUE_ID      = (MSG_USYS + 0x04) // 1:class, 2:classx, 3:vid, 4..7:pid+uid, [7..11: config_fingerprint]
	MSG_SYS_SW_VERSION     = (MSG_USYS + 0x05) // 1:sw-ver_l, 2:sw_-ver_h, 3:sw-ver_u
	MSG_SYS_ERROR          = (MSG_USYS + 0x06) // 1:err_code, 2:msg
	MSG_SYS_IDENTIFY_STATE = (MSG_USYS + 0x07) // 1:state
	MSG_NODETAB_COUNT      = (MSG_USYS + 0x08) // 1:length
	MSG_NODETAB            = (MSG_USYS + 0x09) // 1:version, 2:local num, 3..9: unique
	MSG_PKT_CAPACITY       = (MSG_USYS + 0x0a) // 1:capacity
	MSG_NODE_NA            = (MSG_USYS + 0x0b) // 1:node
	MSG_NODE_LOST          = (MSG_USYS + 0x0c) // 1:node
	MSG_NODE_NEW           = (MSG_USYS + 0x0d) // 1:version, 2:local num, 3..9: unique
	MSG_STALL              = (MSG_USYS + 0x0e) // 1:state
	MSG_FW_UPDATE_STAT     = (MSG_USYS + 0x0f) // 1:stat, 2:timeout

	//-- feature and user config messages
	MSG_UFC           = (MSG_USTRM + 0x10)
	MSG_FEATURE       = (MSG_UFC + 0x00) // 1:feature_num, 2:data
	MSG_FEATURE_NA    = (MSG_UFC + 0x01) // 1:feature_num
	MSG_FEATURE_COUNT = (MSG_UFC + 0x02) // 1:count
	MSG_VENDOR        = (MSG_UFC + 0x03) // 1..n: length,'string',length,'value'
	MSG_VENDOR_ACK    = (MSG_UFC + 0x04) // 1:mode
	MSG_STRING        = (MSG_UFC + 0x05) // 1:namespace, 2:id, 3:stringsize, 4...n: string

	//-- occupancy and bidi-detection messages
	MSG_UBM           = (MSG_USTRM + 0x20)
	MSG_BM_OCC        = (MSG_UBM + 0x00) // 1:mnum, [2,3:time_l, time_h]
	MSG_BM_FREE       = (MSG_UBM + 0x01) // 1:mnum
	MSG_BM_MULTIPLE   = (MSG_UBM + 0x02) // 1:base, 2:size; 3..n:data
	MSG_BM_ADDRESS    = (MSG_UBM + 0x03) // 1:mnum, [2,3:addr_l, addr_h]
	MSG_BM_ACCESSORY  = (MSG_UBM + 0x04) // (reserved, do not use yet) 1:mnum, [2,3:addr_l, addr_h]
	MSG_BM_CV         = (MSG_UBM + 0x05) // 1:addr_l, 2:addr_h, 3:cv_addr_l, 4:cv_addr_h, 5:cv_dat
	MSG_BM_SPEED      = (MSG_UBM + 0x06) // 1:addr_l, 2:addr_h, 3:speed_l, 4:speed_h (from loco)
	MSG_BM_CURRENT    = (MSG_UBM + 0x07) // 1:mnum, 2:current
	MSG_BM_BLOCK_CV   = (MSG_UBM + 0x08) // (deprecated) 1:decvid, 2..5:decuid, 6:offset, 7:idxl, 8:idxh, 9..12:data
	MSG_BM_XPOM       = (MSG_UBM + 0x08) // 1..4:addr/did, 5:0/vid, 6:opcode, 7:cv_l, 8:cv_h, 9:cv_x, 10[..13]: data
	MSG_BM_CONFIDENCE = (MSG_UBM + 0x09) // 1:void, 2:freeze, 3:nosignal
	MSG_BM_DYN_STATE  = (MSG_UBM + 0x0a) // 1:mnum, 2:addr_l, 3:addr_h, 4:dyn_num, 5:value (from loco)
	MSG_BM_RCPLUS     = (MSG_UBM + 0x0b) // 1:mnum, 2:opcode, [3..n:parameter]
	MSG_BM_POSITION   = (MSG_UBM + 0x0c) // 1:addr_l, 2:addr_h, 3:type, 4:location_id_l, 5:location_id_h

	//-- booster messages
	MSG_UBST             = (MSG_USTRM + 0x30)
	MSG_BOOST_STAT       = (MSG_UBST + 0x00) // 1:state (see defines below)
	MSG_BOOST_CURRENT    = (MSG_UBST + 0x01) // (deprecated by DIAGNOSTIC with V0.10) 1:current
	MSG_BOOST_DIAGNOSTIC = (MSG_UBST + 0x02) // [1:enum, 2:value],[3:enum, 4:value] ...
	//                              (MSG_UBST + 0x03)       // was reserved for MSG_NEW_DECODER (deprecated) 1:mnum, 2: dec_vid, 3..6:dec_uid
	//                              (MSG_UBST + 0x04)       // was reserved for MSG_ID_SEARCH_ACK (deprecated) 1:mnum, 2: s_vid, 3..6:s_uid[0..3], 7: dec_vid, 8..11:dec_uid
	//                              (MSG_UBST + 0x05)       // was reserved for MSG_ADDR_CHANGE_ACK (deprecated) 1:mnum, 2: dec_vid, 3..6:dec_uid, 7:addr_l, 8:addr_h

	//-- accessory control messages
	MSG_UACC             = (MSG_USTRM + 0x38)
	MSG_ACCESSORY_STATE  = (MSG_UACC + 0x00) // 1:anum, 2:aspect, 3:total, 4:execute, 5:wait, [6..n:details] (Quittung)
	MSG_ACCESSORY_PARA   = (MSG_UACC + 0x01) // 1:anum, 2:para_num, 3..n: data
	MSG_ACCESSORY_NOTIFY = (MSG_UACC + 0x02) // 1:anum, 2:aspect, 3:total, 4:execute, 5:wait, [6..n:details] (Spontan)

	//-- switch/light control messages
	MSG_ULC       = (MSG_USTRM + 0x40)
	MSG_LC_STAT   = (MSG_ULC + 0x00) // 1,2:port, 3:state
	MSG_LC_NA     = (MSG_ULC + 0x01) // 1,2:port, [3:errcause]
	MSG_LC_CONFIG = (MSG_ULC + 0x02) // (deprecated) 1:type, 2:num, 3:off_val, 4:on_val, 5:dimm_off, 6:dimm_on
	MSG_LC_KEY    = (MSG_ULC + 0x03) // (deprecated) 1:num, 2:state
	MSG_LC_WAIT   = (MSG_ULC + 0x04) // 1,2:port, 3:time
	//                              (MSG_ULC + 0x05)        was reserved for MGS_LC_MAPPING (deprecated)
	MSG_LC_CONFIGX = (MSG_ULC + 0x06) // 1,2:port, [3:p_enum, 4:p_val]  (up to 16)

	//-- macro messages
	MSG_UMAC           = (MSG_USTRM + 0x48)
	MSG_LC_MACRO_STATE = (MSG_UMAC + 0x00) // 1:macro, 2:opcode
	MSG_LC_MACRO       = (MSG_UMAC + 0x01) // 1:macro, 2:item, 3:delay, 4:lstate, 5:lvalue, 6:0 / 4:port[0], 5:port[1], 6:portstat
	MSG_LC_MACRO_PARA  = (MSG_UMAC + 0x02) // 1:macro, 2:para_idx, 3,4,5,6:value

	//-- dcc control messages
	MSG_UGEN                = (MSG_USTRM + 0x60)
	MSG_CS_ALLOC_ACK        = (MSG_UGEN + 0x00) // noch genauer zu klaeren / to be specified
	MSG_CS_STATE            = (MSG_UGEN + 0x01)
	MSG_CS_DRIVE_ACK        = (MSG_UGEN + 0x02)
	MSG_CS_ACCESSORY_ACK    = (MSG_UGEN + 0x03) // 1:addrl, 2:addrh, 3:ack
	MSG_CS_POM_ACK          = (MSG_UGEN + 0x04) // 1:addrl, 2:addrh, 3:addrxl, 4:addrxh, 5:mid, 6:ack
	MSG_CS_DRIVE_MANUAL     = (MSG_UGEN + 0x05) // 1:addrl, 2:addrh, 3:format, 4:active, 5:speed, 6:1-4, 7:5-12, 8:13-20, 9:21-28
	MSG_CS_DRIVE_EVENT      = (MSG_UGEN + 0x06) // 1:addrl, 2:addrh, 3:eventtype, Parameters
	MSG_CS_ACCESSORY_MANUAL = (MSG_UGEN + 0x07) // 1:addrl, 2:addrh, 3:ack
	MSG_CS_RCPLUS_ACK       = (MSG_UGEN + 0x08) // 1:opcode, [2..n:parameter]

	//-- service mode
	MSG_CS_PROG_STATE = (MSG_UGEN + 0x0F) // 1: state, 2:time, 3:cv_l, 4:cv_h, 5:data

	//-- local message
	MSG_ULOCAL     = (MSG_USTRM + 0x70) // only locally used
	MSG_LOGON      = (MSG_ULOCAL + 0x00)
	MSG_LOCAL_PONG = (MSG_ULOCAL + 0x01) // only locally used
)
