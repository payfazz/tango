package fazzrepository

// EmptyResultError appear when data returned as empty
type EmptyResultError struct{}

// Error return error text
func (e *EmptyResultError) Error() string {
	return "fazzrepository: empty result"
}

// NewEmptyResultError creates new instance of EmptyResultError
func NewEmptyResultError() error {
	return &EmptyResultError{}
}
