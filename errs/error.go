package errs

const (
	Parameters = iota
	Unauthorized
	Forbidden
	Validation
	NotFound
	Internal
)

// error response
type E struct {
	C int      `json:"code"`     // Business Code
	M []string `json:"messages"` // Error Messages
}
