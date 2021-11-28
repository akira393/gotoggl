package cmd

import (
	"errors"

	"github.com/akira393/togglclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(reportCmd())
}

func reportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "Manage report resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(
		reportDetailCmd(),
	)
	return cmd
}

func reportDetailCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "detail <since> <until>",
		Short: "Print detial time entry",
		RunE:  runReportDetailCmd,
	}
	return cmd
}

func runReportDetailCmd(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("since and until is required")
	}
	//TODO: バリデーション
	since := args[0]
	until := args[1]
	token := viper.GetString("TOGGL_API_TOKEN")
	if token == "" {
		return errors.New("Unable to locate credentials. You can configure credentials by running 'gotoggl configure'.")
	}

	session := togglclient.OpenSession(token)
	account, err := session.GetAccount()
	if err != nil {
		return err
	}

	workspaceId := account.Data.Workspaces[0].ID

	pagenum, err := getDetailedReportPage(session, workspaceId, since, until)
	if err != nil {
		return err
	}
	for i := 1; i <= pagenum; i++ {
		report, _ := session.GetDetailedReport(workspaceId, since, until, i)

		PrintStructHeader(report.Data[0])
		for _, v := range report.Data {
			PrintStructValues(v)
		}
	}
	return nil
}

func getDetailedReportPage(session togglclient.Session, workspace int, since, until string) (int, error) {
	report, _ := session.GetDetailedReport(workspace, since, until, 1)
	return report.TotalCount/report.PerPage + 1, nil

}
