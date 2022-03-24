package parsers

type Details struct {
	Version string
	Name    string
}

type IParser interface {
	Parse(filename string) (*Details, error)
}
