package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/badgerodon/penv"
)

// ConfigCommand is a struct with metadata.
type ConfigCommand struct {
	Meta
}

// Run is the main program function.
func (c *ConfigCommand) Run(args []string) int {
	// Declare variables
	namearray := [4]string{"SHELTER_ENDPOINT", "SHELTER_ACCESS_KEY", "SHELTER_SECRET_KEY", "SHELTER_BUCKET"}

	if os.Getenv(namearray[0]) == "" {
		// SHELTER_ENDPOINT is the S3 URL
		fmt.Print("Enter the URL of your S3-compatible endpoint: ")
		var endpoint string
		fmt.Scanln(&endpoint)

		// SHELTER_ACCESS_KEY is the "username" key
		fmt.Print("Enter the access key for your S3-compatible endpoint: ")
		var key string
		fmt.Scanln(&key)

		// SHELTER_SECRET_KEY is the password for the "username" that has access to the bucket
		fmt.Print("Enter the secret key for your S3-compatible endpoint: ")
		var secret string
		fmt.Scanln(&secret)

		// SHELTER_BUCKET is your S3 bucket
		fmt.Print("Enter the name of your bucket: ")
		var bucket string
		fmt.Scanln(&bucket)

		envarray := [4]string{endpoint, key, secret, bucket}

		for step := 0; step <= 3; step++ {
			// set env var for the value that was input
			penv.SetEnv(namearray[step], envarray[step])
		}

		fmt.Print("Environment variables have been set. Please relaunch your shell.\n")

	} else {
		fmt.Print("Looks like you may have already set environment variables for Shelter before.")
		fmt.Print("\nDo you want to change any values? (y/n): ")
		var confirm string
		fmt.Scanln(&confirm)

		if confirm == "y" {
			// SHELTER_ENDPOINT is the S3 URL
			fmt.Print("Enter the URL of your S3-compatible endpoint: ")
			var endpoint string
			fmt.Scanln(&endpoint)

			// SHELTER_ACCESS_KEY is the "username" key
			fmt.Print("Enter the access key for your S3-compatible endpoint: ")
			var key string
			fmt.Scanln(&key)

			// SHELTER_SECRET_KEY is the password for the "username" that has access to the bucket
			fmt.Print("Enter the secret key for your S3-compatible endpoint: ")
			var secret string
			fmt.Scanln(&secret)

			// SHELTER_BUCKET is your S3 bucket
			fmt.Print("Enter the name of your bucket: ")
			var bucket string
			fmt.Scanln(&bucket)

			envarray := [4]string{endpoint, key, secret, bucket}

			for step := 0; step <= 3; step++ {
				// set env var for the value that was input
				penv.SetEnv(namearray[step], envarray[step])
			}
			fmt.Print("Environment variables have been set. Please relaunch your shell.\n")
		}
		if confirm == "n" {
			fmt.Print("No environment variables were changed.\n")
		}
		if confirm != "y" && confirm != "n" {
			fmt.Print("You did not answer y or n. Please try again.\n")
		}
	}
	return 0

}

// Synopsis provides a summary of the config command.
func (c *ConfigCommand) Synopsis() string {
	return "Interactively sets needed S3/Min.io environment variables to connect & download/upload objects."
}

// Help gets the help text for this command.
func (c *ConfigCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
