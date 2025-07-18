// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	"github.com/ramendr/ramenctl/pkg/console"
	"github.com/ramendr/ramenctl/pkg/validate"
	"github.com/spf13/cobra"
)

var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate disaster recovery configuration",
}

var ValidateClustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "Validate clusters configuration",
	Run: func(c *cobra.Command, args []string) {
		if err := validate.Clusters(configFile, outputDir); err != nil {
			console.Fatal(err)
		}
	},
}

var ValidateApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Validate protected application",
	Run: func(c *cobra.Command, args []string) {
		if err := validate.Application(configFile, outputDir, drpcName, drpcNamespace); err != nil {
			console.Fatal(err)
		}
	},
}

func init() {
	addDRPCFlags(ValidateApplicationCmd)
	addOutputFlags(ValidateCmd)
	ValidateCmd.AddCommand(ValidateClustersCmd)
	ValidateCmd.AddCommand(ValidateApplicationCmd)
}
