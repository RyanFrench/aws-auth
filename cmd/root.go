// Copyright © 2018 Ryan French <ryan@ryanfrench.co>

package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	roleArn string
)

var rootCmd = &cobra.Command{
	Use:   "aws-role [command]",
	Short: "Assume a role in AWS and optionally run a command",
	Long: `
Assume a role within AWS. This will set your AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY and AWS_SESSION_TOKEN environment variables, allowing you to run a command using the new role. If no command is provided, they will be exported into your current session.


Use "aws-role [command] --help" for more information about a command.`,
	Run:     run,
	Version: "0.1.0",
	Args:    cobra.MinimumNArgs(1),
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	_, err := exec.LookPath("aws")
	if err != nil {
		log.WithError(err).Fatal("aws cli is not installed. For information on how to install the aws cli, please visit https://aws.amazon.com/cli/")
	}
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&roleArn, "role-arn", "r", "", "The arn of the role to assume in AWS (required)")
	rootCmd.MarkFlagRequired("role-arn")
}

func run(cmd *cobra.Command, args []string) {
	roleSessionName, _ := uuid.NewUUID()
	svc := sts.New(session.New())
	input := &sts.AssumeRoleInput{
		DurationSeconds: aws.Int64(3600),
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(roleSessionName.String()),
	}

	assumeRoleResponse, err := svc.AssumeRole(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case sts.ErrCodeMalformedPolicyDocumentException:
				log.WithError(aerr).
					Errorln(sts.ErrCodeMalformedPolicyDocumentException)
			case sts.ErrCodePackedPolicyTooLargeException:
				log.WithError(aerr).
					Errorln(sts.ErrCodePackedPolicyTooLargeException)
			case sts.ErrCodeRegionDisabledException:
				log.WithError(aerr).
					Errorln(sts.ErrCodeRegionDisabledException)
			default:
				log.WithError(aerr).
					Errorln(sts.ErrCodeRegionDisabledException)
			}
		} else {
			log.WithError(err).
				Errorln("Error assuming role")
		}
		return
	}

	command := exec.Command(args[0], args[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Env = append(os.Environ(),
		fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", *assumeRoleResponse.Credentials.AccessKeyId),
		fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", *assumeRoleResponse.Credentials.SecretAccessKey),
		fmt.Sprintf("AWS_SESSION_TOKEN=%s", *assumeRoleResponse.Credentials.SessionToken))

	if err := command.Run(); err != nil {
		log.
			WithField("command", command).
			WithError(err).
			Fatalln("Failed to run command")
	}
}

func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
