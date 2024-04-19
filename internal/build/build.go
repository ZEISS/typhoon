package build

import "fmt"

var (
	version = "dev"
	commit  = ""
	date    = ""
)

// Version returns the version of the build.
func Version() string {
	return version
}

// Build returns the commit of the build.
func Build() string {
	return fmt.Sprintf("%s@%s", version, commit)
}

// Date returns the date of the build.
func Date() string {
	return date
}
