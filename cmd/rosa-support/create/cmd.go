package create

import (
	"github.com/openshift-online/rosa-support/cmd/rosa-support/create/proxy"
	"github.com/openshift-online/rosa-support/cmd/rosa-support/create/sg"
	subnets "github.com/openshift-online/rosa-support/cmd/rosa-support/create/subnets"
	"github.com/openshift-online/rosa-support/cmd/rosa-support/create/vpc"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"add"},
	Short:   "Create a resource from stdin",
	Long:    "Create a resource from stdin",
}

func init() {
	Cmd.AddCommand(vpc.Cmd)
	Cmd.AddCommand(sg.Cmd)
	Cmd.AddCommand(subnets.Cmd)
	Cmd.AddCommand(proxy.Cmd)
}
