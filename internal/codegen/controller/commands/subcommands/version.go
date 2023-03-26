package subcommands

import (
	"github.com/spf13/cobra"
	"projectGenerator/internal/codegen/shared/configs"
)

func GetVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of codegen",
		Long:  `All software has versions. This is codegen's`,
		Run: func(cmd *cobra.Command, args []string) {
			isHelp := configs.GetFlag("help")
			if isHelp != nil && isHelp.(bool) {
				err := cmd.Help()
				if err != nil {
					panic(err)
					return
				}
				return
			}

			cmd.Println("codegen v1.0")
		},
	}
}
