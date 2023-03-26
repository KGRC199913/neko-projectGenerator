package commands

import (
	"github.com/spf13/cobra"
	"projectGenerator/internal/codegen/shared/configs"
)

func GetRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "codegen",
		Short: "Generate a project skeleton",
		Long:  `Generate a project skeleton based on the template and the configuration file.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Please specify a subcommand")
			// list all subcommands: version, generate, each also have --help flag
			cmd.Println("Available Commands:")
			cmd.Println("  version")
			cmd.Println("  config")
			cmd.Println("  generate")

			// Example usage
			cmd.Println("Use \"codegen [command] --help\" for more information about a command.")
		},
	}

	isHelp := false
	cmd.PersistentFlags().BoolVarP(&isHelp, "help", "h", false, "help for codegen")
	configs.SetFlag("help", cmd.PersistentFlags().Lookup("help"))

	return cmd
}
