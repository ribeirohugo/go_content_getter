# Go Content Getter

**Go Content Getter** is a Go application that allows to download all images or content from a specified Url.
By defining a Regex pattern it is possible to adjust fetched image addresses or get different kind of content addresses.

## 1. Configurations

Configurations are stored in program root directory ``config.toml`` file.

| Parameter | Description | Type | Default | Required |
|:---|:---|:---|:---|:---|
| ``host`` | Host and port used for server mode application. | `string` | `localhost:8080` | **NO** |
| ``path`` | Destiny path where files will be stored. | `string` | ` ` | **NO** |
| ``regex`` | Fill this to replace regex expression to get content address from defined ``url``. | `string` | - | **NO** |
| ``url`` | Url address to get all images or content. | `string` | ` ` | **YES** |

## 2. Methods

- ``New(cfg config.Config)`` - Requires a ``Config`` struct with a URL and returns a new ``Getter`` struct.
- ``Get()`` - Returns a slice of image or content addresses, the page title and an error or a ``nil`` value.
- ``Download(folder string, images []string)`` - Receives a folder name and images or content slice addresses and returns and downloads files based on inserted data.

### 2.1. Example

```
cfg := config.Config{Url: "https://domain.com"}

getter := New(cfg)

title, images, err := getter.Get()

getter.Download(title, images)
```

## 3. Main methods

``console``, ``runnable`` and ``server`` are the three main files that allows to create a **Content Getter** with
its different ways of getting content.

* ``console`` - Allows to continuously insert URLs in a command line console to fetch their content, until operation is terminated by user.

* ``runnable`` - Allows to execute the application once and fetch all content based on ``config.toml`` URL attribute defined.

* ``server`` - Creates a Web Server based on config.toml attributes and allows users to insert content URLs through an HTML web page form.
