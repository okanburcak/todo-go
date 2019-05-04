package config

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func loadConfigFiles() map[string]string {
	Init("test")
	path, _ := os.Getwd()

	files := make(map[string]string)
	_ = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			r, e := regexp.MatchString("yaml", info.Name())
			if e == nil && r {
				files[info.Name()] = path
			}
		}
		return nil
	})

	return files
}

func TestConfigFilesExists(t *testing.T) {
	files := loadConfigFiles()

	assert.True(t, len(files["development.yaml"]) > 0)
	assert.True(t, len(files["test.yaml"]) > 0)
	assert.True(t, len(files["production.yaml"]) > 0)

}

func TestConfigFilesDoesNotExists(t *testing.T) {
	files := loadConfigFiles()

	assert.False(t, len(files["dev.yaml"]) > 0)
	assert.False(t, len(files["t.yaml"]) > 0)
	assert.False(t, len(files["prod.yaml"]) > 0)

}

func TestYamlFilesCanUnmarshal(t *testing.T) {
	files := loadConfigFiles()

	for _, v := range files {
		yamlFile, _ := ioutil.ReadFile(v)
		err := yaml.Unmarshal(yamlFile, GetConfig())
		assert.Nil(t, err)
	}
}
