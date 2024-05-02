package vpc

import (
	"os"

	logger "github.com/openshift-online/ocm-common/pkg/log"
	vpcClient "github.com/openshift-online/ocm-common/pkg/test/vpc_client"

	"github.com/spf13/cobra"
)

var args struct {
	region     string
	totalClean bool
	vpcID      string
}
var Cmd = &cobra.Command{
	Use:   "vpc",
	Short: "Delete vpc",
	Long:  "Delete vpc.",
	Example: `  # Delete a vpc with vpc ID
  ocmqe delete vpc --vpc-id <vpc id> --region us-east-2`,

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
		"id of the vpc (required)",
	)
	flags.BoolVarP(
		&args.totalClean,
		"total-clean",
		"",
		false,
		"find the vpc with same name",
	)
	err := Cmd.MarkFlagRequired("vpc-id")
	if err != nil {
		panic(err)
	}
	err = Cmd.MarkFlagRequired("region")
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
	err = vpc.DeleteVPCChain(args.totalClean)
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}
}
