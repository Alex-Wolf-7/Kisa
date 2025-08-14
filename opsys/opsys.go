package opsys

type OpSys struct {
	opSys opSysEnum
	name  string
}

type opSysEnum int

const (
	none opSysEnum = iota
	windows
	mac
)

func NewOpSys(goos string) OpSys {
	if goos == "windows" {
		return OpSys{opSys: windows, name: goos}
	} else {
		return OpSys{opSys: mac, name: goos}
	}
}

func (opSys OpSys) IsWindows() bool {
	return opSys.opSys == windows
}

func (opSys OpSys) IsMac() bool {
	return opSys.opSys == mac
}

func (OpSys OpSys) String() string {
	return OpSys.name
}
