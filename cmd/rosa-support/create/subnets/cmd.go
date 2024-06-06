package subnets

import (
	"os"
	"strings"

	logger "github.com/openshift-online/ocm-common/pkg/log"
	vpcClient "github.com/openshift-online/ocm-common/pkg/test/vpc_client"

	"github.com/spf13/cobra"
)

var args struct {
	region            string
	availabilityZones string
	vpcID             string
	tags              string
}

var Cmd = &cobra.Command{
	Use:   "subnets",
	Short: "Create subnets",
	Long:  "Create subnets.",
	Example: `  # Create a pair of subnets in region 'us-east-2'
  rosa-support create subnets --region us-east-2 --vpc-id <vpc id> --availability-zones <AZs>`,

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
		"Vpc region",
	)
	err := Cmd.MarkFlagRequired("region")
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}
	flags.StringVarP(
		&args.availabilityZones,
		"availability-zones",
		"",
		"",
		"Availability zones to create subnets in",
	)
	err = Cmd.MarkFlagRequired("availability-zones")
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}
	flags.StringVarP(
		&args.vpcID,
		"vpc-id",
		"",
		"",
		"ID of vpc to be created",
	)
	err = Cmd.MarkFlagRequired("vpc-id")
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
	availabilityZones := strings.Split(args.availabilityZones, ",")
	for _, availabilityZone := range availabilityZones {
		subnetMap, err := vpc.PreparePairSubnetByZone(availabilityZone)
		if err != nil {
			panic(err)
		}
		for subnetType, subnet := range subnetMap {
			logger.LogInfo("AVAILABILITY ZONE %s %s SUBNET: %s", availabilityZone, strings.ToUpper(subnetType),
				subnet.ID)
		}
	}
}
