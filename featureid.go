package bidib

// Identifier for feature number
//
//go:generate stringer -type=FeatureID
type FeatureID uint8

const (
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
)
