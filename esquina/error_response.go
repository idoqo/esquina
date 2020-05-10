package esquina

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrorResponse struct {
	Err        error  `json:"-"`               // low-level error
	StatusCode int    `json:"-"`               // http status code
	StatusText string `json:"status"`          // user-level status message
	ErrorText  string `json:"error,omitempty"` // app-level error message
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

// ErrorInternalServer returns error response with status=500 and given error
func ErrorInternalServer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:            err,
		StatusCode: 500,
		StatusText:     "Internal Server Error",
		ErrorText:      err.Error(),
	}
}

var ErrorNotFound = &ErrorResponse{
	StatusCode: 404,
	StatusText: "Page not Found",
}

var ErrorMethodNotAllowed = &ErrorResponse{
	StatusCode: 405,
	StatusText: "Method not Allowed",
}
