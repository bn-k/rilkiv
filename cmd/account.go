package cmd

import (
	"github.com/bn-k/rilkiv/account"
	"github.com/bn-k/rilkiv/app"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "run the account API",
	Long:  "run the account API lorem ipsum",
	RunE: func(cmd *cobra.Command, args []string) error {
		ap, err := app.Provide()
		if err != nil {
			panic(err)
		}

		return account.Run(ap, false)
	},
}


func init() {
	rootCmd.AddCommand(accountCmd)
}
