package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/jacobsa/go-serial/serial"
	"io"
	"io/ioutil"
	"log"
)

func ServoConnect(port io.ReadWriteCloser) error {

	fmt.Print("connecting ")
	con, err := ProtocolRead(port, 0x0080)
	if err != nil {
		fmt.Println("FAIL")
		return err
	}
	if con != 0x0012 {
		fmt.Println("INVALID")
		return fmt.Errorf("invalid response (%.4x)", con)
	}
	fmt.Println("OK")

	return nil
}

func ServoWrite(port io.ReadWriteCloser, config Config) error {

	fmt.Println("writing ")
	for _, w := range config.Writes() {
		fmt.Printf(" %s ", Pad2Len(w.Name, "_", 45))
		err := ProtocolWrite(port, w.Address, w.Value)
		if err != nil {
			fmt.Println("FAIL")
			return err
		}
		fmt.Println("OK")
	}

	return nil
}

func CommandUpload(filename string, portname string, debug bool) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	config := Config{}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	options := serial.OpenOptions{
		PortName:              portname,
		BaudRate:              57600,
		DataBits:              8,
		StopBits:              1,
		MinimumReadSize:       0,
		InterCharacterTimeout: 200,
	}

	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	defer port.Close()

	err = ServoConnect(port)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ServoWrite(port, config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}
