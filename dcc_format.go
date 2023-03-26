package bidib

// DCC Format indicator
type DccFormat uint8

const (
	BIDIB_CS_DRIVE_FORMAT_DCC14  DccFormat = 0
	BIDIB_CS_DRIVE_FORMAT_DCC28  DccFormat = 2
	BIDIB_CS_DRIVE_FORMAT_DCC128 DccFormat = 3
)
