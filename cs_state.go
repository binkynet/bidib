package bidib

// CS State flags
type CsState uint8

const (
	BIDIB_CS_STATE_OFF       CsState = 0x00 // no DCC, DCC-line is static, not toggling
	BIDIB_CS_STATE_STOP      CsState = 0x01 // DCC, all speed setting = 0
	BIDIB_CS_STATE_SOFTSTOP  CsState = 0x02 // DCC, soft stop is progress
	BIDIB_CS_STATE_GO        CsState = 0x03 // DCC on (must be repeated if watchdog is on)
	BIDIB_CS_STATE_GO_IGN_WD CsState = 0x04 // DCC on (watchdog ignored)
	BIDIB_CS_STATE_PROG      CsState = 0x08 // in Programming Mode (ready for commands)
	BIDIB_CS_STATE_PROGBUSY  CsState = 0x09 // in Programming Mode (busy)
	BIDIB_CS_STATE_BUSY      CsState = 0x0D // busy
	BIDIB_CS_STATE_QUERY     CsState = 0xFF
)
