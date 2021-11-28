package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(configureCmd())
}

func configureCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Create config",
		Run:   runConfigulreCmd,
	}
	return cmd
}

func runConfigulreCmd(cmd *cobra.Command, args []string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Print("toggl api token: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	viper.Set("TOGGL_API_TOKEN", scanner.Text())
	viper.WriteConfigAs(home + "/.gotoggl.yaml")

}
