package sg

import (
	"fmt"
	"os"
	"strings"

	logger "github.com/openshift-online/ocm-common/pkg/log"
	vpcClient "github.com/openshift-online/ocm-common/pkg/test/vpc_client"

	"github.com/spf13/cobra"
)

var args struct {
	region     string
	count      int
	vpcID      string
	tags       string
	namePrefix string
}

var Cmd = &cobra.Command{
	Use:   "security-groups",
	Short: "Create security-groups",
	Long:  "Create security-groups.",
	Example: `# Create a number of security groups"
  rosa-helper create security-groups --name-prefix=mysg --region us-east-2 --vpc-id <vpc id>`,

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
		"Region of the security groups",
	)
	flags.StringVarP(
		&args.namePrefix,
		"name-prefix",
		"",
		"",
		"Name prefix of the security groups, they will be named with <prefix>-0,<prefix>-1",
	)

	flags.IntVarP(
		&args.count,
		"count",
		"",
		0,
		"Additional number of security groups to be created for the vpc",
	)
	flags.StringVarP(
		&args.vpcID,
		"vpc-id",
		"",
		"",
		"Vpc ID for the VPC created for the additional security groups",
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
}
func run(cmd *cobra.Command, _ []string) {
	vpc, err := vpcClient.GenerateVPCByID(args.vpcID, args.region)
	if err != nil {
		panic(err)
	}
	preparedSGs := []string{}
	sgDescription := "This security group is created for OCM testing"
	protocol := "tcp"
	for i := 0; i < args.count; i++ {
		sgName := fmt.Sprintf("%s-%d", args.namePrefix, i)
		sg, err := vpc.AWSClient.CreateSecurityGroup(vpc.VpcID, sgName, sgDescription)
		if err != nil {
			panic(err)
		}
		groupID := *sg.GroupId
		cidrPortsMap := map[string]int32{
			vpc.CIDRValue: 8080,
			"0.0.0.0/0":   22,
		}
		for cidr, port := range cidrPortsMap {
			_, err = vpc.AWSClient.AuthorizeSecurityGroupIngress(groupID, cidr, protocol, port, port)
			if err != nil {
				panic(err)
			}
		}

		preparedSGs = append(preparedSGs, groupID)
	}
	logger.LogInfo("ADDITIONAL SECURITY GROUPS: %s", strings.Join(preparedSGs, ","))
}
