package controller

import (
	"projectGenerator/internal/codegen/controller/commands"
	"projectGenerator/internal/codegen/controller/commands/subcommands"
)

func Run() {
	rootCmd := commands.GetRootCmd()
	versionCmd := subcommands.GetVersionCmd()
	initConfigCmd := subcommands.GetInitConfigCmd()
	generateCmd := subcommands.GetGenerateCmd()

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initConfigCmd)
	rootCmd.AddCommand(generateCmd)

	rootCmd.Execute()
}
