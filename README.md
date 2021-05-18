# Go Image Getter

**Go Image Getter** is a Go application that allows to download all images from a specified Url.
By defining a Regex pattern it is possible to adjust fetched image addresses or get different kind of content addresses.

## 1. Configurations

| Parameter | Description | Type | Default | Required |
|:---|:---|:---|:---|:---|
| ``regex`` | Fill this to replace regex expression to get content address from defined ``url``. | `string` | - | **NO** |
| ``url`` | Url address to get all images or content. | `string` | ` ` | **YES** |

## 2. Methods

- ``New(url string, regex ...string)`` - Requires an Url and optionally a regex pattern and returns a new ``Getter`` struct.
- ``Get()`` - Returns a slice of image or content addresses, the page title and an error or a ``nil`` value.
- ``Download(folder string, images []string)`` - Receives a folder name and images or content slice addresses and returns and downloads files based on inserted data.

### 2.1. Example

```
cfg := config.Config{Url: "https://domain.com"}

getter := New(cfg)

title, images, err := getter.Get()

getter.Download(title, images)
```
