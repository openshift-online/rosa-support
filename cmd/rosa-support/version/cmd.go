package version

import (
	"fmt"
	"github.com/openshift-online/rosa-support/pkg/version"
	"github.com/spf13/cobra"
	"os"
)

const (
	use   = "version"
	short = "Prints the version of the tool"
	long  = "Prints the version number of the tool"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Args:  cobra.NoArgs,
		Run:   run,
	}
}

func run(cmd *cobra.Command, args []string) {
	_, _ = fmt.Fprintf(os.Stdout, "%s (build %s)\n", version.Version, version.VersionStamp)
}
