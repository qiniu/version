# version
package version defines the utility information for versioning binary

## Usage

Import this package in your application:

```go
import _ "github.com/qiniu/version"
```

Build it with:

```shell
go build -ldflags "-X 'github.com/qiniu/version.BuildDate=$(date)'" .
```

Then run, it will output like: