package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jjthiago/cli-util/internal/common"
	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(checkMemory)
	rootCmd.AddCommand(clearTmp)
}

var checkMemory = &cobra.Command{
	Use:   "memory",
	Short: "View memory info",
	Long:  `All software has versions. This is Cli-util`,
	Run: func(cmd *cobra.Command, args []string) {

		command := common.NewCommand("")

		fmt.Println(command.Run("free", "-h"))

	},
}

var clearTmp = &cobra.Command{
	Use:   "clean_tmp",
	Short: "Clear /tmp directory",
	Long:  `All software has versions. This is Cli-util`,
	Run: func(cmd *cobra.Command, args []string) {

		if err := os.RemoveAll("/tmp"); err != nil {
			log.Fatalln(err)
		}

	},
}

func bToMb(b uint64) uint64 {

	return b / 1024 / 1024
}
