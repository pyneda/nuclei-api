package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	outputCfgFile string
)

// dumpconfigCmd represents the dumpconfig command
var dumpconfigCmd = &cobra.Command{
	Use:   "dumpconfig",
	Short: "Dumps default configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.SafeWriteConfigAs(outputCfgFile)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Default configuration file created in: ", outputCfgFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(dumpconfigCmd)
	dumpconfigCmd.Flags().StringVarP(&outputCfgFile, "output", "o", ".nuclei-api.yaml", "Output file to write configuration to")
}
