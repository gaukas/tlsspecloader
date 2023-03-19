package tlsspecloader

import "fmt"

type Go struct {
	HexOutput bool
}

func (g *Go) MutualMapUint8(a []uint8, b []string) (aIndexB string, bIndexa string) {
	maxLen := len(a)
	if len(a) != len(b) {
		if len(a) < len(b) {
			fmt.Printf("a is shorter than b, maxLen is set to len(a)")
		} else {
			fmt.Printf("a is longer than b, maxLen is set to len(b)")
			maxLen = len(b)
		}
	}

	// first map: uint8 -> string
	var m1 string
	m1 += "var Uint8MapToString = map[uint8]string{\n"
	for i := 0; i < maxLen; i++ {
		if g.HexOutput {
			m1 += fmt.Sprintf("\t0x%02X: %q,\n", a[i], b[i])
		} else {
			m1 += fmt.Sprintf("\t%d: %q,\n", a[i], b[i])
		}
	}
	m1 += "}\n"

	// second map: string -> uint8
	var m2 string
	m2 += "var StringToUint8 = map[string]uint8{\n"
	for i := 0; i < maxLen; i++ {
		if g.HexOutput {
			m2 += fmt.Sprintf("\t%q: 0x%02X,\n", b[i], a[i])
		} else {
			m2 += fmt.Sprintf("\t%q: %d,\n", b[i], a[i])
		}
	}
	m2 += "}\n"

	return m1, m2
}

func (g *Go) MutualMapUint16(a []uint16, b []string) (aIndexB string, bIndexa string) {
	maxLen := len(a)
	if len(a) != len(b) {
		if len(a) < len(b) {
			fmt.Printf("a is shorter than b, maxLen is set to len(a)")
		} else {
			fmt.Printf("a is longer than b, maxLen is set to len(b)")
			maxLen = len(b)
		}
	}

	// first map: uint16 -> string
	var m1 string
	m1 += "var Uint16MapToString = map[uint16]string{\n"
	for i := 0; i < maxLen; i++ {
		if g.HexOutput {
			m1 += fmt.Sprintf("\t0x%04X: %q,\n", a[i], b[i])
		} else {
			m1 += fmt.Sprintf("\t%d: %q,\n", a[i], b[i])
		}
	}
	m1 += "}\n"

	// second map: string -> uint16
	var m2 string
	m2 += "var StringToUint16 = map[string]uint16{\n"
	for i := 0; i < maxLen; i++ {
		if g.HexOutput {
			m2 += fmt.Sprintf("\t%q: 0x%04X,\n", b[i], a[i])
		} else {
			m2 += fmt.Sprintf("\t%q: %d,\n", b[i], a[i])
		}
	}
	m2 += "}\n"

	return m1, m2
}

func (g *Go) EnumUint8(value []uint8, label []string) string {
	var enum string
	enum += "const (\n"
	for i := 0; i < len(value); i++ {
		if g.HexOutput {
			enum += fmt.Sprintf("\t%s uint8 = 0x%02X\n", label[i], value[i])
		} else {
			enum += fmt.Sprintf("\t%s uint8 = %d\n", label[i], value[i])
		}
	}
	enum += ")\n"

	return enum
}

func (g *Go) EnumUint16(value []uint16, label []string) string {
	var enum string
	enum += "const (\n"
	for i := 0; i < len(value); i++ {
		if g.HexOutput {
			enum += fmt.Sprintf("\t%s uint16 = 0x%04X\n", label[i], value[i])
		} else {
			enum += fmt.Sprintf("\t%s uint16 = %d\n", label[i], value[i])
		}
	}
	enum += ")\n"

	return enum
}
