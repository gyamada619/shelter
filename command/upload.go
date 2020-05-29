package command

import (
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v6"
)

// UploadCommand is a struct with metadata.
type UploadCommand struct {
	Meta
}

// Run is the main program function.
func (c *UploadCommand) Run(args []string) int {

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
	// Name the package that is on the bucket to be package.hart
	packagename := args[0] + "/" + args[1] + "/" + args[0] + ".hart"

	// Check if the file exists and if so upload to the bucket otherwise fail
	if _, err := os.Stat(args[2]); err == nil {
		fmt.Print("Package file found. Uploading to bucket.")
		upload, err := minioClient.FPutObject(bucket, packagename, args[2], minio.PutObjectOptions{})
		if err != nil {
			fmt.Println(err)
			return 1
		}
		fmt.Print("\nFile uploaded: ", upload, " bytes.")
	} else {
		fmt.Print("No file found to upload.")
	}

	return 0
}

// Synopsis provides a summary of the upload command.
func (c *UploadCommand) Synopsis() string {
	return "Uploads a version of a given package to a bucket set in the configuration."
}

// Help gets the help text for this command.
func (c *UploadCommand) Help() string {
	helpText := `

	DESCRIPTION:
	Uploads Habitat .hart files to an S3-compatible bucket.
	File name stored on the bucket is changed to be "packagename.hart".

	ARGUMENTS:
	args[0] = The final name of the package for the file ultimately stored in the bucket (apache). Do not include ".hart".
	args[1] = The version of the package to upload.
	args[2] = Local path to your .hart file.

	EXAMPLES:
	shelter upload test 0.1.0 ./test-0.1.0.hart
	shelter upload apache 1.5.6 ./apache-1.5.6.hart

	`
	return helpText
}
