package tlsspecloader

import "fmt"

type Go struct {
}

func (*Go) MutualMapUint8(a []uint8, b []string) string {
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
		m1 += fmt.Sprintf("\t0x%02X: %q,\n", a[i], b[i])
	}
	m1 += "}\n"

	// second map: string -> uint8
	var m2 string
	m2 += "var StringToUint8 = map[string]uint8{\n"
	for i := 0; i < maxLen; i++ {
		m2 += fmt.Sprintf("\t%q: 0x%02X,\n", b[i], a[i])
	}
	m2 += "}\n"

	return m1 + m2
}

func (*Go) MutualMapUint16(a []uint16, b []string) string {
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
		m1 += fmt.Sprintf("\t0x%04X: %q,\n", a[i], b[i])
	}
	m1 += "}\n"

	// second map: string -> uint16
	var m2 string
	m2 += "var StringToUint16 = map[string]uint16{\n"
	for i := 0; i < maxLen; i++ {
		m2 += fmt.Sprintf("\t%q: 0x%04X,\n", b[i], a[i])
	}
	m2 += "}\n"

	return m1 + m2
}
