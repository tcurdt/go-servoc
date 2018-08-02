package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"
	"testing"
)

type RW struct {
	io.Reader
	io.Writer
}

// Tx 01 .
// Tx 06 .
// Tx 00 .
// Tx 06 .
// Tx 00 .
// Tx 00 .
// Tx 69 i
// Tx cb .
// Rx 01 .
// Rx 06 .
// Rx 00 .
// Rx 06 .
// Rx 00 .
// Rx 00 .
// Rx 69 i
// Rx cb .

func TestProtocolWrite(t *testing.T) {

	var tests = []struct {
		address uint16
		value   uint16
		tx      []uint8
		rx      []uint8
	}{
		{
			0x0006, 0x0000,
			[]uint8{0x01, 0x06, 0x00, 0x06, 0x00, 0x00, 0x69, 0xcb},
			[]uint8{0x01, 0x06, 0x00, 0x06, 0x00, 0x00, 0x69, 0xcb},
		},
	}

	for _, tt := range tests {

		var tx bytes.Buffer

		rw := RW{
			Reader: bytes.NewReader(tt.rx),
			Writer: &tx,
		}

		err := ProtocolWrite(rw, tt.address, tt.value)
		if err != nil {
			t.Errorf("error %v", err)
		}

		bytes := tx.Bytes()

		if !reflect.DeepEqual(bytes, tt.tx) {
			fmt.Printf("tx actual:   %stx expected: %s", hex.Dump(bytes), hex.Dump(tt.tx))
			t.Error("error tx not equal")
		}
	}
}

// Tx 01 .
// Tx 03 .
// Tx 00 .
// Tx a0 .
// Tx 00 .
// Tx 01 .
// Tx 84 .
// Tx 28 (
// Rx 01 .
// Rx 03 .
// Rx 02 .
// Rx 00 .
// Rx 00 .
// Rx b8 .
// Rx 44 D

func TestProtocolRead(t *testing.T) {

	var tests = []struct {
		address uint16
		tx      []uint8
		rx      []uint8
		val     uint16
	}{
		{
			0x00a0,
			[]uint8{0x01, 0x03, 0x00, 0xa0, 0x00, 0x01, 0x84, 0x28},
			[]uint8{0x01, 0x03, 0x02, 0x00, 0x00, 0xb8, 0x44},
			0x0000,
		},
	}

	for _, tt := range tests {

		var tx bytes.Buffer

		rw := RW{
			Reader: bytes.NewReader(tt.rx),
			Writer: &tx,
		}

		val, err := ProtocolRead(rw, tt.address)
		if err != nil {
			t.Errorf("error %v", err)
		}

		bytes := tx.Bytes()

		if !reflect.DeepEqual(bytes, tt.tx) {
			fmt.Printf("tx actual:   %stx expected: %s", hex.Dump(bytes), hex.Dump(tt.tx))
			t.Error("error tx not equal")
		}

		if val != tt.val {
			t.Error("error val not equal")
		}
	}
}
