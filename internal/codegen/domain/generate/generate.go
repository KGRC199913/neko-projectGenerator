package generate

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"projectGenerator/internal/codegen/shared/configs"
	"strings"
	"text/template"
)

func GenerateProject(cmd *cobra.Command, args []string) (bool, error) {
	// read config
	config := configs.GetConfig()

	// check template path
	templatePath := config.TemplatePath
	if templatePath == "" {
		return false, errors.New("template path is empty")
	} else {
		// check template path exists, if not, create it, if yes, continue
		// if create failed, return error
		if _, err := os.Stat(templatePath); os.IsNotExist(err) {
			err := os.MkdirAll(templatePath, 0755)
			if err != nil {
				cmd.Println("Error when creating template path: ", err)
				return false, err
			}
		}
	}

	// check output path
	outputPath := config.OutputPath
	if outputPath == "" {
		return false, errors.New("output path is empty")
	} else {
		// check output path exists, if not, create it, if yes, continue
		// if create failed, return error
		if _, err := os.Stat(outputPath); os.IsNotExist(err) {
			err := os.MkdirAll(outputPath, 0755)
			if err != nil {
				cmd.Println("Error when creating output path: ", err)
				return false, err
			}
		}
	}

	projectName := config.ProjectName
	if projectName == "" {
		return false, errors.New("project name is empty")
	} else {
		// create project folder
		err := os.MkdirAll(outputPath+"\\"+projectName, 0755)
		if err != nil {
			cmd.Println("Error when creating project folder: ", err)
			return false, err
		}
	}

	isUsingGit := config.UsingGit
	if isUsingGit {
		// create git repo
		// if failed, return error
		err := exec.Command("git", "init", "./").Run()
		if err != nil {
			cmd.Println("Error when creating git repo: ", err)
			return false, err
		}
	}

	// copy template to output path
	// if failed, return error
	err := copy(templatePath, outputPath+"\\"+projectName, config.MappingValue)
	if err != nil {
		cmd.Println("Error when copying template to output path: ", err)
		return false, err
	}

	return true, nil
}

func copy(source, destination string, mappingValue map[string]string) error {
	var err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		var relPath = strings.Replace(path, source, "", 1)
		if relPath == "" {
			return nil
		}
		if info.IsDir() {
			return os.Mkdir(filepath.Join(destination, relPath), 0755)
		} else {
			filePath := filepath.Join(source, relPath)
			var data, err1 = os.ReadFile(filePath)
			if err1 != nil {
				return err1
			}
			parse, err1 := template.New(filePath).Parse(string(data))
			if err1 != nil {
				return err1
			}

			file, err1 := os.OpenFile(filepath.Join(destination, relPath), os.O_CREATE|os.O_WRONLY, 0755)
			if err1 != nil {
				return err1
			}

			err1 = parse.Execute(file, mappingValue)
			if err1 != nil {
				return err1
			}

			return nil
		}
	})
	return err
}
