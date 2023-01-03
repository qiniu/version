# qiniu/version

A fast way to add assorted version information for your application.

Note this is tailored primarily to Qiniu's internal needs, and may or may not
suit your specific case.

## License

`qiniu/version` is licensed under [the Apache 2.0 license](./LICENSE).

## Quick start

If you're in a hurry, just import the easy hook package:

```go
import _ "github.com/qiniu/version/v2/easyHook"
```

Build it with:

```shell
go build -ldflags "-X 'github.com/qiniu/version/v2.BuildDate=$(date)'" .
```

Then run your app with `version`, `-version` or `--version`, and you will see
output like:

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
	-X \"github.com/qiniu/version/v2.GitTag=$(git describe --tags)\" \
	-X \"github.com/qiniu/version/v2.BuildComments=${BUILDCOMMENTS}\" \
	-X \"github.com/qiniu/version/v2.Name=${NAME}\" \
```

Check the source code for details.

## Advanced usage

The easy hook package used in the above steps can bring about surprises.
For example, small programs designed to operate on a single input file path
argument will suddenly stop working with a path of `version`, because the
invocation `foo version` will be intercepted by this package and your `func main`
will not even have a chance to run.

If you want more control, you may simply import `github.com/qiniu/version/v2`
and call `version.Print()` or consume the individual values exposed there,
to your own liking.
