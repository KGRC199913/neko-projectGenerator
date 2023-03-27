package generate

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"projectGenerator/internal/codegen/shared/configs"
	"projectGenerator/internal/codegen/shared/constants"
	"strings"
	"text/template"
)

func GenerateProject(cmd *cobra.Command, args []string, mode int) (bool, error) {
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
			cmd.Println("Template path does not exist, failed")
			return false, err
		} else {
			cmd.Println("Template path exists, continue...")
		}
	}

	libPath := config.LibraryPath
	if libPath == "" {
		return false, errors.New("library path is empty")
	} else {
		// check library path exists, if not, create it, if yes, continue
		// if create failed, return error
		if _, err := os.Stat(libPath); os.IsNotExist(err) {
			cmd.Println("Library path does not exist, failed")
			return false, err
		} else {
			cmd.Println("Library path exists, continue...")
		}
	}

	filePath := config.FilePath
	if filePath == "" {
		return false, errors.New("file path is empty")
	} else {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			cmd.Println("File path does not exist, failed")
			return false, err
		} else {
			cmd.Println("File path exists, continue...")
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
			cmd.Println("Output path does not exist, creating...")
			cmd.Println("Creating output path at: ", outputPath)
			err := os.MkdirAll(outputPath, 0755)
			if err != nil {
				cmd.Println("Error when creating output path: ", err)
				return false, err
			}
		} else {
			cmd.Println("Output path exists, continue...")
		}
	}

	//if projectName == "" {
	//	return false, errors.New("project name is empty")
	//} else {
	//	// create project folder
	//	err := os.MkdirAll(outputPath+constants.PathDelimiter+projectName, 0755)
	//	if err != nil {
	//		cmd.Println("Error when creating project folder: ", err)
	//		return false, err
	//	}
	//}

	projectName := config.ProjectName
	if mode == 0 {
		cmd.Println("Generating project using template mode...")
		// generate project using template
		// if failed, return error
		err := generateProject(cmd, outputPath+constants.PathDelimiter+projectName, config)
		if err != nil {
			cmd.Println("Error when generating project: ", err)
			return false, err
		}
	} else if mode == 1 {
		cmd.Println("Generating project using copy mode...")
		// copy template to output path
		// if failed, return error
		err := copy(templatePath, outputPath+constants.PathDelimiter+projectName, config.MappingValue)
		if err != nil {
			cmd.Println("Error when copying template to output path: ", err)
			return false, err
		}
	}

	isUsingGit := config.UsingGit
	if isUsingGit {
		// create git repo
		// if failed, return error
		cmd.Println("Creating git repo at: ", outputPath+constants.PathDelimiter+config.ProjectName)
		err := exec.Command("git", "init", outputPath+constants.PathDelimiter+config.ProjectName).Run()
		if err != nil {
			cmd.Println("Error when creating git repo: ", err)
			return false, err
		}
	}

	return true, nil
}

type PathInfo struct {
	configs.FileInfo
	ParentPath string
	Children   []PathInfo
}

func generateProject(cmd *cobra.Command, destination string, config *configs.Config) error {
	generateTree := parseConfigToPathInfoTree(destination, config)
	// BFS to generate project
	var queue []PathInfo
	queue = append(queue, generateTree)
	for len(queue) > 0 {
		// pop
		cur := queue[0]
		queue = queue[1:]

		// generate current file or folder
		if cur.Type == "folder" {
			// create folder
			cmd.Println("Creating folder at: ", cur.ParentPath+constants.PathDelimiter+cur.Name)
			err := os.MkdirAll(cur.ParentPath+constants.PathDelimiter+cur.Name, 0755)
			if err != nil {
				cmd.Println("Error when creating folder: ", err)
				return err
			}
		} else if cur.Type == "file" {
			// create file
			if cur.IsTemplate {
				cmd.Println("Generating file using template "+cur.ParentPath+constants.PathDelimiter+cur.Name+" from: ", config.TemplatePath+constants.PathDelimiter+cur.Template)
				err := generateFileUsingTemplate(cur.ParentPath+constants.PathDelimiter+cur.Name, config.TemplatePath+constants.PathDelimiter+cur.Template, config.MappingValue)
				if err != nil {
					cmd.Println("Error when generating file using template: ", err)
					return err
				}
			} else if cur.IsLibrary {
				// copy library to destination
				// if failed, return error
				cmd.Println("Copying library from"+config.TemplatePath+constants.PathDelimiter+cur.LibraryName+" to: ", cur.ParentPath+constants.PathDelimiter+cur.Name)
				err := copy(config.TemplatePath+constants.PathDelimiter+cur.LibraryName, cur.ParentPath+constants.PathDelimiter+cur.Name, nil)
				if err != nil {
					cmd.Println("Error when copying library to destination: ", err)
					return err
				}
			} else {
				// create file
				// if failed, return error
				cmd.Println("Creating file from: ", config.FilePath+constants.PathDelimiter+cur.FileInfo.File+" to: ", cur.ParentPath+constants.PathDelimiter+cur.Name)
				err := createFile(config.FilePath+constants.PathDelimiter+cur.FileInfo.File, cur.ParentPath+constants.PathDelimiter+cur.Name)
				if err != nil {
					cmd.Println("Error when creating file: ", err)
					return err
				}
			}
		}

		// push children
		for _, v := range cur.Children {
			queue = append(queue, v)
		}
	}

	return nil
}

func createFile(src, des string) error {
	// create file
	file, err := os.Create(des)
	if err != nil {
		return err
	}
	defer file.Close()

	templateContent, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// write to file
	_, err = file.WriteString(string(templateContent))
	if err != nil {
		return err
	}

	return nil
}

func generateFileUsingTemplate(des string, templatePath string, value map[string]string) error {
	// create file
	file, err := os.Create(des)
	if err != nil {
		return err
	}
	defer file.Close()

	// read content from file with des ./templates/ + content
	// if failed, return error
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	// parse template
	t, err := template.New("template").Parse(string(templateContent))
	if err != nil {
		return err
	}

	// write to file
	err = t.Execute(file, value)
	if err != nil {
		return err
	}

	return nil
}

func parseConfigToPathInfoTree(destination string, config *configs.Config) PathInfo {
	// parse config to PathInfo
	root := PathInfo{
		FileInfo: configs.FileInfo{
			Name:     config.ProjectName,
			Type:     "folder",
			Children: config.ProjectStructure,
		},
		ParentPath: destination,
		Children:   parseChildrenToPathInfo(destination, config.ProjectStructure),
	}
	return root
}

func parseChildrenToPathInfo(curPath string, config []configs.FileInfo) []PathInfo {
	if config == nil || len(config) == 0 {
		return []PathInfo{}
	}

	var result []PathInfo
	for _, v := range config {
		if v.Type == "folder" {
			var children = parseChildrenToPathInfo(curPath+constants.PathDelimiter+v.Name, v.Children)
			if children != nil {
				for _, c := range children {
					c.ParentPath = curPath
				}
			}
			result = append(result, PathInfo{
				FileInfo:   v,
				ParentPath: curPath,
				Children:   children,
			})
		} else {
			result = append(result, PathInfo{
				FileInfo:   v,
				ParentPath: curPath,
			})
		}
	}
	return result
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

			if mappingValue == nil {
				err1 = os.WriteFile(filepath.Join(destination, relPath), data, 0755)
				if err1 != nil {
					return err1
				}
			} else {
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
			}
			return nil
		}
	})
	return err
}
