package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	inv "github.com/redhat-cop/dash/pkg/inventory"
	"github.com/spf13/cobra"
)

const (
	invDryRunDefault = false
	invDryRunUsage = "If true, only print the object that would be sent, without sending it"
	invPathDefault = "."
	invPathUsage   = "Path to Inventory, relative or absolute"
)

var (
	invDryRun bool
	invPath string

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Process an inventory of templates and apply it to a cluster.",
		Long:  `Long version...`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var i inv.Inventory
			var ns string

			yamlFile, err := ioutil.ReadFile(filepath.Join(invPath, "dash.yaml"))
			if err != nil {
				return fmt.Errorf("Couldn't load inventory: %v\n", err)
			}

			err = i.Load(yamlFile, invPath)
			if err != nil {
				return fmt.Errorf("Failed to load inventory: %v\n", err)
			}

			err = i.Process(&ns, &invDryRun)
			if err != nil {
				return fmt.Errorf("Failed to process inventory: %v\n", err)
			}
			return nil
		},
	}
)

func init() {
	runCmd.Flags().StringVarP(&invPath, "inventory", "i", invPathDefault, invPathUsage)
	runCmd.Flags().BoolVarP(&invDryRun, "dry-run", "", invDryRunDefault, invDryRunUsage)
	rootCmd.AddCommand(runCmd)
}
