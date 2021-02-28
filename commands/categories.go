package commands

type Category int

const (
	InfoCategory Category = 1 << iota
	AdminCategory
	AudioCategory
)
