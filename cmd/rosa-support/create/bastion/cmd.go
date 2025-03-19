package bastion

import (
	logger "github.com/openshift-online/ocm-common/pkg/log"
	vpcClient "github.com/openshift-online/ocm-common/pkg/test/vpc_client"
	"github.com/spf13/cobra"
	"os"
)

var args struct {
	region           string
	vpcID            string
	availabilityZone string
	privateKeyPath   string
	keyPairName      string
}

var Cmd = &cobra.Command{
	Use:   "bastion",
	Short: "Create bastion proxy",
	Long:  "Create bastion proxy.",
	Example: `  # Create a bastion proxy in region 'us-east-2'
  rosa-support create bastion --region us-east-2 --availability-zone us-east-2a --vpc-id <vpc id> --keypair-name <name> --private-key-path <path>`,

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
		"Vpc region (required)",
	)
	err := Cmd.MarkFlagRequired("region")
	if err != nil {
		logger.LogError("%s", err.Error())
		os.Exit(1)
	}
	flags.StringVarP(
		&args.vpcID,
		"vpc-id",
		"",
		"",
		"ID of vpc to be used to create bastion proxy (required)",
	)
	err = Cmd.MarkFlagRequired("vpc-id")
	if err != nil {
		logger.LogError("%s", err.Error())
		os.Exit(1)
	}
	flags.StringVarP(
		&args.availabilityZone,
		"availability-zone",
		"",
		"",
		"Availability zone to create public subnet of specified vpc (required)",
	)
	err = Cmd.MarkFlagRequired("availability-zone")
	if err != nil {
		logger.LogError("%s", err.Error())
		os.Exit(1)
	}
	flags.StringVarP(
		&args.keyPairName,
		"keypair-name",
		"",
		"",
		"key pair will be created with the name (required)",
	)
	err = Cmd.MarkFlagRequired("keypair-name")
	if err != nil {
		logger.LogError("%s", err.Error())
		os.Exit(1)
	}
	flags.StringVarP(
		&args.privateKeyPath,
		"private-key-path",
		"",
		"",
		"record generated private ssh key in the given path (required)",
	)
	err = Cmd.MarkFlagRequired("private-key-path")
	if err != nil {
		logger.LogError("%s", err.Error())
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, _ []string) {
	vpc, err := vpcClient.GenerateVPCByID(args.vpcID, args.region)
	if err != nil {
		panic(err)
	}
	proxyUrl, err := vpc.PrepareBastionProxy(args.availabilityZone, args.keyPairName, args.privateKeyPath)
	if err != nil {
		panic(err)
	}

	logger.LogInfo("Bastion PROXY URL: %s", proxyUrl)
}
