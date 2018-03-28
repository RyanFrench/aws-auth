// Copyright © 2018 Ryan French <ryan@ryanfrench.co>

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "aws-role --role-arn [role-arn] [command]",
	Short: "Assume a role in AWS and optionally run a command",
	Long: `
Assume a role within AWS. This will set your AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY and AWS_SESSION_TOKEN environment variables, allowing you to run a command using the new role. If no command is provided, they will be exported into your current session.


Use "aws-role [command] --help" for more information about a command.`,
	Run:     func(cmd *cobra.Command, args []string) {},
	Version: "0.1.0",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match

}
