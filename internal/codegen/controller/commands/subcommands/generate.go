package subcommands

import (
	"github.com/spf13/cobra"
	"projectGenerator/internal/codegen/domain/generate"
	"projectGenerator/internal/codegen/shared/configs"
	"projectGenerator/internal/codegen/shared/constants"
)

const (
	// ConfigMode is the mode to generate a configuration file
	ConfigMode = "config"
	// CopyMode is the mode to generate a project skeleton using the copy method
	CopyMode = "copy"
)

func GetGenerateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a project skeleton",
		Long: `Generate a project skeleton based on the template and the configuration file.
				Usage: codegen generate --config <path> --mode <mode>
				`,
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
			configs.ReadConfig(path.(string))

			mode := configs.GetFlag("mode")
			modeInt := -1
			if mode == ConfigMode {
				modeInt = 0
			} else if mode == CopyMode {
				modeInt = 1
			} else {
				cmd.Println("Invalid mode")
				return
			}

			isSuccess, err := generate.GenerateProject(cmd, args, modeInt)
			if err != nil {
				cmd.Println("Error when generating project: ", err)
			}

			if isSuccess {
				cmd.Println("Project generated successfully")
			} else {
				cmd.Println("Project generated failed")
			}
		},
	}

	path := ""
	cmd.LocalFlags().StringVarP(&path, "path", "p", constants.CurrentDir+"configs", "path to the configuration file")
	configs.SetFlag("path", cmd.LocalFlags().Lookup("path"))

	name := ""
	cmd.LocalFlags().StringVarP(&name, "mode", "m", "config", "generate mode")
	configs.SetFlag("mode", cmd.LocalFlags().Lookup("mode"))

	return cmd
}
