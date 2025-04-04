package config

import (
	"github.com/ribeirohugo/go_content_getter/pkg/patterns"
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

var configContentWithoutOptionalFields = `url = "sub.domain"
path = "/path/to"
`

var configContentInvalid = `url = 123
path = 123
`

func TestConfig(t *testing.T) {
	const (
		contentRegexTest = "[ab]"
		titleRegexTest   = "title"
		pathTest         = "/path/to"
		urlTest          = "sub.domain"
	)

	t.Run("valid config return", func(t *testing.T) {
		t.Run("with all fields", func(t *testing.T) {
			var configTest = Config{
				ContentRegex: contentRegexTest,
				TitleRegex:   titleRegexTest,
				Host:         defaultHost,
				Path:         pathTest,
				URL:          urlTest,
			}

			tempFile, err := createTempFile(configContent)
			require.NoError(t, err)

			defer os.Remove(tempFile.Name())

			cfg, err := Load(tempFile.Name())
			require.NoError(t, err)
			assert.Equal(t, cfg, configTest)

			tempFile.Close()
		})

		t.Run("without optional fields", func(t *testing.T) {
			var configTest = Config{
				ContentRegex: patterns.ImageContentFromHrefURL,
				TitleRegex:   patterns.HTMLTitle,
				Host:         defaultHost,
				Path:         pathTest,
				URL:          urlTest,
			}

			tempFile, err := createTempFile(configContentWithoutOptionalFields)
			require.NoError(t, err)

			defer os.Remove(tempFile.Name())

			cfg, err := Load(tempFile.Name())
			require.NoError(t, err)
			assert.Equal(t, cfg, configTest)

			tempFile.Close()
		})
	})

	t.Run("invalid config return", func(t *testing.T) {
		t.Run("file doesn't exist", func(t *testing.T) {
			cfg, err := Load("")
			assert.Equal(t, Config{}, cfg)
			assert.Error(t, err)
		})

		t.Run("invalid file content", func(t *testing.T) {
			tempFile, err := createTempFile(configContentInvalid)
			require.NoError(t, err)

			defer os.Remove(tempFile.Name())

			cfg, err := Load(tempFile.Name())
			assert.Equal(t, Config{}, cfg)
			assert.Error(t, err)

			tempFile.Close()
		})
	})
}

func createTempFile(fileContent string) (*os.File, error) {
	tempFile, err := os.CreateTemp("", "config.toml")
	if err != nil {
		return nil, err
	}

	_, err = tempFile.WriteString(fileContent)
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}
