package binp

// Printer type. Don't touch the internals.
type Printer struct {
	W []byte
}

// Create a new printer with empty output.
func NewPrinter() *Printer {
	return &Printer{[]byte{}}
}

// Create a new printer with output prefixed with the given byte slice.
func NewPrinterWith(b []byte) *Printer {
	return &Printer{b}
}

// Output a byte.
func (p *Printer) Byte(d byte) *Printer {
	p.W = append(p.W, d)
	return p
}

// Output 2 bigendian bytes.
func (p *Printer) U16(d uint16) *Printer {
	p.W = append(p.W, byte(d>>8), byte(d))
	return p
}

// Output 4 bigendian bytes.
func (p *Printer) U32(d uint32) *Printer {
	p.W = append(p.W, byte(d>>24), byte(d>>16), byte(d>>8), byte(d))
	return p
}

// Output 4 bigendian bytes.
func (p *Printer) U64(d uint64) *Printer {
	p.W = append(p.W, byte(d>>56), byte(d>>48), byte(d>>40), byte(d>>32), byte(d>>24), byte(d>>16), byte(d>>8), byte(d))
	return p
}

// Output a raw byte slice with no length prefix.
func (p *Printer) Bytes(d []byte) *Printer {
	p.W = append(p.W, d...)
	return p
}

// Output a raw string with no length prefix.
func (p *Printer) String(d string) *Printer {
	p.W = append(p.W, []byte(d)...)
	return p
}

// Output a string with a 4 byte bigendian length prefix and no trailing null.
func (p *Printer) U32String(d string) *Printer {
	return p.U32(uint32(len(d))).String(d)
}

// Output bytes with a 4 byte bigendian length prefix and no trailing null.
func (p *Printer) U32Bytes(d []byte) *Printer {
	return p.U32(uint32(len(d))).Bytes(d)
}

// Output a string with a 2 byte bigendian length prefix and no trailing null.
func (p *Printer) U16String(d string) *Printer {
	if len(d) > 0xffff {
		panic("binprinter: string too long")
	}
	return p.U16(uint16(len(d))).String(d)
}

// Output a string with a 1 byte bigendian length prefix and no trailing null.
func (p *Printer) U8String(d string) *Printer {
	if len(d) > 0xff {
		panic("binprinter: string too long")
	}
	return p.Byte(byte(len(d))).String(d)
}

// Output a string terminated by a null-byte
func (p *Printer)String0(d string) *Printer {
	return p.String(d).Byte(0)
}

// Get the output as a byte slice.
func (p *Printer) Out() []byte {
	return p.W
}
