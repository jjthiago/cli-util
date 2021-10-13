package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cli-util",
	Long:  `All software has versions. This is Cli-util`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cli-util/v0.1")
	},
}
