// Package version provides access to some of your binary's versioning
// information.
//
// The easy way to integrate this package, but could bring about surprising
// behavior (read the source for details, it's just a one-liner!):
//
//	import _ "github.com/qiniu/version/v2/easyHook"
//
// And if you want more control and less magic:
//
//	import "github.com/qiniu/version/v2"
//
//	func main() {
//		// make use of one of the prepopulated values
//		_ = version.GitCommit
//		// or output in the package-standard format in your CLI handler
//		version.Print()
//	}
package version
