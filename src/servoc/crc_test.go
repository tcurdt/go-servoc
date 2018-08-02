package main

import (
	"testing"
)

func TestCrc(t *testing.T) {

	var tests = []struct {
		data []uint8
		crc  uint16
	}{
		{[]uint8{0x01, 0x03, 0x00, 0x85, 0x00, 0x01}, 0xe395},
		{[]uint8{0x01, 0x06, 0x00, 0x0a, 0x00, 0x00}, 0xc8a9},
		{[]uint8{0x01, 0x06, 0x00, 0x06, 0x00, 0x00}, 0xcb69},
	}

	for _, tt := range tests {
		buf := tt.data
		crc_found := crc16(buf)
		crc_expected := tt.crc
		if crc_found != crc_expected {
			t.Errorf("invalid crc %.4x != %.4x", crc_found, crc_expected)
		}
	}

}
