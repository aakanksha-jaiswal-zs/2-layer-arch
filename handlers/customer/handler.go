package customer

import (
	"log"
	"net/http"

	"sample-api/stores"
)

type handler struct {
	store stores.Customer
}

func New(store stores.Customer) handler {
	return handler{store: store}
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if ok := h.store.Get(name); ok {
		_, err := w.Write([]byte("Welcome " + name))
		if err != nil {
			log.Println("error in writing response")
		}

		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {

}
