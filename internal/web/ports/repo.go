package ports

// Repository ...
type Repository interface {
	Accounts
	Operators
	Users
	Me
}
