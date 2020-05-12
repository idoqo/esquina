package docs

import "github.com/idoqo/esquina/esquina"

// An internal server error occurs when the operation crashes the application and it cannot be proceed.
// swagger:response errorInternalServer
type InternalServerError struct {
	// in:body
	Body esquina.ErrorResponse
}