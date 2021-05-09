package cmd

import (
	"github.com/bn-k/rilkiv/account"
	"github.com/bn-k/rilkiv/app"
	"github.com/bn-k/rilkiv/db"
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

		return account.Server(ap, doc)
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
}

var accountDropDBCmd = &cobra.Command{
	Use:   "drop",
	Short: "drop the account API database",
	Long:  "drop the account API database",
	RunE: func(cmd *cobra.Command, args []string) error {
		ap, err := app.Provide()
		if err != nil {
			panic(err)
		}

		return db.DropMigrateAll(ap.Conf)
	},
}

func init() {
	rootCmd.AddCommand(accountDropDBCmd)
}
