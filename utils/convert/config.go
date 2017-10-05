package convert

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/roscopecoltran/krakend-admin/utils/errors"
)

func JSON2YAML(j []byte) (string, error) {
	output, err := yaml.JSONToYAML(j)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}
	fmt.Println(string(output))
	return string(output), nil
}

func YAML2JSON(j []byte) (string, error) {
	output, err := yaml.YAMLToJSON(j)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}
	fmt.Println(string(output))
	return string(output), nil
}

func GetYamlFilenames(dirname string) ([]string, error) {
	files := make([]string, 0)
	d, err := ioutil.ReadDir(dirname)
	if err != nil {
		return files, err
	}
	for _, v := range d {
		if !v.IsDir() {
			filename := v.Name()
			fileext := strings.ToLower(path.Ext(filename))
			if fileext == ".yml" || fileext == ".yaml" {
				files = append(files, path.Join(dirname, filename))
			}
		}
	}
	return files, nil
}

func ConcatYamlFiles(files []string) ([]string, error) {
	d := make([]string, 0)
	for _, v := range files {
		content, err := ReadLines(v, "  ")
		if err != nil {
			return d, err
		}
		d = append(d, content...)
	}
	return d, nil
}

func ConcatYamlFilesFromDir(dirname string) []string {
	data := make([]string, 0)
	files, err := GetYamlFilenames(dirname)
	if err == nil {
		data, err = ConcatYamlFiles(files)
		if err != nil {
			errors.PrintError(err)
		}
	}
	return data
}

func ReadLines(path, leftPad string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, leftPad+scanner.Text())
	}
	return lines, scanner.Err()
}
