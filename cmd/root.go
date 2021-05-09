package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var verb bool

var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "api",
	Long:  `api`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVar(&verb, "verb", false, "database verbosity")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigFile("config.yaml")

	err := viper.MergeInConfig()
	if err != nil {
		log.Fatalf("merge config error: %v", err)
	}

	viper.SetConfigFile(".env")
	_ = viper.MergeInConfig()
}
