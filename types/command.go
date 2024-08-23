package types

type Command struct {
	Name    string
	Description string
	Alias   []string
	Flags   []Flag
	Handler func([]Flag)
}
