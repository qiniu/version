package easyHook

import (
	"os"

	"github.com/qiniu/version/v2"
)

func init() {
	// usages: <command> version (or --version / -version)
	if len(os.Args) > 1 && (os.Args[1] == "version" || os.Args[1] == "--version" || os.Args[1] == "-version") {
		version.Print()
		os.Exit(0)
	}

	// TODO: expose version information via exporter as needed
}
