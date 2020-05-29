package main

import (
	"github.com/gyamada619/shelter/command"
	"github.com/mitchellh/cli"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"download": func() (cli.Command, error) {
			return &command.DownloadCommand{
				Meta: *meta,
			}, nil
		},
		"upload": func() (cli.Command, error) {
			return &command.UploadCommand{
				Meta: *meta,
			}, nil
		},
		"list": func() (cli.Command, error) {
			return &command.ListCommand{
				Meta: *meta,
			}, nil
		},
		"config": func() (cli.Command, error) {
			return &command.ConfigCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
