package proto

import (
	"github.com/nextmicro/cli/cmd/next/internal/proto/add"
	"github.com/nextmicro/cli/cmd/next/internal/proto/client"
	"github.com/nextmicro/cli/cmd/next/internal/proto/server"
	"github.com/spf13/cobra"
)

// CmdProto represents the proto command.
var CmdProto = &cobra.Command{
	Use:   "proto",
	Short: "Generate the proto files",
	Long:  "Generate the proto files.",
}

func init() {
	CmdProto.AddCommand(add.CmdAdd)
	CmdProto.AddCommand(client.CmdClient)
	CmdProto.AddCommand(server.CmdServer)
}
