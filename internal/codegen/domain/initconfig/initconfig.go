package initconfig

import (
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

const defaultConfigTmpl = `projectName:
	templatePath:
	outputPath:
	usingGit: true
	mappingValue:
		test: test
`

func CreateConfig(cmd *cobra.Command, args []string, path, name string) bool {
	// check if path exists, if not, create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			cmd.Println("Error when creating template path: ", err)
			return false
		}
	}

	// create config file
	file, err := os.Create(path + "\\" + name + ".yaml")
	if err != nil {
		cmd.Println("Error when creating config file: ", err)
		return false
	}

	parse, err := template.New("config").Parse(defaultConfigTmpl)
	if err != nil {
		return false
	}

	err = parse.Execute(file, nil)

	// close file
	defer file.Close()

	return true
}
