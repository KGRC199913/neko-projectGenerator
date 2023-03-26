package subcommands

import (
	"github.com/spf13/cobra"
	"projectGenerator/internal/codegen/domain/generate"
	"projectGenerator/internal/codegen/shared/configs"
)

func GetGenerateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "Generate a project skeleton",
		Long:  `Generate a project skeleton based on the template and the configuration file.`,
		Run: func(cmd *cobra.Command, args []string) {
			configs.ReadConfig()

			isHelp := configs.GetFlag("help")
			if isHelp != nil && isHelp.(bool) {
				err := cmd.Help()
				if err != nil {
					panic(err)
					return
				}
				return
			}

			isSuccess, err := generate.GenerateProject(cmd, args)
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
}
