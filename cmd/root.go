package cmd

import (
	"helm-update-version/internal"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "helm-update-version",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		run(cmd)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("directory", "d", "", "Directory containing the Helm chart")
	rootCmd.MarkFlagRequired("directory")
	rootCmd.Flags().StringP("version", "v", "", "Specific version to update to")
	rootCmd.Flags().StringP("changetype", "c", "", "major/minor/hotfix")
	rootCmd.Flags().BoolP("verbose", "", false, "detailed output")
}

func run(cmd *cobra.Command) {
	directory, _ := cmd.Flags().GetString("directory")
	version, _ := cmd.Flags().GetString("version")
	changeType, _ := cmd.Flags().GetString("changetype")
	verbose, _ := cmd.Flags().GetBool("verbose")

	// fmt.Println("changetype", changeType)
	// fmt.Println("version", version)

	if version != "" && changeType == "" {
		internal.UpdateNoEdit(directory, version, verbose)

	} else if version != "" && changeType != "" {
		internal.UserVersionUpdate(directory, changeType, version, verbose)

	} else {
		internal.Update(directory, changeType, verbose)

	}
}
