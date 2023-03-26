package bidib

// Command station POM OpCode
type CsPomOpCode uint8

const (
	BIDIB_CS_POM_RD_BLOCK  CsPomOpCode = 0 // bit 0,1: CC-Bits
	BIDIB_CS_POM_RD_BYTE   CsPomOpCode = 1 // bit 2,3: no. of bytes to write (-1)
	BIDIB_CS_POM_WR_BIT    CsPomOpCode = 2 // bit 6,7: standard pom/short form/xpom/reserved
	BIDIB_CS_POM_WR_BYTE   CsPomOpCode = 3
	BIDIB_CS_XWR_BYTE1     CsPomOpCode = 0x43
	BIDIB_CS_XWR_BYTE2     CsPomOpCode = 0x47
	BIDIB_CS_xPOM_reserved CsPomOpCode = 0x80
	BIDIB_CS_xPOM_RD_BLOCK CsPomOpCode = 0x81
	BIDIB_CS_xPOM_WR_BIT   CsPomOpCode = 0x82
	BIDIB_CS_xPOM_WR_BYTE1 CsPomOpCode = 0x83
	BIDIB_CS_xPOM_WR_BYTE2 CsPomOpCode = 0x87
	BIDIB_CS_xPOM_WR_BYTE3 CsPomOpCode = 0x8B
	BIDIB_CS_xPOM_WR_BYTE4 CsPomOpCode = 0x8F
)

// Command station Prog  OpCode
type CsProgOpCode uint8

const (
	BIDIB_CS_PROG_BREAK    CsProgOpCode = 0 // service mode commands (MSG_CS_PROG)
	BIDIB_CS_PROG_QUERY    CsProgOpCode = 1
	BIDIB_CS_PROG_RD_BYTE  CsProgOpCode = 2
	BIDIB_CS_PROG_RDWR_BIT CsProgOpCode = 3
	BIDIB_CS_PROG_WR_BYTE  CsProgOpCode = 4
)
