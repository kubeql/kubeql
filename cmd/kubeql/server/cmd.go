package server

import (
	"flag"
	"io"
	"io/ioutil"
	"strings"

	"github.com/kubeql/kubeql/pkg/log"
)

// ServerCommand represents the "server" command execution.
type ServerCommand struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

// NewServerCommand returns a ServerCommand.
func New(in io.Reader, out, errout io.Writer) *ServerCommand {
	return &ServerCommand{
		stdin:  in,
		stdout: out,
		stderr: errout,
	}
}

var (
	fs     = flag.NewFlagSet("", flag.ContinueOnError)
	listen string
)

func init() {
	fs.SetOutput(ioutil.Discard)
	fs.StringVar(&listen, "listen", ":8080", "")
}

// Run executes the command.
func (cmd *ServerCommand) Run(args ...string) error {
	// Parse flags.
	if err := fs.Parse(args); err != nil {
		return err
	}

	a := NewApp()

	log.Infof("Listening on %q", listen)

	return runServer(a, listen)
}

// Usage returns the help message.
func (cmd *ServerCommand) Usage() string {
	return strings.TrimLeft(`
usage: kubeql server
`, "\n")
}
