package bastion_proxy

import (
	"os"

	logger "github.com/openshift-online/ocm-common/pkg/log"
	vpcClient "github.com/openshift-online/ocm-common/pkg/test/vpc_client"

	"github.com/spf13/cobra"
)

var args struct {
	region     string
	vpcID      string
	instanceID string
}

var Cmd = &cobra.Command{
	Use:   "bastion-proxy",
	Short: "Delete bastion proxy",
	Long:  "Delete bastion proxy.",
	Example: `  # Delete a bastion proxy in region us-east-2
  rosa-support delete bastion-proxy --region us-east-2 --vpc-id <vpc id> --instance-id <instance id>`,

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
		"Region of the vpc (required)",
	)
	flags.StringVarP(
		&args.vpcID,
		"vpc-id",
		"",
		"",
		"ID of the vpc (required)",
	)
	flags.StringVarP(
		&args.instanceID,
		"instance-id",
		"",
		"",
		"Instance ID of the bastion (required)",
	)
	err := Cmd.MarkFlagRequired("region")
	if err != nil {
		panic(err)
	}
	err = Cmd.MarkFlagRequired("vpc-id")
	if err != nil {
		panic(err)
	}
	err = Cmd.MarkFlagRequired("instance-id")
	if err != nil {
		panic(err)
	}
}

func run(cmd *cobra.Command, _ []string) {
	vpc, err := vpcClient.GenerateVPCByID(args.vpcID, args.region)
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}

	var instanceIDs []string
	instanceIDs = append(instanceIDs, args.instanceID)
	filters := []map[string][]string{
		{
			"vpc-id": {
				vpc.VpcID,
			},
		},
	}
	insts, err := vpc.AWSClient.ListInstances(instanceIDs, filters...)
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}

	err = vpc.DestroyBastionProxy(insts[0])
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}
}
