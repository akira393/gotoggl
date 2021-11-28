package cmd

import (
	"fmt"

	"github.com/akira393/togglclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(tagCmd())
}

func tagCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tag",
		Short: "Manage tag resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(
		tagGetCmd(),
	)
	return cmd
}

func tagGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Print tag",
		Run:   runtagGetCmd,
	}
	return cmd
}

func runtagGetCmd(cmd *cobra.Command, args []string) {
	token := viper.GetString("TOGGL_API_TOKEN")
	if token == "" {
		fmt.Println("Unable to locate credentials. You can configure credentials by running 'gotoggl configure'.")
		return
	}
	session := togglclient.OpenSession(token)

	account, err := session.GetAccount()
	if err != nil {
		fmt.Println(err)
		return
	}

	PrintStructHeader(account.Data.Tags[0])
	for _, v := range account.Data.Tags {
		PrintStructValues(v)
	}
}
