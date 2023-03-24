package bidib

const (
	BIDIB_VERSION = (0*256 + 7)

	BIDIB_SYS_MAGIC  = 0xAFFE // full featured BiDiB-Node
	BIDIB_BOOT_MAGIC = 0xB00D // reduced Node, bootloader only

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

	//===============================================================================
	//
	// 3. Type Defines
	//    (useful defines for access to data structures in BiDiB)
	//
	//===============================================================================

	/*typedef struct
	  {
	    union
	      {
	        struct
	          {
	            uint8_t class_switch: 1;
	            uint8_t class_booster: 1;
	            uint8_t class_accessory: 1;
	            uint8_t class_dcc_prog: 1;
	            uint8_t class_dcc_main: 1;
	            uint8_t class_ui: 1;
	            uint8_t class_occupancy: 1;
	            uint8_t class_bridge: 1;
	          };
	        uint8_t class_id;
	      };
	    uint8_t classx_id;
	    uint8_t dcc_vendor;
	    union
	      {
	        struct
	          {
	            uint16_t product_id;
	            uint16_t product_serial;
	          };
	        uint32_t vendor32;
	      };
	  } t_bidib_unique_id;
	*/

	// typedef for control operations - port numbering
	/*typedef union
	  {
	    uint16_t address;
	    struct
	      {
	        uint8_t type;              // BIDIB_PORTTYPE_*
	        union
	          {
	             uint8_t num;          // out number, 0 ... n
	             uint8_t portnum;      // legacy name (do not use)
	          };
	      };
	  } t_bidib_port_idx;
	*/
	// typedef for control operations - execute
	/*typedef struct
	  {
	    union
	      {
	        t_bidib_port_idx port;           // PORT[2]
	        struct                           // legacy names (do not use)
	          {
	            uint8_t type;
	            uint8_t portnum;
	          };
	      };
	    uint8_t portstat;              // state of this output
	  } t_bidib_port;
	*/
	// typedefs for control operations - config (deprecated by configx)
	// note: this structure is deprecated
	/*typedef struct
	    {
	      uint8_t portnum;
	      uint8_t portmode;             // operation mode of port - default state?
	      uint8_t pulstime;             // Holdtime for outputs
	      uint8_t reserved1;            //
	      uint8_t reserved2;            //
	    } t_bidib_switch_cfg;                 // for SWITCH ports

	  typedef t_bidib_switch_cfg t_bidib_sport_cfg; // deprecated name (as of revision 1.24), do not use

	  // note: this structure is deprecated
	  typedef struct
	    {
	      uint8_t portnum;
	      uint8_t brightness_off;       // Brightness in state OFF, range 0..255
	      uint8_t brightness_on;        // Brightness in state ON, range 0..255
	      uint8_t dimm_off;             // time for dimming towards OFF: 0=fast ... 255=slow
	      uint8_t dimm_on;              // time for dimming towards ON: 0=fast ... 255=slow
	    } t_bidib_light_cfg;                  // for LIGHT ports

	  typedef t_bidib_light_cfg t_bidib_lport_cfg; // deprecated name (as of revision 1.24), do not use

	  // note: this structure is deprecated
	  typedef struct
	    {
	      uint8_t portnum;
	      uint8_t dimm_off;             // time for dimming towards OFF: 0=fast ... 255=slow
	      uint8_t dimm_on;              // time for dimming towards ON: 0=fast ... 255=slow
	      uint8_t channel;              // mapping to physical channel
	      uint8_t stretch;              // stretch (make timing slower)
	    } t_bidib_backlight_cfg;              // for BACKLIGHT

	  typedef t_bidib_backlight_cfg t_bidib_backlport_cfg; // deprecated name (as of revision 1.24), do not use

	  // note: this structure is deprecated
	  typedef struct
	    {
	      uint8_t portnum;
	      uint8_t adjust_low;
	      uint8_t adjust_high;
	      uint8_t speed;
	      uint8_t reserved0;
	    } t_bidib_servo_cfg;                  // for Servos

	  // typedefs for accessory operations
	  typedef struct
	    {
	      uint8_t fb_state;             // state of the execution
	      uint8_t error_code;           // 0 or error code
	    } t_bidib_accessory_state;            // for accessory state messages

	  // typedefs for command station operations
	  typedef struct                              //  t_bidib_cs_accessory
	    {
	      union
	        {
	          struct
	            {
	              uint8_t addrl;            // low byte of addr
	              uint8_t addrh;            // high byte of addr
	            };
	          uint16_t addr;                  // true dcc address (start with 0)
	        };
	      union
	        {
	          struct
	            {
	              uint8_t aspect: 5;        // the desired aspect
	              uint8_t activate: 1;      // output (coil) state (only if control_mode == 0)
	              uint8_t control_mode: 1;  // 0: direct coil control; 1:aspect mode
	              uint8_t ext_accessory: 1; // 0: classic dcc; 1:extended accessory control
	            };
	          uint8_t control;
	        };
	      uint8_t time;
	    } t_bidib_cs_accessory;

	  typedef struct                              //  t_bidib_bin_state
	    {
	      union
	        {
	          struct
	            {
	              uint8_t addrl;            // low byte of addr
	              uint8_t addrh;            // high byte of addr
	            };
	          uint16_t addr;                  // true dcc address (start with 0)
	        };
	      union
	        {
	          struct
	            {
	              uint8_t bin_numl;         // low byte of state (this is little endian, DCC is big endian)
	              uint8_t bin_numh;         // high byte of state
	            };
	          uint16_t bin_num;
	        };
	      uint8_t data;
	    } t_bidib_bin_state;

	  typedef struct                              // t_bidib_cs_drive
	    {
	      union
	        {
	          struct
	            {
	              uint8_t addrl;            // low byte of addr
	              uint8_t addrh;            // high byte of addr
	            };
	          uint16_t addr;                  // true dcc address (start with 0)
	        };
	      uint8_t format;                   // BIDIB_CS_DRIVE_FORMAT_DCC14, _DCC28, _DCC128
	      uint8_t active;                   // BIDIB_CS_DRIVE_SPEED_BIT,
	                                              // BIDIB_CS_DRIVE_F1F4_BIT     (1<<1)
	                                              // BIDIB_CS_DRIVE_F5F8_BIT     (1<<2)
	                                              // BIDIB_CS_DRIVE_F9F12_BIT    (1<<3)
	                                              // BIDIB_CS_DRIVE_F13F20_BIT   (1<<4)
	                                              // BIDIB_CS_DRIVE_F21F28_BIT   (1<<5)
	      uint8_t speed;                    // like DCC, MSB=1:forward, MSB=0:revers, speed=1: ESTOP
	      union
	        {
	          struct
	            {
	              uint8_t f4_f1: 4;         // functions f4..f1
	              uint8_t light: 1;         // f0
	            };
	          uint8_t f4_f0;
	        };
	      union
	        {
	          struct
	            {
	              uint8_t f8_f5: 4;         // functions f8..f5
	              uint8_t f12_f9: 4;        // functions f12..f9
	            };
	          uint8_t f12_f5;
	        };
	      union
	        {
	          struct
	            {
	              uint8_t f16_f13: 4;       // functions f16..f13
	              uint8_t f20_f17: 4;       // functions f20..f17
	            };
	          uint8_t f20_f13;
	        };
	      union
	        {
	          struct
	            {
	              uint8_t f24_f21: 4;       // functions f24..f21
	              uint8_t f28_f25: 4;       // functions f28..f25
	            };
	          uint8_t f28_f21;
	        };
	    } t_bidib_cs_drive;

	  // typedefs for command station operations - pom and service mode
	  typedef struct                              // t_bidib_cs_pom
	    {
	      union
	        {
	          struct
	            {
	              union
	                {
	                  struct
	                    {
	                      uint8_t addrl;    // low byte of addr
	                      uint8_t addrh;    // high byte of addr
	                    };
	                  uint16_t addr;          // true dcc address (start with 0)
	                };
	              uint8_t addrxl;           // 0 for normal POM
	              uint8_t addrxh;           // 0 for normal POM
	              uint8_t mid;              // manufactorer ID: 0 for normal POM, else VendorID like DCC
	            };
	          uint8_t did[5];               // true dcc address (start with 0)
	        };
	      uint8_t opcode;                   // 0=RdBlock, 1=RdByte, 2=WrBit, 3=WrByte
	                                              // 80=reserved, 81=XRdBlock, 82=XWrBit, 83=XWrByte
	                                              // see below: BIDIB_CS_POM_RD_BLOCK ...
	      union
	        {
	          struct
	            {
	              uint8_t cv_addrl;         // low byte of cv addr
	              uint8_t cv_addrh;         // high byte of cv addr
	            };
	          uint16_t cv_addr;               // true cv address (start with 0)
	        };
	      uint8_t cv_addrx;                 //
	      uint8_t data[4];
	    } t_bidib_cs_pom;

	  typedef struct                              // t_bidib_cs_prog
	    {
	      uint8_t opcode;                   // 0=Break, 1=RdByte, 2=WrBit, 3=WrByte
	      union
	        {
	          struct
	            {
	              uint8_t cv_addrl;         // low byte of cv addr
	              uint8_t cv_addrh;         // high byte of cv addr
	            };
	          uint16_t cv_addr;               // true cv address (start with 0)
	        };
	      uint8_t data;
	    } t_bidib_cs_prog;

	  typedef struct                              // t_bidib_cs_prog_state
	    {
	      uint8_t result;                   //
	      uint8_t time;                     //
	      union
	        {
	          struct
	            {
	              uint8_t cv_addrl;         // low byte of cv addr
	              uint8_t cv_addrh;         // high byte of cv addr
	            };
	          uint16_t cv_addr;               // true cv address (start with 0)
	        };
	      uint8_t data;
	    } t_bidib_cs_prog_state;
	*/
	//===============================================================================
	//
	// 4. Feature Codes
	//
	//===============================================================================

	//-- occupancy
	FEATURE_BM_SIZE                  FeatureID = 0 // number of occupancy detectors
	FEATURE_BM_ON                    FeatureID = 1 // occupancy detection on/off
	FEATURE_BM_SECACK_AVAILABLE      FeatureID = 2 // secure ack available
	FEATURE_BM_SECACK_ON             FeatureID = 3 // secure ack on/off
	FEATURE_BM_CURMEAS_AVAILABLE     FeatureID = 4 // occupancy detectors with current measurement results
	FEATURE_BM_CURMEAS_INTERVAL      FeatureID = 5 // interval for current measurements
	FEATURE_BM_DC_MEAS_AVAILABLE     FeatureID = 6 // (dc) measurements available, even if track voltage is off
	FEATURE_BM_DC_MEAS_ON            FeatureID = 7 // dc measurement enabled
	FEATURE_BM_ADDR_DETECT_AVAILABLE FeatureID = 8 // detector ic capable to detect loco address
	FEATURE_BM_ADDR_DETECT_ON        FeatureID = 9 // address detection enabled
	//-- bidi detection
	FEATURE_BM_ADDR_AND_DIR       FeatureID = 10 // addresses contain direction
	FEATURE_BM_ISTSPEED_AVAILABLE FeatureID = 11 // speed messages available
	FEATURE_BM_ISTSPEED_INTERVAL  FeatureID = 12 // speed update interval
	FEATURE_BM_CV_AVAILABLE       FeatureID = 13 // CV readback available
	FEATURE_BM_CV_ON              FeatureID = 14 // CV readback enabled
	//-- booster
	FEATURE_BST_VOLT_ADJUSTABLE     FeatureID = 15 // booster output voltage is adjustable
	FEATURE_BST_VOLT                FeatureID = 16 // booster output voltage setting (unit: V)
	FEATURE_BST_CUTOUT_AVAIALABLE   FeatureID = 17 // booster can do cutout for railcom
	FEATURE_BST_CUTOUT_ON           FeatureID = 18 // cutout is enabled
	FEATURE_BST_TURNOFF_TIME        FeatureID = 19 // time in ms until booster turns off in case of a short (unit 2ms)
	FEATURE_BST_INRUSH_TURNOFF_TIME FeatureID = 20 // time in ms until booster turns off in case of a short after the first power up (unit 2ms)
	FEATURE_BST_AMPERE_ADJUSTABLE   FeatureID = 21 // booster output current is adjustable
	FEATURE_BST_AMPERE              FeatureID = 22 // booster output current value (special coding)
	FEATURE_BST_CURMEAS_INTERVAL    FeatureID = 23 // current update interval
	FEATURE_BST_CV_AVAILABLE        FeatureID = 24 // (deprecated, now synonym to 13) CV readback available
	FEATURE_BST_CV_ON               FeatureID = 25 // (deprecated, now synonym to 14) CV readback enabled
	FEATURE_BST_INHIBIT_AUTOSTART   FeatureID = 26 // 1: Booster does no automatic BOOST_ON when DCC at input wakes up.
	FEATURE_BST_INHIBIT_LOCAL_ONOFF FeatureID = 27 // 1: Booster announces local STOP/GO key stroke only, no local action

	//-- bidi detection
	FEATURE_BM_DYN_STATE_INTERVAL FeatureID = 28 // transmit interval of MSG_BM_DYN_STATE (unit 100ms)
	FEATURE_BM_RCPLUS_AVAILABLE   FeatureID = 29 // 1: RailcomPlus messages available
	//-- occupancy
	FEATURE_BM_TIMESTAMP_ON FeatureID = 30 // 1: OCC will be sent with timestamp
	//-- bidi detection
	FEATURE_BM_POSITION_ON     FeatureID = 31 // position messages enabled
	FEATURE_BM_POSITION_SECACK FeatureID = 32 // secure position ack interval (unit: 10ms), 0: none

	//-- accessory
	FEATURE_ACCESSORY_COUNT       FeatureID = 40 // number of objects
	FEATURE_ACCESSORY_SURVEILLED  FeatureID = 41 // 1: announce if operated outside bidib
	FEATURE_ACCESSORY_MACROMAPPED FeatureID = 42 // 1..n: no of accessory aspects are mapped to macros

	//-- control
	FEATURE_CTRL_INPUT_COUNT              FeatureID = 50 // number of inputs for keys
	FEATURE_CTRL_INPUT_NOTIFY             FeatureID = 51 // 1: report a keystroke to host
	FEATURE_CTRL_SWITCH_COUNT             FeatureID = 52 // number of switch ports (direct controlled)
	FEATURE_CTRL_LIGHT_COUNT              FeatureID = 53 // number of light ports (direct controlled)
	FEATURE_CTRL_SERVO_COUNT              FeatureID = 54 // number of servo ports (direct controlled)
	FEATURE_CTRL_SOUND_COUNT              FeatureID = 55 // number of sound ports (direct controlled)
	FEATURE_CTRL_MOTOR_COUNT              FeatureID = 56 // number of motor ports (direct controlled)
	FEATURE_CTRL_ANALOGOUT_COUNT          FeatureID = 57 // number of analog ports (direct controlled)
	FEATURE_CTRL_STRETCH_DIMM             FeatureID = 58 // additional time stretch for dimming (for light ports)
	FEATURE_CTRL_BACKLIGHT_COUNT          FeatureID = 59 // number of backlight ports (intensity direct controlled)
	FEATURE_CTRL_MAC_LEVEL                FeatureID = 60 // supported macro level
	FEATURE_CTRL_MAC_SAVE                 FeatureID = 61 // number of permanent storage places for macros
	FEATURE_CTRL_MAC_COUNT                FeatureID = 62 // number of macros
	FEATURE_CTRL_MAC_SIZE                 FeatureID = 63 // length of each macro (entries)
	FEATURE_CTRL_MAC_START_MAN            FeatureID = 64 // (local) manual control of macros enabled
	FEATURE_CTRL_MAC_START_DCC            FeatureID = 65 // (local) dcc control of macros enabled
	FEATURE_CTRL_PORT_QUERY_AVAILABLE     FeatureID = 66 // 1: node will answer to output queries via MSG_LC_PORT_QUERY
	FEATURE_SWITCH_CONFIG_AVAILABLE       FeatureID = 67 // (deprecated, version >= 0.6 uses availability of PCFG_IO_CTRL) 1: node has possibility to configure switch ports
	FEATURE_CTRL_PORT_FLAT_MODEL          FeatureID = 70 // node uses flat port model, "low" number of addressable ports
	FEATURE_CTRL_PORT_FLAT_MODEL_EXTENDED FeatureID = 71 // node uses flat port model, "high" number of addressable ports
	/* deprecated names (as of revision 1.24), do not use */
	FEATURE_CTRL_SPORT_COUNT       FeatureID = 52 // (deprecated)
	FEATURE_CTRL_LPORT_COUNT       FeatureID = 53 // (deprecated)
	FEATURE_CTRL_ANALOG_COUNT      FeatureID = 57 // (deprecated)
	FEATURE_SPORT_CONFIG_AVAILABLE FeatureID = 67 // (deprecated)

	//-- dcc gen
	FEATURE_GEN_SPYMODE             FeatureID = 100 // 1: watch bidib handsets
	FEATURE_GEN_WATCHDOG            FeatureID = 101 // 0: no watchdog, 1: permanent update of MSG_CS_SET_STATE required, unit 100ms
	FEATURE_GEN_DRIVE_ACK           FeatureID = 102 // not used
	FEATURE_GEN_SWITCH_ACK          FeatureID = 103 // not used
	FEATURE_GEN_LOK_DB_SIZE         FeatureID = 104 //
	FEATURE_GEN_LOK_DB_STRING       FeatureID = 105 //
	FEATURE_GEN_POM_REPEAT          FeatureID = 106 // supported service modes
	FEATURE_GEN_DRIVE_BUS           FeatureID = 107 // 1: this node drive the dcc bus.
	FEATURE_GEN_LOK_LOST_DETECT     FeatureID = 108 // 1: command station annouces lost loco
	FEATURE_GEN_NOTIFY_DRIVE_MANUAL FeatureID = 109 // 1: dcc gen reports manual operation
	FEATURE_GEN_START_STATE         FeatureID = 110 // 1: power up state, 0=off, 1=on
	FEATURE_GEN_RCPLUS_AVAILABLE    FeatureID = 111 // 1: supports rcplus messages

	FEATURE_STRING_SIZE       FeatureID = 252 // length of user strings, 0:n.a (default); allowed 8..24
	FEATURE_RELEVANT_PID_BITS FeatureID = 253 // how many bits of 'vendor32' are relevant for PID (default 16, LSB aligned)
	FEATURE_FW_UPDATE_MODE    FeatureID = 254 // 0: no fw-update, 1: intel hex (max. 10 byte / record)
	FEATURE_EXTENSION         FeatureID = 255 // 1: reserved for future expansion

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

	BIDIB_CS_STATE_OFF       CsState = 0x00 // no DCC, DCC-line is static, not toggling
	BIDIB_CS_STATE_STOP      CsState = 0x01 // DCC, all speed setting = 0
	BIDIB_CS_STATE_SOFTSTOP  CsState = 0x02 // DCC, soft stop is progress
	BIDIB_CS_STATE_GO        CsState = 0x03 // DCC on (must be repeated if watchdog is on)
	BIDIB_CS_STATE_GO_IGN_WD CsState = 0x04 // DCC on (watchdog ignored)
	BIDIB_CS_STATE_PROG      CsState = 0x08 // in Programming Mode (ready for commands)
	BIDIB_CS_STATE_PROGBUSY  CsState = 0x09 // in Programming Mode (busy)
	BIDIB_CS_STATE_BUSY      CsState = 0x0D // busy
	BIDIB_CS_STATE_QUERY     CsState = 0xFF

	BIDIB_CS_DRIVE_FORMAT_DCC14  DccFormat = 0
	BIDIB_CS_DRIVE_FORMAT_DCC28  DccFormat = 2
	BIDIB_CS_DRIVE_FORMAT_DCC128 DccFormat = 3

	BIDIB_CS_DRIVE_SPEED_BIT  = (1 << 0)
	BIDIB_CS_DRIVE_F1F4_BIT   = (1 << 1) // also FL
	BIDIB_CS_DRIVE_F0F4_BIT   = (1 << 1) // additional define, it is the same bit
	BIDIB_CS_DRIVE_F5F8_BIT   = (1 << 2)
	BIDIB_CS_DRIVE_F9F12_BIT  = (1 << 3)
	BIDIB_CS_DRIVE_F13F20_BIT = (1 << 4)
	BIDIB_CS_DRIVE_F21F28_BIT = (1 << 5)

	BIDIB_CS_POM_RD_BLOCK  = 0 // bit 0,1: CC-Bits
	BIDIB_CS_POM_RD_BYTE   = 1 // bit 2,3: no. of bytes to write (-1)
	BIDIB_CS_POM_WR_BIT    = 2 // bit 6,7: standard pom/short form/xpom/reserved
	BIDIB_CS_POM_WR_BYTE   = 3
	BIDIB_CS_XWR_BYTE1     = 0x43
	BIDIB_CS_XWR_BYTE2     = 0x47
	BIDIB_CS_xPOM_reserved = 0x80
	BIDIB_CS_xPOM_RD_BLOCK = 0x81
	BIDIB_CS_xPOM_WR_BIT   = 0x82
	BIDIB_CS_xPOM_WR_BYTE1 = 0x83
	BIDIB_CS_xPOM_WR_BYTE2 = 0x87
	BIDIB_CS_xPOM_WR_BYTE3 = 0x8B
	BIDIB_CS_xPOM_WR_BYTE4 = 0x8F

	BIDIB_CS_PROG_BREAK    = 0 // service mode commands (MSG_CS_PROG)
	BIDIB_CS_PROG_QUERY    = 1
	BIDIB_CS_PROG_RD_BYTE  = 2
	BIDIB_CS_PROG_RDWR_BIT = 3
	BIDIB_CS_PROG_WR_BYTE  = 4

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
