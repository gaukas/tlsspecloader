package tlsspecloader

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToUint8(u string) (uint8, error) {
	if !strings.HasPrefix(u, "0x") {
		return 0, fmt.Errorf("invalid string %s", u)
	}

	u = strings.TrimPrefix(u, "0x")
	if len(u) != 2 {
		return 0, fmt.Errorf("invalid string %s", u)
	}

	v, err := strconv.ParseUint(u, 16, 8)
	if err != nil {
		return 0, err
	}

	return uint8(v), nil
}

func ArrStringToUint8(in []string) []uint8 {
	out := make([]uint8, len(in))
	for i := 0; i < len(in); i++ {
		u8, err := StringToUint8(in[i])
		if err != nil {
			continue
		}
		out[i] = u8
	}

	return out
}

// StringUint8ToUint16 converts a string in the format of
// "0x0A, 0x0B" to uint16 0x0A0B
func StringUint8ToUint16(u string) (uint16, error) {
	split := strings.Split(u, ",")
	if len(split) != 2 {
		return 0, fmt.Errorf("invalid string %s", u)
	}

	u1, err := StringToUint8(strings.TrimSpace(split[0]))
	if err != nil {
		return 0, err
	}

	u2, err := StringToUint8(strings.TrimSpace(split[1]))
	if err != nil {
		return 0, err
	}

	return uint16(u1)<<8 | uint16(u2), nil
}

// in: ["0x01, 0x02", "0x01, 0x03", ...], a uint16 array represented as strings
// out: [0x0102, 0x0103, ...], a uint16 array
func ArrStringUint8ToUint16(in []string) []uint16 {
	out := make([]uint16, len(in))
	for i := 0; i < len(in); i++ {
		u16, err := StringUint8ToUint16(in[i])
		if err != nil {
			continue
		}
		out[i] = u16
	}

	return out
}

// StringToUint16 converts a string in the format of
// "0x0A0B" to uint16 0x0A0B
func StringToUint16(u string) (uint16, error) {
	if !strings.HasPrefix(u, "0x") {
		return 0, fmt.Errorf("invalid string %s", u)
	}

	u = strings.TrimPrefix(u, "0x")
	if len(u) > 4 {
		return 0, fmt.Errorf("invalid string %s", u)
	}

	v, err := strconv.ParseUint(u, 16, 16)
	if err != nil {
		return 0, err
	}

	return uint16(v), nil
}

// in: ["0x0102", "0x0103", ...], a uint16 array represented as strings
// out: [0x0102, 0x0103, ...], a uint16 array
func ArrStringToUint16(in []string) []uint16 {
	out := make([]uint16, len(in))
	for i := 0; i < len(in); i++ {
		u16, err := StringToUint16(in[i])
		if err != nil {
			continue
		}
		out[i] = u16
	}

	return out
}
