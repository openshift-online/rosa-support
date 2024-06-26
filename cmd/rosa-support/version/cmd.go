/*
Copyright (c) 2024 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

import (
	"fmt"
	"os"

	"github.com/openshift-online/rosa-support/pkg/version"
	"github.com/spf13/cobra"
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
	_, _ = fmt.Fprintf(os.Stdout, "%s\n", version.Version)
}
