package command

import "github.com/mitchellh/cli"

// Meta contain the meta-option that nearly all subcommand inherited.
type Meta struct {
	UI cli.Ui
}
