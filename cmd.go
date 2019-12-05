package tfmbm

import (
	"os"

	"github.com/tkuchiki/tfmbm/converter"
	"github.com/tkuchiki/tfmbm/parser"
	"gopkg.in/alecthomas/kingpin.v2"
)

type Cmd struct {
}

func NewCmd() *Cmd {
	return &Cmd{}
}

const version = "0.0.1"

func (c *Cmd) Run() error {
	fpath := kingpin.Arg("filepath", "Terraform file path").Required().String()
	kingpin.Version(version)
	kingpin.Parse()

	p := parser.NewV1Parser()
	f, err := os.Open(*fpath)
	if err != nil {
		return err
	}

	bindings, err := p.ParseBinding(f)
	if err != nil {
		return err
	}

	err = converter.ConvertProjectIamMembers(bindings)
	if err != nil {
		return err
	}

	return nil
}
