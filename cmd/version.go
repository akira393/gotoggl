package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version  string
	Revision string
	BuildTag string
)

func init() {
	rootCmd.AddCommand(versionCmd())
}

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("gotoggl version: %s, revision: %s, buildTag: %s\n", Version, Revision, BuildTag)
		},
	}

	return cmd
}
