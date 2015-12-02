package noop

import (
	"bytes"
	"io"
	"time"

	"github.com/abrander/alerto/plugins"
)

func init() {
	plugins.Register("noop", NewNoop)
}

func NewNoop() plugins.Plugin {
	return new(Noop)
}

type (
	Noop struct {
		Delay time.Duration `json:"delay" description:"Amount of time to do nothing"`
	}
)

func (n *Noop) Run(transport plugins.Transport, request plugins.Request) plugins.Result {
	time.Sleep(n.Delay)
	return plugins.NewResult(plugins.Ok, nil, "noop ;-)")
}

func (n *Noop) Exec(cmd string, arguments ...string) (io.Reader, io.Reader, error) {
	var stdoutBuf, stderrBuf bytes.Buffer
	return &stdoutBuf, &stderrBuf, nil
}

// Ensure compliance
var _ plugins.Agent = (*Noop)(nil)
var _ plugins.Transport = (*Noop)(nil)
