package tag

import (
	"os"

	awsClient "github.com/openshift-online/ocm-common/pkg/aws/aws_client"
	logger "github.com/openshift-online/ocm-common/pkg/log"
	"github.com/spf13/cobra"
)

var args struct {
	region      string
	resourceID  string
	tagKey      string
	tagValue    string
	profileName string
}
var Cmd = &cobra.Command{
	Use:   "tag",
	Short: "Delete tag",
	Long:  "Delete tag.",
	Example: `  # Delete a tag from the resource
  rosa-support delete tag --resource-id <vpc id> --region us-east-2 --tag-key key`,

	Run: run,
}

func init() {
	flags := Cmd.Flags()
	flags.SortFlags = false
	flags.StringVarP(
		&args.region,
		"region",
		"",
		"",
		"Region of the resource (required)",
	)
	flags.StringVarP(
		&args.profileName,
		"profile-name",
		"",
		"",
		"profile name to pass into aws client",
	)
	flags.StringVarP(
		&args.resourceID,
		"resource-id",
		"",
		"",
		"id of the resource (required)",
	)
	flags.StringVarP(
		&args.tagKey,
		"tag-key",
		"",
		"",
		"tag key of the resource (required)",
	)
	flags.StringVarP(
		&args.tagValue,
		"tag-value",
		"",
		"",
		"tag value of the resource",
	)

	err := Cmd.MarkFlagRequired("resource-id")
	if err != nil {
		panic(err)
	}
	err = Cmd.MarkFlagRequired("region")
	if err != nil {
		panic(err)
	}
	err = Cmd.MarkFlagRequired("tag-key")
	if err != nil {
		panic(err)
	}
}
func run(_ *cobra.Command, _ []string) {
	client, err := awsClient.CreateAWSClient(args.profileName, args.region)
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}
	_, err = client.RemoveResourceTag(args.resourceID, args.tagKey, args.tagValue)
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}
}
