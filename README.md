# Go Content Getter

**Go Content Getter** is a Go application that allows to download all images or content from a specified Url.
By defining a Regex pattern it is possible to adjust fetched image addresses or get different kind of content addresses.

## 1. Configurations

Configurations are stored in program root directory ``config.toml`` file.

| Parameter         | Description                                                       | Type       | Default                       | Required |
|:------------------|:------------------------------------------------------------------|:-----------|:------------------------------|:---------|
| `host`            | Host and port used for server mode application.                   | `string`   | `localhost:8080`              | **NO**   |
| `path`            | Destination path where files will be stored.                      | `string`   | ` `                           | **NO**   |
| `content_regex`   | Regular expression to find content over defined `url` address.    | `string`   | `ImageSrc` <sup>1</sup>       | **NO**   |
| `title_regex`     | Regular expression to find page title over defined `url` address. | `string`   | `HTMLTitle` <sup>2</sup>      | **NO**   |
| `url`             | Url address list to get all images or content.                    | `[]string` | ` `                           | **YES**  |
| `allowed_origins` | List of allowed CORS origins for server mode.                     | `[]string` | `[""http://localhost:3000""]` | **NO**   |

<sup>1</sup> - `ImageSrc` is the following regex source: 
``
src=["'](http[s]?://[a-zA-Z0-9/._-]+(?::[0-9]+)?/[a-zA-Z0-9/._-]*[.](?:jpg|gif|png))(?:[?&#].*)?["']
``

<sup>2</sup> - `HTMLTitle` is the following regex source:
``
(?:\<title\>)(.*)(?:<\/title\>)
``

## 2. Domain

1. `Config` - (optional entity) holds configuration to start a `Getter` source

2. `Page` - Base entity that holds a Page to find for data. Has page `Title` and `Content` attributes.

3. `Target` - Are URLs found in a `Page` to be downloaded after.

4. `download` package allows to download one or many targets.

5. `File` - Is the result from a downloaded `Target` and hold filename and content in bytes.

6. `store` package allows to store in a file, the content from a `File`.

7. `source` package holds a `Getter` with `Get` and `GetAndStore` methods, that compiles all the sequence
from previous entities.

## 3. Main methods

``console``, ``runnable`` and ``server`` are the three main files that allows to create a **Content Getter** with
its different ways of getting content.

* ``console`` - Allows to continuously insert URLs in a command line console to fetch their content, until operation is
terminated by user.

* ``runnable`` - Allows to execute the application once and fetch all content based on ``config.toml`` URL attribute defined.

* ``server`` - Creates a Web Server based on config.toml attributes and allows users to insert content URLs through an
HTML web page form.

## 4. Implementation

```
cfg, err := config.Load("config.toml")
if err != nil {
    log.Fatal(err)
}

sourceGetter := source.New(cfg.Path, cfg.ContentRegex, cfg.TitleRegex)

content, err := sourceGetter.GetAndStore(url)
if err != nil {
    log.Println(err)
}
```

## 5. Server

Server mode allows to run a web server API that allows to use `go_content_getter` features through a web interface.

### 5.1 Build and Run

To run the server mode, execute the following command:

``go run cmd/server/server.go``

Then open your browser and navigate to ``http://localhost:8080``
(or the host and port defined in your config.toml file).

### 5.2. Frontend

The frontend can be accessed using ReactJS `frontend` folder, and it allows to access server API with user interface.
To run this application execute the following command:

``npm start``

For more information, please refer to the [frontend/README.md](frontend/README.md) file.
