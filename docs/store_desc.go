package docs

import "github.com/idoqo/esquina/dto"

// A store created by a user. They contain catalogs (products that are similar)
// which are also created by the store owner.

// swagger:route POST /store stores createStore
// Creates a new store owned by the authenticated user in the database.
// responses:
//		200: responseStoreCreated
//		500: errorInternalServer

// Store was successfully created
// swagger:response responseStoreCreated
type StoreCreatedResponse struct {
	// in:body
	Body dto.Store
}
