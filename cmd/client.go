package cmd

import (
	"fmt"

	"github.com/akira393/togglclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(clientCmd())
}

func clientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client",
		Short: "Manage client resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(
		clientGetCmd(),
	)
	return cmd
}

func clientGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Print client",
		Run:   runClientGetCmd,
	}
	return cmd
}

func runClientGetCmd(cmd *cobra.Command, args []string) {
	token := viper.GetString("TOGGL_API_TOKEN")
	if token == "" {
		fmt.Println("Unable to locate credentials. You can configure credentials by running 'gotoggl configure'.")
		return
	}
	session := togglclient.OpenSession(token)
	client, err := session.GetClients()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("index", ",", "id", ",", "name")
	for i, v := range client {
		fmt.Println(i, ",", v.ID, ",", v.Name)
	}
	return
}
