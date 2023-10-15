package upgrade

import (
	"fmt"

	"github.com/nextmicro/cli/cmd/next/internal/base"
	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the next tools",
	Long:  "Upgrade the next tools. Example: next upgrade",
	Run:   Run,
}

// Run upgrade the next tools.
func Run(_ *cobra.Command, _ []string) {
	err := base.GoInstall(
		"github.com/nextmicro/cli/cmd/next@latest",
		"github.com/nextmicro/cli/cmd/protoc-gen-go-http@latest",
		"github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
		"github.com/google/gnostic/cmd/protoc-gen-openapi@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}
