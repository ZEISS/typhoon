package ports

// Build is the interface that wraps the methods to get the build information.
type Build interface {
	Version() (string, error)
	Build() (string, error)
	Date() (string, error)
}
