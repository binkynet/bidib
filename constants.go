package bidib

const (
	BIDIB_VERSION = (0*256 + 7)

	BIDIB_SYS_MAGIC  = 0xAFFE // full featured BiDiB-Node
	BIDIB_BOOT_MAGIC = 0xB00D // reduced Node, bootloader only

	//===============================================================================
	//
	// 6. FW Update (useful defines)
	//
	//===============================================================================

	BIDIB_MSG_FW_UPDATE_OP_ENTER   = 0x00 // node should enter update mode
	BIDIB_MSG_FW_UPDATE_OP_EXIT    = 0x01 // node should leave update mode
	BIDIB_MSG_FW_UPDATE_OP_SETDEST = 0x02 // set destination memory
	BIDIB_MSG_FW_UPDATE_OP_DATA    = 0x03 // data chunk
	BIDIB_MSG_FW_UPDATE_OP_DONE    = 0x04 // end of data

	BIDIB_MSG_FW_UPDATE_STAT_READY = 0   // ready
	BIDIB_MSG_FW_UPDATE_STAT_EXIT  = 1   // exit ack'd
	BIDIB_MSG_FW_UPDATE_STAT_DATA  = 2   // waiting for data
	BIDIB_MSG_FW_UPDATE_STAT_ERROR = 255 // there was an error

	BIDIB_FW_UPDATE_ERROR_NO_DEST  = 1 // destination not yet set
	BIDIB_FW_UPDATE_ERROR_RECORD   = 2 // error in hex record type
	BIDIB_FW_UPDATE_ERROR_ADDR     = 3 // record out of range
	BIDIB_FW_UPDATE_ERROR_CHECKSUM = 4 // checksum error on record
	BIDIB_FW_UPDATE_ERROR_SIZE     = 5 // size error
	BIDIB_FW_UPDATE_ERROR_APPCRC   = 6 // crc error on application, cant start

	//===============================================================================
	//
	// 7. System Messages, Serial Link, BiDiBus
	//
	//===============================================================================

	// 6.a) Serial Link

	BIDIB_PKT_MAGIC  = 0xFE // frame delimiter for serial link
	BIDIB_PKT_ESCAPE = 0xFD

	// 6.b) defines for BiDiBus, system messages
	// (system messages are 9 bits, bit8 is set (1), bits 0..7 do have even parity)
	BIDIBUS_SYS_MSG = 0x40 // System Part of BiDiBus

	BIDIBUS_POWER_UPx     = 0x7F // formerly Bus Reset (now reserved)
	BIDIBUS_POWER_UPx_par = 0xFF // formerly Bus Reset (including parity)
	BIDIBUS_LOGON         = 0x7E // Logon Prompt
	BIDIBUS_LOGON_par     = 0x7E // Logon Prompt (including parity)
	BIDIBUS_BUSY          = 0x7D // Interface Busy
	BIDIBUS_BUSY_par      = 0x7D // Interface Busy (including parity)

	// from Node
	BIDIBUS_NODE_READY = 0
	BIDIBUS_NODE_BUSY  = 1

	//===============================================================================
	//
	// 8. Booster and Command Station Handling (useful defines)
	//
	//===============================================================================

	BIDIB_BST_STATE_OFF         = 0x00 // Booster turned off
	BIDIB_BST_STATE_OFF_SHORT   = 0x01 // Booster is off, output shortend
	BIDIB_BST_STATE_OFF_HOT     = 0x02 // Booster off and too hot
	BIDIB_BST_STATE_OFF_NOPOWER = 0x03 // Booster has no mains
	BIDIB_BST_STATE_OFF_GO_REQ  = 0x04 // Booster off and local go request is present
	BIDIB_BST_STATE_OFF_HERE    = 0x05 // Booster off (was turned off by a local key)
	BIDIB_BST_STATE_OFF_NO_DCC  = 0x06 // Booster is off (no DCC input)
	BIDIB_BST_STATE_ON          = 0x80 // Booster on
	BIDIB_BST_STATE_ON_LIMIT    = 0x81 // Booster on and critical current flows
	BIDIB_BST_STATE_ON_HOT      = 0x82 // Booster on and is getting hot
	BIDIB_BST_STATE_ON_STOP_REQ = 0x83 // Booster on and a local stop request is present
	BIDIB_BST_STATE_ON_HERE     = 0x84 // Booster on (was turned on by a local key)

	BIDIB_BST_DIAG_I = 0x00 // Current
	BIDIB_BST_DIAG_V = 0x01 // Voltage
	BIDIB_BST_DIAG_T = 0x02 // Temperatur

	BIDIB_CS_DRIVE_SPEED_BIT  = (1 << 0)
	BIDIB_CS_DRIVE_F1F4_BIT   = (1 << 1) // also FL
	BIDIB_CS_DRIVE_F0F4_BIT   = (1 << 1) // additional define, it is the same bit
	BIDIB_CS_DRIVE_F5F8_BIT   = (1 << 2)
	BIDIB_CS_DRIVE_F9F12_BIT  = (1 << 3)
	BIDIB_CS_DRIVE_F13F20_BIT = (1 << 4)
	BIDIB_CS_DRIVE_F21F28_BIT = (1 << 5)

	BIDIB_CS_PROG_START         = 0x00 // service mode answer (MSG_CS_PROG_STATE)
	BIDIB_CS_PROG_RUNNING       = 0x01 // generic rule:  MSB: 0: running, 1: finished
	BIDIB_CS_PROG_OKAY          = 0x80 //               Bit6: 0: okay,    1: fail
	BIDIB_CS_PROG_STOPPED       = 0xC0
	BIDIB_CS_PROG_NO_LOCO       = 0xC1
	BIDIB_CS_PROG_NO_ANSWER     = 0xC2
	BIDIB_CS_PROG_SHORT         = 0xC3
	BIDIB_CS_PROG_VERIFY_FAILED = 0xC4

	//===============================================================================
	//
	// 9. IO-Control and Macro (useful defines)
	//
	//===============================================================================

	// Accessory parameter
	BIDIB_ACCESSORY_PARA_HAS_ESTOP = 250 // following boolean declares whether ESTOP is available
	BIDIB_ACCESSORY_PARA_OPMODE    = 251 // following data links the mode accessory that governs the composite
	BIDIB_ACCESSORY_PARA_STARTUP   = 252 // following data defines initialisation behavior
	BIDIB_ACCESSORY_PARA_MACROMAP  = 253 // following data defines a mapping
	BIDIB_ACCESSORY_SWITCH_TIME    = 254 //
	BIDIB_ACCESSORY_PARA_NOTEXIST  = 255 // following data contains the number of the unknown parameter

	// Accessory aspects
	BIDIB_ACCESSORY_ASPECT_STOP      = 0x00 // stop signal, idle state
	BIDIB_ACCESSORY_ASPECT_OPERATING = 0x01 // normal operation (for operating mode accessories)
	BIDIB_ACCESSORY_ASPECT_ESTOP     = 0xFE // emergency stop (during movement)
	BIDIB_ACCESSORY_ASPECT_UNKNOWN   = 0xFF // illegal aspect (error void) or unknown aspect (error feedback)

	// Accessory states
	BIDIB_ACC_STATE_DONE            = 0x00 // done
	BIDIB_ACC_STATE_WAIT            = 0x01 // not done, time (like railcom spec) following
	BIDIB_ACC_STATE_NO_FB_AVAILABLE = 0x02 // ...and no feedback available
	BIDIB_ACC_STATE_ERROR           = 0x80 // error, error code following

	BIDIB_ACC_STATE_ERROR_MORE     = 0x40 // more errors are present
	BIDIB_ACC_STATE_ERROR_NONE     = 0x00 // no (more) errors
	BIDIB_ACC_STATE_ERROR_VOID     = 0x01 // no processing possible, illegal aspect
	BIDIB_ACC_STATE_ERROR_CURRENT  = 0x02 // current comsumption to high
	BIDIB_ACC_STATE_ERROR_LOWPOWER = 0x03 // supply too low
	BIDIB_ACC_STATE_ERROR_FUSE     = 0x04 // fuse blown
	BIDIB_ACC_STATE_ERROR_TEMP     = 0x05 // temp too high
	BIDIB_ACC_STATE_ERROR_POSITION = 0x06 // feedback error
	BIDIB_ACC_STATE_ERROR_MAN_OP   = 0x07 // manually operated
	BIDIB_ACC_STATE_ERROR_BULB     = 0x10 // bulb blown
	BIDIB_ACC_STATE_ERROR_SERVO    = 0x20 // servo broken
	BIDIB_ACC_STATE_ERROR_SELFTEST = 0x3F // internal error

	// Accessory state details

	BIDIB_ACC_DETAIL_CURR_ANGLE1DEG5   = 0x01 // uint8   current rotation angle, values 0..239 [unit 1.5°]
	BIDIB_ACC_DETAIL_TARGET_ANGLE1DEG5 = 0x02 // uint8   targeted rotation angle, values 0..239 [unit 1.5°]
	BIDIB_ACC_DETAIL_TIMESTAMP         = 0x40 // uint16  system timestamp

	// Macro / Output Portparameters
	// type codes
	BIDIB_PORTTYPE_SWITCH     = 0  // standard port (on/off)
	BIDIB_PORTTYPE_LIGHT      = 1  // light port
	BIDIB_PORTTYPE_SERVO      = 2  // servo port
	BIDIB_PORTTYPE_SOUND      = 3  // sound
	BIDIB_PORTTYPE_MOTOR      = 4  // motor
	BIDIB_PORTTYPE_ANALOGOUT  = 5  // analog
	BIDIB_PORTTYPE_BACKLIGHT  = 6  // backlight (different operation then light port)
	BIDIB_PORTTYPE_SWITCHPAIR = 7  // width: 2, exclusive usage
	BIDIB_PORTTYPE_INPUT      = 15 // simple input (open/closed)
	/* deprecated names (as of revision 1.24), do not use! */
	BIDIB_OUTTYPE_SPORT     = 0 // (deprecated) standard port
	BIDIB_OUTTYPE_LPORT     = 1 // (deprecated) light port
	BIDIB_OUTTYPE_SERVO     = 2 // (deprecated) servo port
	BIDIB_OUTTYPE_SOUND     = 3 // (deprecated) sound
	BIDIB_OUTTYPE_MOTOR     = 4 // (deprecated) motor
	BIDIB_OUTTYPE_ANALOG    = 5 // (deprecated) analog
	BIDIB_OUTTYPE_BACKLIGHT = 6 // (deprecated) backlight

	// Port configuration ENUMs (P_ENUM)
	// P_ENUM 0..63:     8 bit values
	// P_ENUM 64..127:  16 bit values
	// P_ENUM 128..191: 24 bit values
	// P_ENUM 192..254: reserved
	// P_ENUM 255:       0 bit values
	BIDIB_PCFG_NONE           = 0x00 // uint8   no parameters available / error code
	BIDIB_PCFG_LEVEL_PORT_ON  = 0x01 // uint8   'analog' value for ON
	BIDIB_PCFG_LEVEL_PORT_OFF = 0x02 // uint8   'analog' value for OFF
	BIDIB_PCFG_DIMM_UP        = 0x03 // uint8   rate of increase for dimm up [unit 1/255 absolute brightness per 10ms]
	BIDIB_PCFG_DIMM_DOWN      = 0x04 // uint8   rate of decrease for dimm down [unit 1/255 absolute brightness per 10ms]
	BIDIB_PCFG_OUTPUT_MAP     = 0x06 // uint8   if there is a output mapping (like DMX)
	BIDIB_PCFG_SERVO_ADJ_L    = 0x07 // uint8   Servo Adjust Low
	BIDIB_PCFG_SERVO_ADJ_H    = 0x08 // uint8   Servo Adjust High
	BIDIB_PCFG_SERVO_SPEED    = 0x09 // uint8   Servo Speed
	BIDIB_PCFG_IO_CTRL        = 0x0a // uint8   (deprecated) IO setup
	BIDIB_PCFG_TICKS          = 0x0b // uint8   puls time for output [unit 10ms]
	BIDIB_PCFG_IS_PAIRED      = 0x0c // bool    (deprecated) not used, reserved
	BIDIB_PCFG_SWITCH_CTRL    = 0x0d // uint8   electrical behaviour for switch ports, as two nibbles
	BIDIB_PCFG_INPUT_CTRL     = 0x0e // uint8   electrical behaviour for input ports
	// 16 bit values
	BIDIB_PCFG_DIMM_UP_8_8   = 0x43 // uint16  rate of increase for dimm up [unit 1/65535 absolute brightness per 10ms]
	BIDIB_PCFG_DIMM_DOWN_8_8 = 0x44 // uint16  rate of decrease for dimm down [unit 1/65535 absolute brightness per 10ms]
	BIDIB_PCFG_PAIRED_PORT   = 0x45 // uint16  (deprecated) not used, reserved
	// 24 bit values
	BIDIB_PCFG_RGB      = 0x80 // uint24  RGB value of a coloured output. first byte R, second G, third B
	BIDIB_PCFG_RECONFIG = 0x81 // uint24  Reconfiguration: ACT_TYPE PORTMAP_L PORTMAP_H (only if FEATURE_CTRL_PORT_FLAT_MODEL > 0)
	// special
	BIDIB_PCFG_CONTINUE = 0xFF // none    an addtional message will follow

	// control codes - limited to one nibble, here for PORTs

	BIDIB_PORT_TURN_OFF     = 0 // for standard
	BIDIB_PORT_TURN_ON      = 1 // for standard
	BIDIB_PORT_TO_0         = 0 // for pair
	BIDIB_PORT_TO_1         = 1 // for pair
	BIDIB_PORT_DIMM_OFF     = 2
	BIDIB_PORT_DIMM_ON      = 3
	BIDIB_PORT_TURN_ON_NEON = 4
	BIDIB_PORT_BLINK_A      = 5
	BIDIB_PORT_BLINK_B      = 6
	BIDIB_PORT_FLASH_A      = 7
	BIDIB_PORT_FLASH_B      = 8
	BIDIB_PORT_DOUBLE_FLASH = 9
	BIDIB_PORT_QUERY        = 15

	// Macro Global States
	BIDIB_MACRO_OFF      = 0x00
	BIDIB_MACRO_START    = 0x01
	BIDIB_MACRO_RUNNING  = 0x02
	BIDIB_MACRO_RESTORE  = 0xFC // 252
	BIDIB_MACRO_SAVE     = 0xFD // 253
	BIDIB_MACRO_DELETE   = 0xFE
	BIDIB_MACRO_NOTEXIST = 0xFF

	// Macro System Commands (Level 2)
	// These are opcodes inside a macro-syscommand of level 2
	BIDIB_MSYS_END_OF_MACRO  = 255 // end of macro (EOF)
	BIDIB_MSYS_START_MACRO   = 254 // start a macro
	BIDIB_MSYS_STOP_MACRO    = 253 // stop a macro
	BIDIB_MSYS_BEGIN_CRITCAL = 252 // current macro will ignore stop requests
	BIDIB_MSYS_END_CRITCAL   = 251 // current macro can be stopped by a stop (default)

	BIDIB_MSYS_FLAG_QUERY  = 250 // deprecated (by QUERY0 and QUERY1)
	BIDIB_MSYS_FLAG_QUERY1 = 250 // query flag and pause as long as flag is not set (advance if set)
	BIDIB_MSYS_FLAG_SET    = 249 // set flag
	BIDIB_MSYS_FLAG_CLEAR  = 248 // reset flag

	BIDIB_MSYS_INPUT_QUERY1 = 247 // query input for 'pressed / activated' and advance, if input is set
	BIDIB_MSYS_INPUT_QUERY0 = 246 // query input for 'released' and and advance, if input is released
	BIDIB_MSYS_DELAY_RANDOM = 245 // make a random delay
	BIDIB_MSYS_DELAY_FIXED  = 244 // make a fixed delay

	BIDIB_MSYS_ACC_OKAY_QIN1   = 243 // query input for 'pressed / activated' and send okay to accessory-module, if pressed, else send nok. (not waiting)
	BIDIB_MSYS_ACC_OKAY_QIN0   = 242 // query input for 'released' and send okay to accessory-module, if pressed, else send nok. (not waiting)
	BIDIB_MSYS_ACC_OKAY_NF     = 241 // send okay to accessory-module, no feedback available
	BIDIB_MSYS_SERVOMOVE_QUERY = 240 // query servo movement and pause as long as moving

	BIDIB_MSYS_FLAG_QUERY0 = 239 // query flag and pause as long as flag is set (advance if not set)

	// Macro global parameters
	BIDIB_MACRO_PARA_SLOWDOWN  = 0x01
	BIDIB_MACRO_PARA_REPEAT    = 0x02 // 0=forever, 1=once, 2..250 n times
	BIDIB_MACRO_PARA_START_CLK = 0x03 // TCODE defines Startpoint

	// here additional run parameters are to be defined. like:
	// start condition: from DCC, DCC addr low, DCC addr high
	//                  from system clock: time
	//                  from input
	// stop condition:

	//===============================================================================
	//
	// 10. Defines for RailcomPlus
	//
	//===============================================================================

	// phase generally in bit 0
	RC_P0 = (0 << 0)
	RC_P1 = (1 << 0)
	// type generally in bit 1
	RC_TYPE_LOCO = (0 << 1)
	RC_TYPE_ACC  = (1 << 1)

	// Note: in bidib, we have little endian; on DCC, the order is big endian.
	/*typedef union
	    {
	      uint8_t as_uint8[5];
	      struct
	        {
	          uint8_t mun_0;        // manufacturer unique number
	          uint8_t mun_1;
	          uint8_t mun_2;
	          uint8_t mun_3;
	          uint8_t mid;          // manufacturer ID (like DCC vendor ID)
	        };
	    } t_rcplus_unique_id;         // as Central ID or Decoder ID

	  typedef struct
	    {
	      t_rcplus_unique_id cid;
	      uint8_t sid;          // session number
	    } t_rcplus_tid;
	*/
	// a) for MSG_CS_RCPLUS

	RC_BIND      = 0 // 2:dec_mun[0],3:dec_mun[1],4:dec_mun[2],5:dec_mun[3], 6:dec_mid, 7:new_addrl, 8:new_addrh
	RC_PING      = 1 // 2:interval
	RC_GET_TID   = 2 // -
	RC_SET_TID   = 3 // 2:cid[0],3:cid[1],4:cid[2],5:cid[3],6:cid[4],7:sid
	RC_PING_ONCE = 4 // -
	RC_FIND      = 6 // 2:dec_mun[0],3:dec_mun[1],4:dec_mun[2],5:dec_mun[3], 6:dec_mid
	// expanded (redundant aliases):
	RC_PING_ONCE_P0 = (RC_PING_ONCE | RC_P0)
	RC_PING_ONCE_P1 = (RC_PING_ONCE | RC_P1)
	RC_FIND_P0      = (RC_FIND | RC_P0)
	RC_FIND_P1      = (RC_FIND | RC_P1)

	// b) for MSG_CS_RCPLUS_ACK

	//      RC_BIND                     0 // 2:ack, 3:dec_mun[0],4:dec_mun[1],5:dec_mun[2],6:dec_mun[3], 7:dec_mid
	//      RC_PING                     1 // 2:interval
	RC_TID = 2 // 2:cid[0],3:cid[1],4:cid[2],5:cid[3],6:cid[4], 7:sid
	//      RC_PING_ONCE_*              4 // 2:ack
	//      RC_FIND_*                   6 // 2:ack, 3:dec_mun[0],4:dec_mun[1],5:dec_mun[2],6:dec_mun[3], 7:dec_mid

	// c) for MSG_BM_RCPLUS

	RC_BIND_ACCEPTED  = (0 << 2) // 3:dec_mun[0],4:dec_mun[1],5:dec_mun[2],6:dec_mun[3], 7:dec_mid, 8:addr_l, 9:addr_h
	RC_COLLISION      = (1 << 2)
	RC_PING_COLLISION = (RC_COLLISION | (0 << 1)) // -
	RC_FIND_COLLISION = (RC_COLLISION | (1 << 1)) // 3:dec_mun[0],4:dec_mun[1],5:dec_mun[2],6:dec_mun[3], 7:dec_mid (of find command)
	RC_PONG_OKAY      = (2 << 2)                  // 3:dec_mun[0],4:dec_mun[1],5:dec_mun[2],6:dec_mun[3], 7:dec_mid (of found decoder)
	RC_PONG_NEW       = (3 << 2)                  // 3:dec_mun[0],4:dec_mun[1],5:dec_mun[2],6:dec_mun[3], 7:dec_mid (of found decoder)
	// expanded (redundant aliases):
	RC_BIND_ACCEPTED_LOCO      = (RC_BIND_ACCEPTED | RC_TYPE_LOCO) // no phase!
	RC_BIND_ACCEPTED_ACCESSORY = (RC_BIND_ACCEPTED | RC_TYPE_ACC)  // no phase!
	RC_PING_COLLISION_P0       = (RC_PING_COLLISION | RC_P0)       // no type!
	RC_PING_COLLISION_P1       = (RC_PING_COLLISION | RC_P1)       // no type!
	RC_FIND_COLLISION_P0       = (RC_FIND_COLLISION | RC_P0)       // no type!
	RC_FIND_COLLISION_P1       = (RC_FIND_COLLISION | RC_P1)       // no type!
	RC_PONG_OKAY_LOCO_P0       = (RC_PONG_OKAY | RC_TYPE_LOCO | RC_P0)
	RC_PONG_OKAY_LOCO_P1       = (RC_PONG_OKAY | RC_TYPE_LOCO | RC_P1)
	RC_PONG_OKAY_ACCESSORY_P0  = (RC_PONG_OKAY | RC_TYPE_ACC | RC_P0)
	RC_PONG_OKAY_ACCESSORY_P1  = (RC_PONG_OKAY | RC_TYPE_ACC | RC_P1)
	RC_PONG_NEW_LOCO_P0        = (RC_PONG_NEW | RC_TYPE_LOCO | RC_P0)
	RC_PONG_NEW_LOCO_P1        = (RC_PONG_NEW | RC_TYPE_LOCO | RC_P1)
	RC_PONG_NEW_ACCESSORY_P0   = (RC_PONG_NEW | RC_TYPE_ACC | RC_P0)
	RC_PONG_NEW_ACCESSORY_P1   = (RC_PONG_NEW | RC_TYPE_ACC | RC_P1)

	//===============================================================================
	//
	// 11. Defines for M4
	//
	//===============================================================================

	// reminder
	//  MSG_CS_M4     (MSG_DGEN + 0x09)
	//  MSG_CS_M4_ACK (MSG_UGEN + 0x09)

	// a) opcodes for MSG_CS_M4

	M4_GET_TID     = 0x00 // -
	M4_SET_TID     = 0x01 // followed by 4 Bytes UID (track identifier) and 2 Bytes SID (session counter), for the beacon to be sent.
	M4_SET_BEACON  = 0x02 // followed by 1 Byte with the time interval between beacon broadcasts. 0 means disabling.
	M4_SET_SEARCH  = 0x03 // followed by 1 Byte with the time interval between searches with C=0 U=0. 0 means disabling.
	M4_BIND_ADDR   = 0x04 // followed by 4 Bytes UID (decoder) and 2 Bytes address.
	M4_UNBIND_ADDR = 0x05 // followed by 2 Bytes with the on-track address.

	// b) opcodes for MSG_CS_M4_ACK

	M4_TID           = 0x80 // followed by 4 Bytes UID and 2 Bytes SID. This is the answer to GET_TID and SET_TID.
	M4_BEACON        = 0x82 // followed by 1 Byte with the interval. This is the answer to BEACON.
	M4_SEARCH        = 0x83 // followed by 1 Byte with the interval. This is the answer to SEARCH.
	M4_BIND_ADDR_ACK = 0x84 // followed by 1 Byte with the ack code (0:not available, 1:ok, 2:delayed, 5:worked and inquiry successful),
	// 4 Bytes DID and 2 Bytes address. This is the answer to BIND_ADDR.
	M4_UNBIND_ADDR_ACK = 0x85 // followed by 1 Byte ack, 2 Bytes address, similar to MSG_CS_DRIVE_ACK
	M4_NEW_LOCO        = 0x86 // followed by 4 Bytes with the DID of the newly found decoder.

	// c) opcodes for MSG_CS_POM/MSG_BM_XPOM

	BIDIB_CS_POM4_RD_BYTE1 = 0xC0
	BIDIB_CS_POM4_RD_BYTE2 = 0xC1
	BIDIB_CS_POM4_RD_BYTE4 = 0xC2
	BIDIB_CS_POM4_RD_BYTE8 = 0xC3 // reserved???
	BIDIB_CS_POM4_WR_BYTE  = 0xC4
)
