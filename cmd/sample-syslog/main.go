package main

import (
	"fmt"
	"log"
	"log/syslog"
	"strings"

	"github.com/jessevdk/go-flags"
)

// TODO: Make use of Facility and Severity provided by the CLI
var opts struct {
	Address   string `long:"address" description:"address of the syslog server" default:"127.0.0.1:514"`
	Transport string `long:"transport" description:"transport to use (tcp|udp)" default:"udp"`
	Facility  string `long:"facility" description:"name of the syslog facility to send msgs to" default:"local0"`
	Severity  string `long:"severity" description:"severity of the message" default:"emerg"`
	Args      struct {
		Message []string
	} `positional-args:"yes" required:"yes"`
}

func must(err error) {
	if err == nil {
		return
	}

	log.Fatal(err)
}

func main() {
	_, err := flags.Parse(&opts)
	must(err)

	switch opts.Transport {
	case "udp":
	case "tcp":
	default:
		log.Fatal(fmt.Errorf("Unknown transport %s", opts.Transport))
	}

	logger, err := syslog.Dial(opts.Transport, opts.Address,
		syslog.LOG_LOCAL0|syslog.LOG_USER, "custom-tag")
	must(err)

	defer logger.Close()

	fmt.Fprintf(logger, strings.Join(opts.Args.Message, " "))
}
