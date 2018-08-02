package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type ProtocolCommand interface {
	Read(r io.Reader) error
	Write(w io.Writer) error
}

func ProtocolBytes(p ProtocolCommand) ([]uint8, error) {

	var request_buffer bytes.Buffer

	writer := bufio.NewWriter(&request_buffer)
	err := p.Write(writer)
	if err != nil {
		return nil, err
	}
	writer.Flush()

	return request_buffer.Bytes(), nil
}

type ProtocolCommandAddressValue struct {
	Command uint16
	Address uint16
	Value   uint16
}

func (p *ProtocolCommandAddressValue) Write(w io.Writer) error {
	err := binary.Write(w, binary.BigEndian, p.Command)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, p.Address)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, p.Value)
	if err != nil {
		return err
	}
	return nil
}
func (p *ProtocolCommandAddressValue) Read(r io.Reader) error {
	err := binary.Read(r, binary.BigEndian, &p.Command)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.BigEndian, &p.Address)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.BigEndian, &p.Value)
	if err != nil {
		return err
	}
	return nil
}

type ProtocolCommandValue struct {
	Command uint16
	Unknown uint8
	Value   uint16
}

func (p *ProtocolCommandValue) Write(w io.Writer) error {
	err := binary.Write(w, binary.BigEndian, p.Command)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, p.Unknown)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, p.Value)
	if err != nil {
		return err
	}
	return nil
}
func (p *ProtocolCommandValue) Read(r io.Reader) error {
	err := binary.Read(r, binary.BigEndian, &p.Command)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.BigEndian, &p.Unknown)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.BigEndian, &p.Value)
	if err != nil {
		return err
	}
	return nil
}

func ProtocolRequest(command ProtocolCommand) ([]uint8, error) {

	buf_bytes, err := ProtocolBytes(command)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(buf_bytes)
	writer := bufio.NewWriter(buf)

	err = binary.Write(writer, binary.LittleEndian, crc16(buf_bytes))
	if err != nil {
		return nil, err
	}

	writer.Flush()

	return buf.Bytes(), nil
}

func ProtocolResponse(port io.ReadWriter, command ProtocolCommand) ([]uint8, error) {

	var buf bytes.Buffer
	tee := io.TeeReader(port, &buf)

	err := command.Read(tee)
	if err != nil {
		return nil, err
	}
	buf_bytes := buf.Bytes()

	crc_calculated := crc16(buf_bytes)

	var crc_read uint16
	err = binary.Read(tee, binary.LittleEndian, &crc_read)
	if err != nil {
		return nil, err
	}

	buf_bytes = buf.Bytes()

	if crc_read != crc_calculated {
		return nil, fmt.Errorf("checksum should be %.4x but was %.4x - %s", crc_calculated, crc_read, Hex(buf_bytes))
	}

	return buf_bytes, nil
}

func ProtocolRead(port io.ReadWriter, address uint16) (uint16, error) {

	request := &ProtocolCommandAddressValue{}
	request.Command = 0x0103
	request.Address = address
	request.Value = 0x0001

	request_bytes, err := ProtocolRequest(request)
	if err != nil {
		return 0, err
	}

	// fmt.Printf("Read req %s\n", Hex(request_bytes))

	_, err = port.Write(request_bytes)
	if err != nil {
		return 0, err
	}

	response := &ProtocolCommandValue{}
	response_bytes, err := ProtocolResponse(port, response)
	if err != nil {
		return 0, err
	}

	_ = response_bytes
	// fmt.Printf("Read res %s\n", Hex(response_bytes))

	return response.Value, nil
}

func ProtocolWrite(port io.ReadWriter, address uint16, value uint16) error {

	request := &ProtocolCommandAddressValue{}
	request.Command = 0x0106
	request.Address = address
	request.Value = value

	request_bytes, err := ProtocolRequest(request)
	if err != nil {
		return err
	}

	// fmt.Printf("Write req %s\n", Hex(request_bytes))

	_, err = port.Write(request_bytes)
	if err != nil {
		return err
	}

	response := &ProtocolCommandAddressValue{}
	response_bytes, err := ProtocolResponse(port, response)
	if err != nil {
		return err
	}

	_ = response_bytes
	// fmt.Printf("Write res %s\n", Hex(response_bytes))

	return nil
}
