package main

import (
	"github.com/alecthomas/kingpin"
	"os"
)

var (
	app      = kingpin.New("jmc-config", "Command line tool to interact with JMC IHSV57 servos")
	debug    = app.Flag("debug", "Be more verbose.").Short('d').Bool()
	portname = app.Flag("port", "Serial port").Short('p').Default("/dev/cu.serial").String()

	upload   = app.Command("upload", "Upload configuration to servo.")
	filename = upload.Flag("config", "Path to configuration yaml.").Short('c').Default("example.yaml").String()
)

func main() {

	app.UsageTemplate(kingpin.LongHelpTemplate)
	app.HelpFlag.Short('h')

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case upload.FullCommand():
		CommandUpload(*filename, *portname, *debug)
	}

	os.Exit(0)
}
