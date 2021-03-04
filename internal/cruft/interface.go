package cruft

// Crufter defines an interface for encrufting strings:
type Crufter interface {
	Encruft(value string) (string, error)
}
