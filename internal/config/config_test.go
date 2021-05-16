package config

import (
	"io/ioutil"
	"os"
	"testing"
)

var configContent = `regex = "[ab]"
url="sub.domain"
`

var configTest = Config{
	Regex: "[ab]",
	Url:   "sub.domain",
}

func TestConfig(t *testing.T) {

	tempFile, _ := createTempFile()

	defer os.Remove(tempFile.Name())

	cfg, _ := Load(tempFile.Name())

	if cfg != configTest {
		t.Errorf("Wrong config file output,\n got: %v,\n want: %v.", cfg, configTest)
	}

	tempFile.Close()
}

func createTempFile() (*os.File, error) {

	tempFile, err := ioutil.TempFile("", "config.toml")
	if err != nil {
		return nil, err
	}

	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(configContent)
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}
