package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dash",
		Short: "Dash is a super fast automation framework for Kubernetes",
		Long: `A fast and extensible automation framework for Kubernetes.
      We focus on supporting multiple templating engines in one tool, and encouraging declarative resource management.
      For more info, check out https://github.com/redhat-cop/dash`,
	}
)

func errorOut(cmd *cobra.Command) {
	cmd.Help()
	os.Exit(1)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
