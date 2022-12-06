package version

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
)

const (
	unknownProperty = "N/A"
)

// Information of versioning
var (
	// GoVersion is the version of the Go toolchain that built the binary(for example, "go1.19.2").
	GoVersion = unknownProperty
	// GitCommit shows the revision identifier for the current commit or checkout.
	GitCommit = unknownProperty
	// GitCommitDate shows the time associated with GitCommit, in RFC3339 format
	GitCommitDate = unknownProperty
	// GitTreeState shows "dirty" indicating the source tree had local modifications, otherwise it is invisible.
	GitTreeState = unknownProperty
	// GitTag shows latest tag if injected by go -ldflags. otherwise it is invisible by default.
	GitTag = unknownProperty
	// BuildDate shows the built time for the associated binary
	BuildDate = unknownProperty
	// Platform composes with GOARCH and GOOS
	Platform = unknownProperty
	// Compiler shows the toolchain flag used (typically "gc")
	Compiler = unknownProperty
)

func init() {
	// usages: <command> version or <command> --version
	if len(os.Args) > 1 && (os.Args[1] == "version" || os.Args[1] == "--version") {
		Version()
		os.Exit(0)
	}

	// TODO: expose version information via exporter as needed
}

var once sync.Once

// Version prints the information of versioning
func Version() {
	once.Do(func() {
		collectFromBuildInfo()
		collectFromRuntime()
	})

	format := "%s:\t%s\n"
	xprintf(format, "Go version", GoVersion)
	xprintf(format, "Git commit", GitCommit)
	xprintf(format, "Commit date", GitCommitDate)

	if GitTreeState != unknownProperty {
		xprintf(format, "Git tree state", GitTreeState)
	}

	xprintf(format, "Built date", BuildDate)
	xprintf(format, "OS/Arch", Platform)
	xprintf(format, "Compiler", Compiler)

	if GitTag != unknownProperty {
		xprintf(format, "Git tag", GitTag)
	}

}

// collectFromBuildInfo tries to set the build information embedded in the running binary via Go module.
// It doesn't override data if were already set by Go -ldflags.
func collectFromBuildInfo() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	for _, kv := range info.Settings {
		switch kv.Key {
		case "vcs.revision":
			if GitCommit == unknownProperty && kv.Value != "" {
				GitCommit = kv.Value
			}
		case "vcs.time":
			if GitCommitDate == unknownProperty && kv.Value != "" {
				GitCommitDate = kv.Value
			}

		case "vcs.modified":
			if GitTreeState == unknownProperty && kv.Value != "" {
				GitTreeState = "dirty"
			}
		}
	}
}

// collectFromRuntime tries to set the build information embedded in the running binary via go runtime.
// It doesn't override data if were already set by Go -ldflags.
func collectFromRuntime() {
	if GoVersion == unknownProperty {
		GoVersion = runtime.Version()
	}

	if Platform == unknownProperty {
		Platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	}

	if Compiler == unknownProperty {
		Compiler = runtime.Compiler
	}
}

// xprintf prints a message to standard output.
func xprintf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
