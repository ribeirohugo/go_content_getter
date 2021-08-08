package config

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	pathTest  = "/path/to"
	regexTest = "[ab]"
	urlTest   = "sub.domain"
)

var configContent = `regex = "[ab]"
url = "sub.domain"
path = "/path/to"
`

var configTest = Config{
	Host:  defaultHost,
	Path:  pathTest,
	Regex: regexTest,
	Url:   urlTest,
}

func TestConfig(t *testing.T) {
	tempFile, err := createTempFile()
	if err != nil {
		t.Fatalf("Unexpected error creating temp file: %v.", err.Error())
	}

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

	_, err = tempFile.WriteString(configContent)
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}
