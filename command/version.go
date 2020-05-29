package command

import (
	"bytes"
	"fmt"
)

// VersionCommand is a struct with metadata.
type VersionCommand struct {
	Meta

	Name     string
	Version  string
	Revision string
}

// Run is the main program function.
func (c *VersionCommand) Run(args []string) int {
	var versionString bytes.Buffer

	fmt.Fprintf(&versionString, "%s version %s", c.Name, c.Version)
	if c.Revision != "" {
		fmt.Fprintf(&versionString, " (%s)", c.Revision)
	}

	c.UI.Output(versionString.String())
	return 0
}

// Synopsis provides a summary of the version command.
func (c *VersionCommand) Synopsis() string {
	return fmt.Sprintf("Print %s version and quit.", c.Name)
}

// Help gets the help text for this command.
func (c *VersionCommand) Help() string {
	helpText := `

	DESCRIPTION:
	Shows the current version of Shelter.

	EXAMPLES:
	shelter --version

	`
	return helpText
}
