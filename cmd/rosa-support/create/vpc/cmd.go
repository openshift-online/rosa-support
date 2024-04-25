package vpc

import (
	"os"
	"strings"

	logger "github.com/openshift-online/ocm-common/pkg/log"
	vpcClient "github.com/openshift-online/ocm-common/pkg/test/vpc_client"

	"github.com/spf13/cobra"
)

var args struct {
	region       string
	name         string
	cidr         string
	tags         string
	findExisting bool
}
var Cmd = &cobra.Command{
	Use:   "vpc",
	Short: "Create vpc",
	Long:  "Create vpc.",
	Example: `  # Create a vpc named "myvpc"
  rosa-helper create vpc --name=myvpc --region us-east-2`,

	Run: run,
}

func init() {
	flags := Cmd.Flags()
	flags.SortFlags = false
	flags.StringVarP(
		&args.name,
		"name",
		"n",
		"",
		"Name of the vpc",
	)
	flags.StringVarP(
		&args.region,
		"region",
		"",
		"",
		"Vpc region",
	)
	flags.StringVarP(
		&args.cidr,
		"cidr",
		"",
		"",
		"Cidr of the vpc",
	)
	flags.StringVarP(
		&args.tags,
		"tags",
		"",
		"",
		"Vpc tags, example: tagName:tagValue,tagName2:tagValue2",
	)
	flags.BoolVarP(
		&args.findExisting,
		"find-existing",
		"",
		false,
		"Find the vpc with the same name from current region. if it does not exist, create a new one",
	)
}
func run(cmd *cobra.Command, _ []string) {
	vpc, err := vpcClient.PrepareVPC(args.name, args.region, args.cidr, args.findExisting)
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}

	logger.LogInfo("VPC ID: %s", vpc.VpcID)
	logger.LogInfo("VPC REGION: %s", vpc.Region)
	logger.LogInfo("VPC NAME: %s", vpc.VPCName)
	var tagMap map[string]string
	if args.tags != "" {
		tags := strings.Split(args.tags, ",")
		tagMap = map[string]string{}
		for _, tag := range tags {
			split := strings.Split(tag, ":")

			if len(split) == 2 {
				tagMap[split[0]] = tagMap[split[1]]
			} else {
				tagMap[split[0]] = ""
			}

		}
		_, err = vpc.AWSClient.TagResource(vpc.VpcID, tagMap)
		if err != nil {
			logger.LogError(err.Error())
			os.Exit(1)
		}
	}
}
