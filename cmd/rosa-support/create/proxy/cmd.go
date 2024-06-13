package proxy

import (
	"fmt"
	"os"

	logger "github.com/openshift-online/ocm-common/pkg/log"
	vpcClient "github.com/openshift-online/ocm-common/pkg/test/vpc_client"

	"github.com/spf13/cobra"
)

var args struct {
	region           string
	vpcID            string
	availabilityZone string
	privateKeyPath   string
	keyPairName      string
	caFilePath       string
}

var Cmd = &cobra.Command{
	Use:   "proxy",
	Short: "Create proxy",
	Long:  "Create proxy.",
	Example: `  # Create a proxy
  rosa-support create proxy --region us-east-2 --vpc-id <vpc id> --availability-zone <AZ> --ca-file <path> --keypair-name <name> --private-key-path <path>`,
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
	flags.StringVarP(
		&args.vpcID,
		"vpc-id",
		"",
		"",
		"Creates a pair of subnets (required)",
	)
	flags.StringVarP(
		&args.availabilityZone,
		"availability-zone",
		"",
		"",
		"Creates a proxy in the indicated AZ (required)",
	)

	flags.StringVarP(
		&args.caFilePath,
		"ca-file",
		"",
		"",
		"Creates a proxy and stores the ca file (required)",
	)

	flags.StringVarP(
		&args.keyPairName,
		"keypair-name",
		"",
		"",
		"Create a key pair with the name (required)",
	)
	flags.StringVarP(
		&args.privateKeyPath,
		"private-key-path",
		"",
		"",
		"Stores key pair in the given path (required)",
	)

	err := Cmd.MarkFlagRequired("vpc-id")
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}
	err = Cmd.MarkFlagRequired("region")
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}
	err = Cmd.MarkFlagRequired("availability-zone")
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}
	err = Cmd.MarkFlagRequired("ca-file")
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}
	err = Cmd.MarkFlagRequired("keypair-name")
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}

	err = Cmd.MarkFlagRequired("private-key-path")
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, _ []string) {
	vpc, err := vpcClient.GenerateVPCByID(args.vpcID, args.region)
	if err != nil {
		panic(err)
	}
	_, ip, ca, err := vpc.LaunchProxyInstance(args.availabilityZone, args.keyPairName, args.privateKeyPath)
	if err != nil {
		panic(err)
	}
	httpProxy := fmt.Sprintf("http://%s:8080", ip)
	httpsProxy := fmt.Sprintf("https://%s:8080", ip)
	file, err := os.OpenFile(args.caFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(ca)
	if err != nil {
		panic(err)
	}
	logger.LogInfo("HTTP PROXY: %s", httpProxy)
	logger.LogInfo("HTTPs PROXY: %s", httpsProxy)
	logger.LogInfo("CA FILE PATH: %s", args.caFilePath)
}
