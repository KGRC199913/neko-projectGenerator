package initconfig

import (
	"github.com/spf13/cobra"
	"os"
	"projectGenerator/internal/codegen/shared/constants"
	"text/template"
)

const defaultConfigTmpl = `projectName: {{.projectName}}
templatePath: {{.currentDir}}templates
outputPath: {{.currentDir}}output
libraryPath: {{.currentDir}}libs
filePath: {{.currentDir}}files
usingGit: true
projectStructure:
- name: libs
  type: folder
  children:
    - name: test.jar
      type: file
      libraryName: test.jar
      isLibrary: true
- name: src 
  type: folder
  children:
    - name: main
      type: folder
    - name: test
      type: folder
- name: .gitignore
  type: file
  template: ___gitignore.tmpl
  isTemplate: true
mappingValue:
  - test: test
  - test1: test1
`

func CreateConfig(cmd *cobra.Command, args []string, path, name, projectName, envOs string) bool {
	// check if path exists, if not, create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			cmd.Println("Error when creating template path: ", err)
			return false
		}
	}

	// create config file
	file, err := os.Create(path + constants.PathDelimiter + name)
	if err != nil {
		cmd.Println("Error when creating config file: ", err)
		return false
	}

	parse, err := template.New("config").Parse(defaultConfigTmpl)
	if err != nil {
		return false
	}

	data := make(map[string]interface{})
	if projectName != "" {
		data["projectName"] = projectName
	} else {
		data["projectName"] = "placeholder"
	}
	data["delimiter"] = constants.PathDelimiter
	data["currentDir"] = constants.CurrentDir

	err = parse.Execute(file, data)

	// close file
	defer file.Close()

	return true
}
