# dailyreport

A Go program to generate and share daily work reports. License: MGPL.v1.4. All rights reserved.

## Quick Start

1. Install Go 1.20 or later. Install protoc and protoc-gen-go.

    ```bash
    $ go version
    go version go1.20.5 linux/amd64
    ```

    ```bash
    sudo apt install -y protobuf-compiler
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    ```

    And add `$GOPATH/bin` to your `$PATH` environment variable. In Linux, like this:

    ```bash
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```

2. Clone this repository.

    ```bash
    git clone https://github.com/mizumoto-cn/dailyreport.git
    ```

3. Run

    ```bash
    make setup
    ```

    to copy the template and configuration files to the directories.
    then after completing the [configuration](#configuration-file), run:

    ```bash
    make run-cmd
    ```

    ![()](./template/run-cmd.png)

    Or you can run the program directly after build

    ```bash
    make build-run
    ```

    ![()](./template/makebuild.png)
    ![()](./template/makebuildrun.png)

4. Use Crontab for periodically running the program. (Linux or **Windows WSL**(recommended))

    ```bash
    crontab -e
    ```

    Add the following line to the crontab file. After you build it.

    ```bash
    55 17 * * 1-5 /path/to/dailyreport/bin/dailyreport
    ```

    This will run the program at 17:55 from Monday to Friday.

    ![()](./template/crontab.png)
    ![()](./template/crontab2.png)

> See [template](./template) folder for the template files.
> > [template/dailyreport.txt](./template/dailyreport.txt) is the template for the daily report content.
> > [template/template.html](./template/template.html) is the template for the daily report email.
> > [template/config.yaml.tmp](./template/config.yaml.tmp) is the template for the configuration file.

## Template & Configuration Guide

### Configuration File

The configuration file is a YAML file. The default file is `configs/config.yaml`. You can change the file path by setting the `init()` function in [cmd/main.go](./cmd/main.go).

```go
var (
    // flagconf is the config flag.
    flagconf string
)

func init() {
    flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
}

func main() {
    flag.Parse()
    c := config.New(
        config.WithSource(
            file.NewSource(flagconf),
        ),
    )
```

A typical configuration file looks like [this](./template/config.yaml.tmp):

```yaml
name: dailyreport
smtp_dialer:
  insecure_skip_verify: true
  host: smtp.gmail.com
  port: 587
  username: username  # your email address
  password: password  # your email password
  to: 
    - mail@example.com
    - mail2@example.com
  template_path: ./template.html  # mail body template

path:
  contents_path: ./dailyreport.txt  # daily report content
```

### Template File & Contents

The template file contains a format string for the e-mail body. You can add `%s` placeholders into the template. Standard html syntax is supported.

e.g.:

`template`:

```html
Hello, %s %s
```

`contents`:

```txt
World!
wubba dubba
```

`result`:

```html
Hello, <p>World!</p> <p>wubba dubba</p>
```

## License

This project is licensed under the Mizumoto General Public License, Version 1.4. You can find the full license text in [LICENSE](./LICENSE/Mizumoto.General.Public.License.v1.4.md).

The _entities_\* on the [Disqualified Entities List](./LICENSE/List_of_Disqualified_Entities.md) are prohibited from using _files_\*\* from this project in any way.

---
> \*/\*\*: See chapter [#Restrictions for Users](./LICENSE/Mizumoto.General.Public.License.v1.4.md/#restrictions-for-users) for definitions of _entities_ and _files_.
