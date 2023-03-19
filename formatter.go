package tlsspecloader

import (
	"fmt"
	"strconv"
	"strings"
)

// HexFormatter is a formatter that converts uint8/uint16 values in hex formatted string
// to a uint8/uint16 and vice versa.
type HexFormatter struct {
}

func (*HexFormatter) ToUint8(u string) (uint8, error) {
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

func (h *HexFormatter) ToUint8Arr(in []string) []uint8 {
	out := make([]uint8, len(in))
	for i := 0; i < len(in); i++ {
		u8, err := h.ToUint8(in[i])
		if err != nil {
			continue // will skip this value, leaving it as 0
		}
		out[i] = u8
	}

	return out
}

// DuoToUint16 converts a string in the format of
// "0x0A, 0x0B" to uint16 0x0A0B
func (h *HexFormatter) DuoToUint16(u string) (uint16, error) {
	split := strings.Split(u, ",")
	if len(split) != 2 {
		return 0, fmt.Errorf("invalid string %s", u)
	}

	u1, err := h.ToUint8(strings.TrimSpace(split[0]))
	if err != nil {
		return 0, err
	}

	u2, err := h.ToUint8(strings.TrimSpace(split[1]))
	if err != nil {
		return 0, err
	}

	return uint16(u1)<<8 | uint16(u2), nil
}

// in: ["0x01, 0x02", "0x01, 0x03", ...], a uint16 array represented in uint8 duo strings.
// out: [0x0102, 0x0103, ...], a uint16 array
func (h *HexFormatter) DuoToUint16Arr(in []string) []uint16 {
	out := make([]uint16, len(in))
	for i := 0; i < len(in); i++ {
		u16, err := h.DuoToUint16(in[i])
		if err != nil {
			continue
		}
		out[i] = u16
	}

	return out
}

// ToUint16 converts a string in the format of
// "0x0A0B" to uint16 0x0A0B
func (*HexFormatter) ToUint16(u string) (uint16, error) {
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
func (h *HexFormatter) ToUint16Arr(in []string) []uint16 {
	out := make([]uint16, len(in))
	for i := 0; i < len(in); i++ {
		u16, err := h.ToUint16(in[i])
		if err != nil {
			continue
		}
		out[i] = u16
	}

	return out
}

type DecFormatter struct {
}

func (*DecFormatter) ToUint8(u string) (uint8, error) {
	v, err := strconv.ParseUint(u, 10, 8)
	if err != nil {
		return 0, err
	}

	return uint8(v), nil
}

func (d *DecFormatter) ToUint8Arr(in []string) []uint8 {
	out := make([]uint8, len(in))
	for i := 0; i < len(in); i++ {
		u8, err := d.ToUint8(in[i])
		if err != nil {
			continue
		}
		out[i] = u8
	}

	return out
}

func (*DecFormatter) ToUint16(u string) (uint16, error) {
	v, err := strconv.ParseUint(u, 10, 16)
	if err != nil {
		return 0, err
	}

	return uint16(v), nil
}

func (d *DecFormatter) ToUint16Arr(in []string) []uint16 {
	out := make([]uint16, len(in))
	for i := 0; i < len(in); i++ {
		u16, err := d.ToUint16(in[i])
		if err != nil {
			continue
		}
		out[i] = u16
	}

	return out
}
