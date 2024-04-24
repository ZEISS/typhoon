package ports

// Repositories is the interface that wraps the methods to access data.
type Repositories interface {
	Systems
	Teams
	Accounts
	Operators
}
