package cmd

import (
	"fmt"

	"github.com/akira393/gotoggl/togglclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(projectCmd())
}

func projectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "Manage project resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(
		projectGetCmd(),
	)
	return cmd
}

func projectGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Print project",
		Run:   runprojectGetCmd,
	}
	return cmd
}

func runprojectGetCmd(cmd *cobra.Command, args []string) {
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
	project, err := session.GetProjects(account.Data.Workspaces[0].ID)

	if err != nil {
		fmt.Println(err)
		return
	}

	PrintStructHeader(project[0])
	for _, v := range project {
		PrintStructValues(v)
	}
}
