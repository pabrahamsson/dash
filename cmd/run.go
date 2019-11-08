package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	inv "github.com/redhat-cop/dash/pkg/inventory"
	"github.com/spf13/cobra"
)

const (
	invPathDefault = "."
	invPathUsage   = "Path to Inventory, relative or absolute"
)

var (
	invPath string

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Process an inventory of templates and apply it to a cluster.",
		Long:  `Long version...`,
		Run: func(cmd *cobra.Command, args []string) {
			var i inv.Inventory
			var ns string

			yamlFile, err := ioutil.ReadFile(filepath.Join(invPath, "dash.yaml"))
			if err != nil {
				fmt.Printf("Error: Couldn't load dash inventory: %v\n\n", err)
				errorOut(cmd)
			}

			i.Load(yamlFile, invPath)
			err = i.Process(&ns)
			if err != nil {
				fmt.Println("Error: " + err.Error())
				errorOut(cmd)
			}
		},
	}
)

func init() {
	runCmd.Flags().StringVarP(&invPath, "inventory", "i", invPathDefault, invPathUsage)
	rootCmd.AddCommand(runCmd)
}
