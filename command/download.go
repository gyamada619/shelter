package command

import (
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v6"
)

// DownloadCommand is a struct with metadata.
type DownloadCommand struct {
	Meta
}

// Run is the main program function.
func (c *DownloadCommand) Run(args []string) int {

	// Initialize variables, grab them from ENV.
	endpoint := os.Getenv("SHELTER_ENDPOINT")
	accessKeyID := os.Getenv("SHELTER_ACCESS_KEY")
	secretAccessKey := os.Getenv("SHELTER_SECRET_KEY")
	bucket := os.Getenv("SHELTER_BUCKET")
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Example format packagename would be proget/0.1.0/proget.hart
	// Example format localpath would be ./proget/proget-0.1.0.hart
	packagename := args[0] + "/" + args[1] + "/" + args[0] + ".hart"
	localpath := "./" + args[0] + "/" + args[0] + "-" + args[1] + ".hart"

	// Download object from minIO
	err = minioClient.FGetObject(bucket, packagename, localpath, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return 1
	}

	// check localpath to ensure it was downloaded
	if _, err := os.Stat(localpath); err == nil {
		fmt.Print("Package ", args[0], " was downloaded successfully.\n")
	} else {
		fmt.Print("We failed you, Luke Skywalker")
	}

	return 0
}

// Synopsis provides a summary of the download command.
func (c *DownloadCommand) Synopsis() string {
	return "Downloads Habitat .hart files from a bucket."
}

// Help gets the help text for this command.
func (c *DownloadCommand) Help() string {
	helpText := `

	DESCRIPTION:
	Downloads Habitat .hart files from a bucket.

	ARGUMENTS:
	args[0] = The name of the package (apache). Do not include ".hart".
	args[1] = The version of the package to download.

	EXAMPLES:
	shelter download test 0.1.0
	shelter download apache 1.5.6

	`
	return helpText
}
