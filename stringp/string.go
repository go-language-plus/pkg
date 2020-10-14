package stringp

// DefaultStringBase if not set during type conversion, the default is decimal
const DefaultStringBase = 10

// Builder A Builder is the base struct for type conversion
type Builder struct {
	Base int
}
