package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var count int

func init() {

	uuidCmd.Flags().IntVarP(&count, "count", "c", 1, "count of uuid")

	rootCmd.AddCommand(uuidCmd)
}

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate one or more unique id(s)",
	Long:  `All software has versions. This is Cli-util`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < count; i++ {
			fmt.Println(uuid.New())
		}

	},
}
