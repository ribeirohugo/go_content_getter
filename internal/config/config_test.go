package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var configContent = `contentRegex = "[ab]"
titleRegex = "title"
url = "sub.domain"
path = "/path/to"
`

func TestConfig(t *testing.T) {
	const (
		contentRegexTest = "[ab]"
		titleRegexTest   = "title"
		pathTest         = "/path/to"
		urlTest          = "sub.domain"
	)

	var configTest = Config{
		ContentRegex: contentRegexTest,
		TitleRegex:   titleRegexTest,
		Host:         defaultHost,
		Path:         pathTest,
		URL:          urlTest,
	}

	tempFile, err := createTempFile()
	require.NoError(t, err)

	defer os.Remove(tempFile.Name())

	cfg, err := Load(tempFile.Name())
	require.NoError(t, err)
	assert.Equal(t, cfg, configTest)

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
