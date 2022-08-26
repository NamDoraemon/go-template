package cmd

import (
	"errors"
	"github.com/urfave/cli/v2"
	"os"
)

func SetupCmd() error {
	commandLine := cli.NewApp()
	commandLine.EnableBashCompletion = true
	commandLine.Action = func(c *cli.Context) error {
		return errors.New("please enter your command")
	}
	commandLine.Commands = []*cli.Command{
		{Name: "http", Usage: "Start up HTTP Serving", Action: HTTPServe},
		{Name: "grpc", Usage: "Start up Grpc Serving", Action: GRPCServe},
	}
	return commandLine.Run(os.Args)
}
