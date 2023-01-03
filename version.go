package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

const (
	unknownProperty = "N/A"
)

// Information of versioning
var (
	// GoVersion is the version of the Go toolchain that built the binary(for example, "go1.19.2").
	// and will be filled via runtime.Version() automatically if not specified.
	GoVersion = unknownProperty
	// GitCommit shows the revision identifier for the current commit or checkout.
	// and will be filled via debug.BuildInfo.Settings automatically if not specified.
	GitCommit = unknownProperty
	// GitCommitDate shows the time associated with GitCommit, in RFC3339 format
	// and will be filled via debug.BuildInfo.Settings automatically if not specified.
	GitCommitDate = unknownProperty
	// GitTreeState shows "dirty" indicating the source tree had local modifications, otherwise it is invisible.
	GitTreeState = unknownProperty
	// GitTag shows latest tag if injected by go -ldflags, otherwise it is invisible.
	GitTag = unknownProperty
	// BuildDate shows the built time for the associated binary if injected by go -ldflags. otherwise it is invisible.
	BuildDate = unknownProperty
	// Platform composes with GOARCH and GOOS automatically.
	Platform = unknownProperty
	// Compiler shows the toolchain flag used (typically "gc")
	Compiler = unknownProperty
	// BuildComments provides extra information if needed.
	BuildComments = unknownProperty
	// Name shows the name of your binary if provided, otherwise it is invisible.
	Name = unknownProperty
)

// This is for preventing access to the unpopulated properties.
func init() {
	collectFromBuildInfo()
	collectFromRuntime()
}

// Print prints out the collected version information.
func Print() {
	xprintf := func(k string, v string) {
		fmt.Printf("%s:\t%s\n", k, v)
	}

	if Name != unknownProperty {
		xprintf("Name", Name)
	}

	xprintf("Go version", GoVersion)
	xprintf("Git commit", GitCommit)
	xprintf("Commit date", GitCommitDate)

	if GitTreeState != unknownProperty {
		xprintf("Git state", GitTreeState)
	}

	if BuildDate != unknownProperty {
		xprintf("Build date", BuildDate)
	}

	if BuildComments != unknownProperty {
		xprintf("Build comments", BuildComments)
	}

	xprintf("OS/Arch", Platform)
	xprintf("Compiler", Compiler)

	if GitTag != unknownProperty {
		xprintf("Git tag", GitTag)
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
			if GitTreeState == unknownProperty && kv.Value == "true" {
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
