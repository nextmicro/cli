package main

import (
	"log"

	"github.com/nextmicro/cli/cmd/next/internal/change"
	"github.com/nextmicro/cli/cmd/next/internal/project"
	"github.com/nextmicro/cli/cmd/next/internal/proto"
	"github.com/nextmicro/cli/cmd/next/internal/run"
	"github.com/nextmicro/cli/cmd/next/internal/upgrade"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "next",
	Short:   "Next: An elegant toolkit for Go microservices.",
	Long:    `Next: An elegant toolkit for Go microservices.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(change.CmdChange)
	rootCmd.AddCommand(run.CmdRun)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
