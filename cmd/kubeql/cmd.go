package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/kubeql/kubeql/cmd/kubeql/server"
	"github.com/kubeql/kubeql/pkg/log"
)

var (
	// ErrUsage is returned when a usage message was printed and the process
	// should simply exit with an error.
	ErrUsage = errors.New("usage")

	// ErrUnknownCommand is returned when a CLI command is not specified.
	ErrUnknownCommand = errors.New("unknown command")
)

// Command is a wrapper for specific functions like server or worker
type Command interface {
	Run(args ...string) error
	Usage() string
}

func main() {
	m := NewMain()
	if err := m.Run(os.Args[1:]...); err == ErrUsage {
		os.Exit(2)
	} else if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// Main represents the main program execution.
type Main struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

// NewMain returns a new instance of Main connect to the standard input/output.
func NewMain() *Main {
	return &Main{
		stdin:  os.Stdin,
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
}

var (
	fs       = flag.NewFlagSet("", flag.ContinueOnError)
	help     bool
	logLevel string
)

func init() {
	fs.SetOutput(ioutil.Discard)
	fs.BoolVar(&help, "help", false, "print usage")
	fs.BoolVar(&help, "h", false, "")
	fs.StringVar(&logLevel, "log-level", "debug", "")
}

// Run executes the program.
func (m *Main) Run(args ...string) error {
	// Parse flags.
	if err := fs.Parse(args); err != nil {
		return err
	}

	var cmd Command
	switch fs.Arg(0) {
	case "server":
		cmd = server.New(m.stdin, m.stdout, m.stderr)
	case "":
		fmt.Fprintln(m.stdout, m.Usage())
		return nil
	default:
		return ErrUnknownCommand
	}

	if help {
		fmt.Fprintln(m.stdout, cmd.Usage())
		return nil
	}

	// New logger instance
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	if level, err := logrus.ParseLevel(logLevel); err != nil {
		logger.Level = logrus.DebugLevel
		logger.Debugf("Cannot parse log level: %s, set debug level as default", logLevel)
	} else {
		logger.Level = level
	}
	// Set logrus as default logger
	log.Set(logger)

	return cmd.Run(fs.Args()[1:]...)
}

// Usage returns the help message.
func (m *Main) Usage() string {
	return strings.TrimLeft(`
kubeql is a tool.
Usage:
	kubeql command [arguments]

The commands are:
    server      start kubeql web server application

Use "kubeql [command]" for more information about a command.
`, "\n")
}
