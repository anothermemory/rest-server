package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var grpcCmd = &cobra.Command{
	Use: "grpc",
	//Aliases: []string{"grpc"},
	Short: "starts grpc server",
	Long:  "long grpc command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GRPC")
	},
}

func init() {
	RootCmd.AddCommand(grpcCmd)
}
