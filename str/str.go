// Package str some string conversion related packages
package str

// DefaultIntBase if not set during type conversion, the default is decimal
const DefaultIntBase = 10

// Options A Builder is the base struct for type conversion
type Options struct {
	base int
}
