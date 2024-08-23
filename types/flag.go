package types

type Flag struct {
	Name  string
	Alias []string // Alias for the flag (e.g. -v for -verbose)
	Value string
}
