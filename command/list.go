package command

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/minio/minio-go/v6"
)

// ListCommand is a struct with metadata.
type ListCommand struct {
	Meta
}

// Run is the main program function.
func (c *ListCommand) Run(args []string) int {
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

	// Check the bucket to see if the package prefix exists.
	checkCh := make(chan struct{})
	defer close(checkCh)
	isRecursive := false
	objectCheckCh := minioClient.ListObjectsV2(bucket, args[0], isRecursive, checkCh)

	// Check to see if the package you are searching for exists.
	check := <-objectCheckCh
	key := check.Key
	if key == "" {
		fmt.Print("Package ", args[0], " does not exist.")
		return 1
	}

	// If the package prefix exists then get the versions available.
	doneCh := make(chan struct{})
	defer close(doneCh)
	isRecursive = true
	objectCh := minioClient.ListObjectsV2(bucket, args[0], isRecursive, doneCh)

	// Print available versions of the package.
	fmt.Printf("Available versions for package %s:\n", args[0])
	for object := range objectCh {
		objString := string(object.Key)
		split := strings.Split(objString, "/")
		fmt.Println("- ", split[1])
		if err != nil {
			fmt.Println(object.Err)
			return 1
		}

	}

	return 0
}

// Synopsis provides a summary of the list command.
func (c *ListCommand) Synopsis() string {
	return "Gets a list of versions of a given package in a bucket set in the configuration."
}

// Help gets the help text for this command.
func (c *ListCommand) Help() string {
	helpText := `

	DESCRIPTION:
	Searches the Bucket for the package prefix requested.
	
	ARGUMENTS:
	args[0] = The name of the package (apache). Do not include ".hart".
	

	EXAMPLES:
	shelter list test
	shelter list apache 

	`
	return helpText
}
