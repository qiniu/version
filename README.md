# version
fast way to add version information for your application

## Usage

Import this package in your application:

```go
import _ "github.com/qiniu/version"
```

Build it with:

```shell
go build -ldflags "-X 'github.com/qiniu/version.BuildDate=$(date)'" .
```

Then run your app with parameters `version` or `--version`, you will get output like:

```shell
✗ ./examples version
Go version:     go1.19.3
Git commit:     56dac6d5e895d5d6474b840d99fb3c7cfbdf26e3
Commit date:    2022-12-06T12:22:24Z
Git state:      dirty
Built date:     Wed Dec  7 14:11:24 CST 2022
OS/Arch:        darwin/amd64
Compiler:       gc
```

Also, there are some other useful options for your reference. For example:

```shell
LDFLAGS="${LDFLAGS} \
	-X \"github.com/qiniu/version.GitTag=$(git describe --tags)\" \
	-X \"github.com/qiniu/version.BuildComments=${BUILDCOMMENTS}\" \
	-X \"github.com/qiniu/version.Name=${NAME}\" \
```