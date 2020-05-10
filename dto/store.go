package dto

import "net/http"

type Store struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OwnerID     int    `json:"owner_id"`
	Description string `json:"description"`
}

type StoreList struct {
	Count int    `json:"count"`
	List []Store `json:"list"`
}

func (store *Store) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (store *Store) Bind(r *http.Request) error {
	return nil
}
