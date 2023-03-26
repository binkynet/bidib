// Code generated by "stringer -type=FeatureID"; DO NOT EDIT.

package bidib

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FEATURE_BM_SIZE-0]
	_ = x[FEATURE_BM_ON-1]
	_ = x[FEATURE_BM_SECACK_AVAILABLE-2]
	_ = x[FEATURE_BM_SECACK_ON-3]
	_ = x[FEATURE_BM_CURMEAS_AVAILABLE-4]
	_ = x[FEATURE_BM_CURMEAS_INTERVAL-5]
	_ = x[FEATURE_BM_DC_MEAS_AVAILABLE-6]
	_ = x[FEATURE_BM_DC_MEAS_ON-7]
	_ = x[FEATURE_BM_ADDR_DETECT_AVAILABLE-8]
	_ = x[FEATURE_BM_ADDR_DETECT_ON-9]
	_ = x[FEATURE_BM_ADDR_AND_DIR-10]
	_ = x[FEATURE_BM_ISTSPEED_AVAILABLE-11]
	_ = x[FEATURE_BM_ISTSPEED_INTERVAL-12]
	_ = x[FEATURE_BM_CV_AVAILABLE-13]
	_ = x[FEATURE_BM_CV_ON-14]
	_ = x[FEATURE_BST_VOLT_ADJUSTABLE-15]
	_ = x[FEATURE_BST_VOLT-16]
	_ = x[FEATURE_BST_CUTOUT_AVAIALABLE-17]
	_ = x[FEATURE_BST_CUTOUT_ON-18]
	_ = x[FEATURE_BST_TURNOFF_TIME-19]
	_ = x[FEATURE_BST_INRUSH_TURNOFF_TIME-20]
	_ = x[FEATURE_BST_AMPERE_ADJUSTABLE-21]
	_ = x[FEATURE_BST_AMPERE-22]
	_ = x[FEATURE_BST_CURMEAS_INTERVAL-23]
	_ = x[FEATURE_BST_CV_AVAILABLE-24]
	_ = x[FEATURE_BST_CV_ON-25]
	_ = x[FEATURE_BST_INHIBIT_AUTOSTART-26]
	_ = x[FEATURE_BST_INHIBIT_LOCAL_ONOFF-27]
	_ = x[FEATURE_BM_DYN_STATE_INTERVAL-28]
	_ = x[FEATURE_BM_RCPLUS_AVAILABLE-29]
	_ = x[FEATURE_BM_TIMESTAMP_ON-30]
	_ = x[FEATURE_BM_POSITION_ON-31]
	_ = x[FEATURE_BM_POSITION_SECACK-32]
	_ = x[FEATURE_ACCESSORY_COUNT-40]
	_ = x[FEATURE_ACCESSORY_SURVEILLED-41]
	_ = x[FEATURE_ACCESSORY_MACROMAPPED-42]
	_ = x[FEATURE_CTRL_INPUT_COUNT-50]
	_ = x[FEATURE_CTRL_INPUT_NOTIFY-51]
	_ = x[FEATURE_CTRL_SWITCH_COUNT-52]
	_ = x[FEATURE_CTRL_LIGHT_COUNT-53]
	_ = x[FEATURE_CTRL_SERVO_COUNT-54]
	_ = x[FEATURE_CTRL_SOUND_COUNT-55]
	_ = x[FEATURE_CTRL_MOTOR_COUNT-56]
	_ = x[FEATURE_CTRL_ANALOGOUT_COUNT-57]
	_ = x[FEATURE_CTRL_STRETCH_DIMM-58]
	_ = x[FEATURE_CTRL_BACKLIGHT_COUNT-59]
	_ = x[FEATURE_CTRL_MAC_LEVEL-60]
	_ = x[FEATURE_CTRL_MAC_SAVE-61]
	_ = x[FEATURE_CTRL_MAC_COUNT-62]
	_ = x[FEATURE_CTRL_MAC_SIZE-63]
	_ = x[FEATURE_CTRL_MAC_START_MAN-64]
	_ = x[FEATURE_CTRL_MAC_START_DCC-65]
	_ = x[FEATURE_CTRL_PORT_QUERY_AVAILABLE-66]
	_ = x[FEATURE_SWITCH_CONFIG_AVAILABLE-67]
	_ = x[FEATURE_CTRL_PORT_FLAT_MODEL-70]
	_ = x[FEATURE_CTRL_PORT_FLAT_MODEL_EXTENDED-71]
	_ = x[FEATURE_CTRL_SPORT_COUNT-52]
	_ = x[FEATURE_CTRL_LPORT_COUNT-53]
	_ = x[FEATURE_CTRL_ANALOG_COUNT-57]
	_ = x[FEATURE_SPORT_CONFIG_AVAILABLE-67]
	_ = x[FEATURE_GEN_SPYMODE-100]
	_ = x[FEATURE_GEN_WATCHDOG-101]
	_ = x[FEATURE_GEN_DRIVE_ACK-102]
	_ = x[FEATURE_GEN_SWITCH_ACK-103]
	_ = x[FEATURE_GEN_LOK_DB_SIZE-104]
	_ = x[FEATURE_GEN_LOK_DB_STRING-105]
	_ = x[FEATURE_GEN_POM_REPEAT-106]
	_ = x[FEATURE_GEN_DRIVE_BUS-107]
	_ = x[FEATURE_GEN_LOK_LOST_DETECT-108]
	_ = x[FEATURE_GEN_NOTIFY_DRIVE_MANUAL-109]
	_ = x[FEATURE_GEN_START_STATE-110]
	_ = x[FEATURE_GEN_RCPLUS_AVAILABLE-111]
	_ = x[FEATURE_STRING_SIZE-252]
	_ = x[FEATURE_RELEVANT_PID_BITS-253]
	_ = x[FEATURE_FW_UPDATE_MODE-254]
	_ = x[FEATURE_EXTENSION-255]
}

const (
	_FeatureID_name_0 = "FEATURE_BM_SIZEFEATURE_BM_ONFEATURE_BM_SECACK_AVAILABLEFEATURE_BM_SECACK_ONFEATURE_BM_CURMEAS_AVAILABLEFEATURE_BM_CURMEAS_INTERVALFEATURE_BM_DC_MEAS_AVAILABLEFEATURE_BM_DC_MEAS_ONFEATURE_BM_ADDR_DETECT_AVAILABLEFEATURE_BM_ADDR_DETECT_ONFEATURE_BM_ADDR_AND_DIRFEATURE_BM_ISTSPEED_AVAILABLEFEATURE_BM_ISTSPEED_INTERVALFEATURE_BM_CV_AVAILABLEFEATURE_BM_CV_ONFEATURE_BST_VOLT_ADJUSTABLEFEATURE_BST_VOLTFEATURE_BST_CUTOUT_AVAIALABLEFEATURE_BST_CUTOUT_ONFEATURE_BST_TURNOFF_TIMEFEATURE_BST_INRUSH_TURNOFF_TIMEFEATURE_BST_AMPERE_ADJUSTABLEFEATURE_BST_AMPEREFEATURE_BST_CURMEAS_INTERVALFEATURE_BST_CV_AVAILABLEFEATURE_BST_CV_ONFEATURE_BST_INHIBIT_AUTOSTARTFEATURE_BST_INHIBIT_LOCAL_ONOFFFEATURE_BM_DYN_STATE_INTERVALFEATURE_BM_RCPLUS_AVAILABLEFEATURE_BM_TIMESTAMP_ONFEATURE_BM_POSITION_ONFEATURE_BM_POSITION_SECACK"
	_FeatureID_name_1 = "FEATURE_ACCESSORY_COUNTFEATURE_ACCESSORY_SURVEILLEDFEATURE_ACCESSORY_MACROMAPPED"
	_FeatureID_name_2 = "FEATURE_CTRL_INPUT_COUNTFEATURE_CTRL_INPUT_NOTIFYFEATURE_CTRL_SWITCH_COUNTFEATURE_CTRL_LIGHT_COUNTFEATURE_CTRL_SERVO_COUNTFEATURE_CTRL_SOUND_COUNTFEATURE_CTRL_MOTOR_COUNTFEATURE_CTRL_ANALOGOUT_COUNTFEATURE_CTRL_STRETCH_DIMMFEATURE_CTRL_BACKLIGHT_COUNTFEATURE_CTRL_MAC_LEVELFEATURE_CTRL_MAC_SAVEFEATURE_CTRL_MAC_COUNTFEATURE_CTRL_MAC_SIZEFEATURE_CTRL_MAC_START_MANFEATURE_CTRL_MAC_START_DCCFEATURE_CTRL_PORT_QUERY_AVAILABLEFEATURE_SWITCH_CONFIG_AVAILABLE"
	_FeatureID_name_3 = "FEATURE_CTRL_PORT_FLAT_MODELFEATURE_CTRL_PORT_FLAT_MODEL_EXTENDED"
	_FeatureID_name_4 = "FEATURE_GEN_SPYMODEFEATURE_GEN_WATCHDOGFEATURE_GEN_DRIVE_ACKFEATURE_GEN_SWITCH_ACKFEATURE_GEN_LOK_DB_SIZEFEATURE_GEN_LOK_DB_STRINGFEATURE_GEN_POM_REPEATFEATURE_GEN_DRIVE_BUSFEATURE_GEN_LOK_LOST_DETECTFEATURE_GEN_NOTIFY_DRIVE_MANUALFEATURE_GEN_START_STATEFEATURE_GEN_RCPLUS_AVAILABLE"
	_FeatureID_name_5 = "FEATURE_STRING_SIZEFEATURE_RELEVANT_PID_BITSFEATURE_FW_UPDATE_MODEFEATURE_EXTENSION"
)

var (
	_FeatureID_index_0 = [...]uint16{0, 15, 28, 55, 75, 103, 130, 158, 179, 211, 236, 259, 288, 316, 339, 355, 382, 398, 427, 448, 472, 503, 532, 550, 578, 602, 619, 648, 679, 708, 735, 758, 780, 806}
	_FeatureID_index_1 = [...]uint8{0, 23, 51, 80}
	_FeatureID_index_2 = [...]uint16{0, 24, 49, 74, 98, 122, 146, 170, 198, 223, 251, 273, 294, 316, 337, 363, 389, 422, 453}
	_FeatureID_index_3 = [...]uint8{0, 28, 65}
	_FeatureID_index_4 = [...]uint16{0, 19, 39, 60, 82, 105, 130, 152, 173, 200, 231, 254, 282}
	_FeatureID_index_5 = [...]uint8{0, 19, 44, 66, 83}
)

func (i FeatureID) String() string {
	switch {
	case i <= 32:
		return _FeatureID_name_0[_FeatureID_index_0[i]:_FeatureID_index_0[i+1]]
	case 40 <= i && i <= 42:
		i -= 40
		return _FeatureID_name_1[_FeatureID_index_1[i]:_FeatureID_index_1[i+1]]
	case 50 <= i && i <= 67:
		i -= 50
		return _FeatureID_name_2[_FeatureID_index_2[i]:_FeatureID_index_2[i+1]]
	case 70 <= i && i <= 71:
		i -= 70
		return _FeatureID_name_3[_FeatureID_index_3[i]:_FeatureID_index_3[i+1]]
	case 100 <= i && i <= 111:
		i -= 100
		return _FeatureID_name_4[_FeatureID_index_4[i]:_FeatureID_index_4[i+1]]
	case 252 <= i && i <= 255:
		i -= 252
		return _FeatureID_name_5[_FeatureID_index_5[i]:_FeatureID_index_5[i+1]]
	default:
		return "FeatureID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}