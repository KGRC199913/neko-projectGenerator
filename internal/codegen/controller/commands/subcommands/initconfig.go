package subcommands

import (
	"github.com/spf13/cobra"
	"projectGenerator/internal/codegen/domain/initconfig"
	"projectGenerator/internal/codegen/shared/configs"
)

func GetInitConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Create an empty configuration file",
		Long: `Create an empty configuration file.
				Usage: codegen config --path <path> --name <name>`,
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

			path := configs.GetFlag("path")
			name := configs.GetFlag("name")
			initconfig.CreateConfig(cmd, args, path.(string), name.(string))
		},
	}

	path := ""
	cmd.LocalFlags().StringVarP(&path, "path", "p", "./configs", "path to the configuration file")
	configs.SetFlag("path", cmd.LocalFlags().Lookup("path"))

	name := ""
	cmd.LocalFlags().StringVarP(&name, "name", "n", "config.yaml", "name of the configuration file")
	configs.SetFlag("name", cmd.LocalFlags().Lookup("name"))

	return cmd
}
