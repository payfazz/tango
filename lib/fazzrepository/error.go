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

// IsEmptyResult check if instance of error is EmptyResultError
func IsEmptyResult(it interface{}) bool {
	switch it.(type) {
	case EmptyResultError:
		return true
	case *EmptyResultError:
		return true
	default:
		return false
	}
}
