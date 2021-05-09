package cmd

import (
	"github.com/spf13/cobra"


)

var notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "run the notify API",
	Long:  "run the notify API lorem ipsum",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(notifyCmd)
}
