package fazzrepository

// emptyResultError appear when data returned as empty
type emptyResultError struct{}

// Error return error text
func (e *emptyResultError) Error() string {
	return "no rows in result set"
}

// NewEmptyResultError creates new instance of EmptyResultError
func NewEmptyResultError() error {
	return &emptyResultError{}
}

// IsEmptyResult check if instance of error is EmptyResultError
func IsEmptyResult(err error) bool {
	if _, ok := err.(*emptyResultError); ok {
		return true
	}
	return false
}
